package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rahmat412/go-microservice-template/internal/dto"
	"github.com/rahmat412/go-microservice-template/internal/service"
	"github.com/rahmat412/go-toolbox/httputil"
	"github.com/rahmat412/go-toolbox/logging"
)

type UserHandler struct {
	logger    *logging.Logger
	svc       service.UserServiceProvider
	validator *validator.Validate
}

func NewUserHandler(logger *logging.Logger,
	svc service.UserServiceProvider, validator *validator.Validate) UserHandler {
	return UserHandler{
		logger:    logger,
		svc:       svc,
		validator: validator,
	}
}

func (u *UserHandler) RegisterRoutes(router chi.Router) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", u.RegisterUserHandler)
		r.Get("/{id}", u.GetUserByIDHandler)
		r.Put("/{id}", u.UpdateUserHandler)
		r.Delete("/{id}", u.DeleteUserHandler)
	})
}

func (u UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteBadRequestResponse(w, "JSON can't be decoded", nil)
		return
	}

	if err := u.validator.Struct(req); err != nil {
		httputil.HandleError(w, u.logger, err)
		return
	}

	response, err := u.svc.CreateUser(r.Context(), req)
	if err != nil {
		httputil.HandleError(w, u.logger, err)
		return
	}

	httputil.WriteSuccessResponse(w, response, "User created successfully")
}

func (u UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		httputil.WriteBadRequestResponse(w, "User ID is required", nil)
		return
	}

	// Convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		httputil.WriteBadRequestResponse(w, "Invalid user ID", nil)
		return
	}

	response, err := u.svc.GetUserByID(r.Context(), userID)
	if err != nil {
		httputil.HandleError(w, u.logger, err)
		return
	}

	httputil.WriteSuccessResponse(w, response, "User retrieved successfully")
}

func (u UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		httputil.WriteBadRequestResponse(w, "User ID is required", nil)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		httputil.WriteBadRequestResponse(w, "Invalid user ID", nil)
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteBadRequestResponse(w, "JSON can't be decoded", nil)
		return
	}

	if err := u.validator.Struct(req); err != nil {
		httputil.HandleError(w, u.logger, err)
		return
	}

	response, err := u.svc.UpdateUser(r.Context(), userID, &req)
	if err != nil {
		httputil.HandleError(w, u.logger, err)
		return
	}

	httputil.WriteSuccessResponse(w, response, "User updated successfully")
}

func (u UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		httputil.WriteBadRequestResponse(w, "User ID is required", nil)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		httputil.WriteBadRequestResponse(w, "Invalid user ID", nil)
		return
	}

	err = u.svc.DeleteUser(r.Context(), userID)
	if err != nil {
		httputil.HandleError(w, u.logger, err)
		return
	}

	httputil.WriteSuccessResponse(w, nil, "User deleted successfully")
}
