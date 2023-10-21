package controllers

import (
	"encoding/json"
	"net/http"

	"authservice/models"

	"go.uber.org/zap"
)

// RegisterController is the Register route handler
type RegisterController struct {
	logger *zap.Logger
}

type RegisterRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

// NewRegisterController returns a fresh Register controller
func NewRegisterController(logger *zap.Logger) *RegisterController {
	return &RegisterController{
		logger: logger,
	}
}

// adds the user to the database of users
// Register - register a new user
// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param input body RegisterRequest true "User register info"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (ctrl *RegisterController) RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	// Create a variable to hold the request data
	var registerRequest RegisterRequest

	// Decode the JSON request body into the RegisterRequest struct
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		ctrl.logger.Error("Error decoding request body", zap.Error(err))
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request body"))
		return
	}

	// Now you can access the data from registerRequest
	email := registerRequest.Email
	passwordHash := registerRequest.PasswordHash

	// You can also add validation for other fields if needed

	// Validate and then add the user
	check := models.AddUserObject(email, "", passwordHash, "", 0)

	// If check is false, it means the user already exists
	if !check {
		ctrl.logger.Warn("User already exists", zap.String("email", email))
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("Email or Username already exists"))
		return
	}

	ctrl.logger.Info("User created", zap.String("email", email))
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Created"))
}
