package controllers

import (
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
	// extra error handling should be done at server side to prevent malicious attacks
	if _, ok := r.Header["Email"]; !ok {
		ctrl.logger.Warn("Email was not found in the header")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email Missing"))
		return
	}
	if _, ok := r.Header["Username"]; !ok {
		ctrl.logger.Warn("Username was not found in the header")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		ctrl.logger.Warn("Passwordhash was not found in the header")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Passwordhash Missing"))
		return
	}
	if _, ok := r.Header["Fullname"]; !ok {
		ctrl.logger.Warn("Fullname was not found in the header")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Fullname Missing"))
		return
	}

	// validate and then add the user
	check := models.AddUserObject(r.Header["Email"][0], r.Header["Username"][0], r.Header["Passwordhash"][0],
		r.Header["Fullname"][0], 0)
	// if false means username already exists
	if !check {
		ctrl.logger.Warn("User already exists", zap.String("email", r.Header["Email"][0]), zap.String("username", r.Header["Username"][0]))
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("Email or Username already exists"))
		return
	}
	ctrl.logger.Info("User created", zap.String("email", r.Header["Email"][0]), zap.String("username", r.Header["Username"][0]))
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Created"))
}
