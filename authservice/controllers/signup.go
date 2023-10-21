package controllers

import (
	"encoding/json"
	"net/http"

	"authservice/models"

	"go.uber.org/zap"
)

// SignupController is the Signup route handler
type SignupController struct {
	logger *zap.Logger
}
type SignupRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
}
type SignupResponse struct {
	Token string `json:"token"`
}

// NewSignupController returns a frsh Signup controller

func NewSignupController(logger *zap.Logger) *SignupController {
	return &SignupController{
		logger: logger,
	}
}

// adds the user to the database of users
// Signup - register a new user
// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param input body SignupRequest true "User signup info"
// @Success 200 {object} SignupResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signup [post]
func (ctrl *SignupController) SignupHandler(rw http.ResponseWriter, r *http.Request) {
	// Create a variable to hold the request data
	var signupRequest SignupRequest

	// Decode the JSON request body into the SignupRequest struct
	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		ctrl.logger.Error("Error decoding request body", zap.Error(err))
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request body"))
		return
	}

	// Now you can access the data from signupRequest
	email := signupRequest.Email
	passwordHash := signupRequest.PasswordHash

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
