package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/v1/features/users/domains"
	"server/v1/features/users/services"
	"server/v1/utils"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// CreateUser handles the HTTP request for creating a new user
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domains.User

	body := r.Body
	defer body.Close()

	// Decode the JSON request body into a user
	errDecode := json.NewDecoder(body).Decode(&user)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
		return
	}

	// Call the service to create the user
	createdUser, errCreate := c.userService.CreateUser(user)
	if errCreate != nil {
		http.Error(w, errCreate.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the created user into JSON for the response
	response, errResponse := json.Marshal(createdUser)
	if errResponse != nil {
		http.Error(w, errResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Query())
	urlQuery := r.URL.Query()
	filter := utils.GetQueryFilter(urlQuery)

	users, err := c.userService.GetUsers(filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := c.userService.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
