package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goldfish-api/internal/core/ports"
	"net/http"
	"strings"
)

type HTTPUserHandler struct {
	userService ports.UserService
	router *mux.Router
}

// NewHTTPUserHandler creates a new instance of the user handler
func NewHTTPUserHandler(router *mux.Router, userService ports.UserService) *HTTPUserHandler {
	userHandler := HTTPUserHandler{
		userService: userService,
		router:      router,
	}

	// set routes of user handler
	userHandler.init()

	return &userHandler
}

// init initializes the handler and sets up the routes
func (handler *HTTPUserHandler) init() {
	handler.router.HandleFunc("/register", handler.Register)
}

func (handler *HTTPUserHandler) Register(writer http.ResponseWriter, request *http.Request) {
	var requestBody UserRegisterRequest

	// parse and validate request body
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil || !requestBody.IsValid() {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	// save user
	user := requestBody.ToUser()
	jwt, err := handler.userService.Create(&user)

	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			http.Error(writer, DUPLICATE_EMAIL, http.StatusBadRequest)
			return
		}

		http.Error(writer, INTERNAL_SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	fmt.Fprint(writer, "Registered... :D" + jwt)
}