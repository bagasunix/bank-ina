package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

func decodeByEntityIdEndpoint(c *gin.Context) (request interface{}, err error) {
	id := c.Param("id")
	if id == "" {
		return nil, errors.ErrInvalidAttributes("id")
	}

	// Periksa apakah id adalah tipe int
	intId, err := strconv.Atoi(id)
	if err == nil {
		// Gunakan tipe int
		requestBuilder := requests.NewEntityIdBuilder()
		requestBuilder.SetId(intId)
		return requestBuilder.Build(), nil
	}

	// Periksa apakah id adalah tipe uuid.UUID
	uuidId, err := uuid.FromString(id)
	if err == nil {
		// Gunakan tipe uuid.UUID
		requestBuilder := requests.NewEntityIdBuilder()
		requestBuilder.SetId(uuidId)
		return requestBuilder.Build(), nil
	}

	requestBuilder := requests.NewEntityIdBuilder()
	requestBuilder.SetId(uuidId)
	return requestBuilder.Build(), nil
}
