package usecases

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/domains/data/models"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/entities"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
	"github.com/bagasunix/bank-ina/server/endpoints/responses"
)

type Task interface {
	CreateTask(ctx *gin.Context, req *requests.CreateTask) (response *responses.EntityId, err error)
	UpdateTask(ctx *gin.Context, req *requests.UpdateTask) (response *responses.ViewEntity[*entities.Task], err error)
	DeleteTask(ctx *gin.Context, req *requests.EntityId) (response *responses.Empty, err error)
	GetAllTask(ctx *gin.Context) (response *responses.ListEntity[entities.Task], err error)
	ViewTask(ctx *gin.Context, req *requests.EntityId) (response *responses.ViewEntity[*entities.Task], err error)
}
type task struct {
	repo   repositories.Repositories
	logger *zap.Logger
}

// CreateTask implements Task.
func (t *task) CreateTask(ctx *gin.Context, req *requests.CreateTask) (response *responses.EntityId, err error) {
	authUser := ctx.Value("authorization_payload").(*entities.User)
	responseBuilder := responses.NewEntityIdBuilder()
	if req.Validate() != nil {
		return responseBuilder.Build(), req.Validate()
	}

	checkName := t.repo.GetTask().GetByName(ctx, req.Title)
	if len(checkName.Value) != 0 {
		return responseBuilder.Build(), errors.CustomError(errors.ERR_DUPLICATE_KEY)
	}

	userID, _ := strconv.Atoi(authUser.ID)
	mBuild := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      userID,
	}

	if err := t.repo.GetTask().CreateTask(ctx, &mBuild); err != nil {
		return responseBuilder.Build(), err
	}

	return responseBuilder.SetId(mBuild.ID).Build(), nil
}

// DeleteTask implements Task.
func (t *task) DeleteTask(ctx *gin.Context, req *requests.EntityId) (response *responses.Empty, err error) {
	responseBuild := responses.NewEmptyBuilder()
	if err = req.Validate(); err != nil {
		responseBuild.SetCode(400)
		responseBuild.SetMsg("Gagal validasi id")
		return responseBuild.Build(), err
	}
	result := t.repo.GetTask().GetByID(ctx, strconv.Itoa(req.Id.(int)))
	if result.Error != nil {
		responseBuild.SetCode(404)
		responseBuild.SetMsg("Data yang akan dihapus tidak ada")
		return responseBuild.Build(), result.Error
	}

	if err = t.repo.GetTask().DeleteTask(ctx, strconv.Itoa(req.Id.(int))); err != nil {
		responseBuild.SetCode(404)
		responseBuild.SetMsg("Data gagal dihapus")
		return responseBuild.Build(), result.Error
	}
	responseBuild.SetMsg("Data berhasil dihapus")
	responseBuild.SetCode(200)
	return responseBuild.Build(), nil
}

// GetAllTask implements Task.
func (t *task) GetAllTask(ctx *gin.Context) (response *responses.ListEntity[entities.Task], err error) {
	responseBuilder := responses.NewListEntityBuilder[entities.Task]()
	mBuild := t.repo.GetTask().GetAllTask(ctx)
	if mBuild.Error != nil {
		return responseBuilder.Build(), mBuild.Error
	}

	var tasks []entities.Task
	for _, i := range mBuild.Value {
		taskBuild := entities.NewTaskBuilder()
		taskBuild.SetId(strconv.Itoa(i.ID))
		taskBuild.SetTitle(i.Title)
		taskBuild.SetDescription(i.Description)
		taskBuild.SetStatus(i.Status)
		taskBuild.SetUserID(strconv.Itoa(i.UserID))
		taskBuild.SetCreatedAt(i.CreatedAt)
		taskBuild.SetUpdatedAt(i.UpdatedAt)
		tasks = append(tasks, *taskBuild.Build())
	}

	return responseBuilder.SetData(tasks).Build(), nil
}

// UpdateTask implements Task.
func (t *task) UpdateTask(ctx *gin.Context, req *requests.UpdateTask) (response *responses.ViewEntity[*entities.Task], err error) {
	authUser := ctx.Value("authorization_payload").(*entities.User)
	responseBuilder := responses.NewViewEntityBuilder[*entities.Task]()
	if req.Validate() != nil {
		return responseBuilder.Build(), req.Validate()
	}

	// checkName := t.repo.GetTask().GetByName(ctx, req.Title)
	// if len(checkName.Value) != 0 {
	// 	return responseBuilder.Build(), errors.ErrDataAlready(req.Title)
	// }

	ids, _ := strconv.Atoi(req.ID)
	userID, _ := strconv.Atoi(authUser.ID)

	mBuild := models.NewTaskBuilder()
	mBuild.SetId(ids)
	mBuild.SetTitle(req.Title)
	mBuild.SetDescription(req.Description)
	mBuild.SetStatus(req.Status)
	mBuild.SetUserID(userID)

	if err = t.repo.GetTask().UpdateTask(ctx, req.ID, mBuild.Build()); err != nil {
		return responseBuilder.Build(), err
	}

	resultTask := t.repo.GetTask().GetByID(ctx, req.ID)
	if err = resultTask.Error; err != nil {
		return responseBuilder.Build(), err
	}

	entitiesTask := entities.NewTaskBuilder()
	entitiesTask.SetId(strconv.Itoa(resultTask.Value.ID))
	entitiesTask.SetTitle(resultTask.Value.Title)
	entitiesTask.SetDescription(resultTask.Value.Description)
	entitiesTask.SetStatus(resultTask.Value.Status)
	entitiesTask.SetCreatedAt(resultTask.Value.CreatedAt)
	entitiesTask.SetUpdatedAt(resultTask.Value.UpdatedAt)
	return responseBuilder.SetData(entitiesTask.Build()).Build(), nil
}

// ViewTask implements Task.
func (t *task) ViewTask(ctx *gin.Context, req *requests.EntityId) (response *responses.ViewEntity[*entities.Task], err error) {
	responseBuilder := responses.NewViewEntityBuilder[*entities.Task]()
	if err = req.Validate(); err != nil {
		return responseBuilder.Build(), err
	}
	mBuild := t.repo.GetTask().GetByID(ctx, strconv.Itoa(req.Id.(int)))
	if mBuild.Error != nil {
		return responseBuilder.Build(), mBuild.Error
	}

	entitiesTask := entities.NewTaskBuilder()
	entitiesTask.SetId(strconv.Itoa(mBuild.Value.ID))
	entitiesTask.SetTitle(mBuild.Value.Title)
	entitiesTask.SetDescription(mBuild.Value.Description)
	entitiesTask.SetStatus(mBuild.Value.Status)
	entitiesTask.SetCreatedAt(mBuild.Value.CreatedAt)
	entitiesTask.SetUpdatedAt(mBuild.Value.UpdatedAt)

	return responseBuilder.SetData(entitiesTask.Build()).Build(), nil
}

func NewTask(logger *zap.Logger, repo repositories.Repositories) Task {
	a := new(task)
	a.repo = repo
	a.logger = logger
	return a
}
