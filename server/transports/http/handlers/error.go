package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/endpoints/responses"
)

// EncodeError encode errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	responseBuilder := responses.NewErrorBuilder()

	if strings.Contains(err.Error(), errors.ERR_NOT_AUTHORIZED) {
		w.WriteHeader(http.StatusUnauthorized)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_INVALID_KEY) {
		w.WriteHeader(http.StatusBadRequest)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_NOT_FOUND) {
		w.WriteHeader(http.StatusNotFound)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_ALREADY_EXISTS) {
		w.WriteHeader(http.StatusConflict)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(err.Error(), errors.ERR_TOKEN) {
		w.WriteHeader(http.StatusUnauthorized)
		responseBuilder.SetError(err)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	errors.HandlerReturnedVoid(err)
	w.WriteHeader(http.StatusInternalServerError)
	responseBuilder.SetError(err)
	json.NewEncoder(w).Encode(responseBuilder.Build())
}
