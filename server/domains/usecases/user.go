package usecases

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/pkg/hash"
	"github.com/bagasunix/bank-ina/server/domains/data/models"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/entities"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
	"github.com/bagasunix/bank-ina/server/endpoints/responses"
)

type User interface {
	CreateUser(ctx *gin.Context, req *requests.CreateUser) (response *responses.EntityId, err error)
	UpdateUser(ctx *gin.Context, req *requests.UpdateUser) (response *responses.ViewEntity[*entities.User], err error)
	DeleteUser(ctx *gin.Context, req *requests.EntityId) (response *responses.Empty, err error)
	GetAllUser(ctx *gin.Context) (response *responses.ListEntity[entities.User], err error)
	ViewUser(ctx *gin.Context, req *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error)
}

type user struct {
	repo   repositories.Repositories
	logger *zap.Logger
}

// CreateUser implements User.
func (u *user) CreateUser(ctx *gin.Context, req *requests.CreateUser) (response *responses.EntityId, err error) {
	responseBuilder := responses.NewEntityIdBuilder()
	if req.Validate() != nil {
		return responseBuilder.Build(), req.Validate()
	}

	checkName := u.repo.GetUser().GetByEmail(ctx, req.Name)
	if len(checkName.Value) != 0 {
		return responseBuilder.Build(), errors.CustomError(errors.ERR_DUPLICATE_KEY)
	}

	mBuild := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hash.HashAndSalt([]byte(req.Password)),
	}

	if err := u.repo.GetUser().CreateUser(ctx, &mBuild); err != nil {
		return responseBuilder.Build(), err
	}

	return responseBuilder.SetId(mBuild.ID).Build(), nil
}

// DeleteUser implements User.
func (u *user) DeleteUser(ctx *gin.Context, req *requests.EntityId) (response *responses.Empty, err error) {
	responseBuild := responses.NewEmptyBuilder()
	if err = req.Validate(); err != nil {
		responseBuild.SetCode(400)
		responseBuild.SetMsg("Gagal validasi id")
		return responseBuild.Build(), err
	}
	result := u.repo.GetUser().GetByID(ctx, strconv.Itoa(req.Id.(int)))
	if result.Error != nil {
		responseBuild.SetCode(404)
		responseBuild.SetMsg("Data yang akan dihapus tidak ada")
		return responseBuild.Build(), result.Error
	}

	if err = u.repo.GetUser().DeleteUser(ctx, strconv.Itoa(req.Id.(int))); err != nil {
		responseBuild.SetCode(404)
		responseBuild.SetMsg("Data gagal dihapus")
		return responseBuild.Build(), result.Error
	}
	responseBuild.SetMsg("Data berhasil dihapus")
	responseBuild.SetCode(200)
	return responseBuild.Build(), nil
}

// GetAllUser implements User.
func (u *user) GetAllUser(ctx *gin.Context) (response *responses.ListEntity[entities.User], err error) {
	responseBuilder := responses.NewListEntityBuilder[entities.User]()
	mBuild := u.repo.GetUser().GetAllUser(ctx)
	if mBuild.Error != nil {
		return responseBuilder.Build(), mBuild.Error
	}

	var users []entities.User
	for _, i := range mBuild.Value {
		userBuild := entities.NewUserBuilder()
		userBuild.SetID(strconv.Itoa(i.ID))
		userBuild.SetName(i.Name)
		userBuild.SetEmail(i.Email)
		userBuild.SetCreatedAt(i.CreatedAt)
		userBuild.SetUpdatedAt(i.UpdatedAt)
		users = append(users, *userBuild.Build())
	}

	return responseBuilder.SetData(users).Build(), nil
}

// UpdateUser implements User.
func (u *user) UpdateUser(ctx *gin.Context, req *requests.UpdateUser) (response *responses.ViewEntity[*entities.User], err error) {
	responseBuilder := responses.NewViewEntityBuilder[*entities.User]()
	if req.Validate() != nil {
		return responseBuilder.Build(), req.Validate()
	}

	checkName := u.repo.GetUser().GetByEmail(ctx, req.Email)
	if len(checkName.Value) != 0 {
		return responseBuilder.Build(), errors.ErrDataAlready(req.Email)
	}

	ids, _ := strconv.Atoi(req.ID)

	mBuild := models.NewUserBuilder()
	mBuild.SetId(ids)
	mBuild.SetName(req.Name)
	mBuild.SetEmail(req.Email)
	mBuild.SetPassword(req.Password)

	if err = u.repo.GetUser().UpdateUser(ctx, req.ID, mBuild.Build()); err != nil {
		return responseBuilder.Build(), err
	}

	resultUser := u.repo.GetUser().GetByID(ctx, req.ID)
	if err = resultUser.Error; err != nil {
		return responseBuilder.Build(), err

	}

	entitiesUser := entities.NewUserBuilder()
	entitiesUser.SetID(strconv.Itoa(resultUser.Value.ID))
	entitiesUser.SetName(resultUser.Value.Name)
	entitiesUser.SetEmail(resultUser.Value.Email)
	entitiesUser.SetCreatedAt(resultUser.Value.CreatedAt)
	entitiesUser.SetUpdatedAt(resultUser.Value.UpdatedAt)

	return responseBuilder.SetData(entitiesUser.Build()).Build(), nil
}

// ViewUser implements User.
func (u *user) ViewUser(ctx *gin.Context, req *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error) {
	responseBuilder := responses.NewViewEntityBuilder[*entities.User]()
	if err = req.Validate(); err != nil {
		return responseBuilder.Build(), err
	}
	mBuild := u.repo.GetUser().GetByID(ctx, strconv.Itoa(req.Id.(int)))
	if mBuild.Error != nil {
		return responseBuilder.Build(), mBuild.Error
	}

	userBuild := entities.NewUserBuilder()
	userBuild.SetID(strconv.Itoa(mBuild.Value.ID))
	userBuild.SetName(mBuild.Value.Name)
	userBuild.SetEmail(mBuild.Value.Email)
	userBuild.SetCreatedAt(mBuild.Value.CreatedAt)
	userBuild.SetUpdatedAt(mBuild.Value.UpdatedAt)

	return responseBuilder.SetData(userBuild.Build()).Build(), nil

}

func NewUser(logger *zap.Logger, repo repositories.Repositories) User {
	a := new(user)
	a.repo = repo
	a.logger = logger
	return a
}
