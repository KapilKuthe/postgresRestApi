// handler/handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"postgresRestApi/database"
	"strconv"

	"github.com/gorilla/mux"
)

// gets the entire user list
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetUsers()
	if err != nil {
		// if fetching users fails
		RespondJSON(w, http.StatusInternalServerError, nil, "Failed to fetch users", err)
		return
	}

	RespondJSON(w, http.StatusOK, users, "Users fetched successfully", nil)
}

// getting the user for single user based on id
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		// if the request id is invalid
		RespondJSON(w, http.StatusBadRequest, nil, "Invalid user ID", err)
		return
	}

	user, err := database.GetUser(uint(userID))
	if err != nil {
		// if DB fails to retrive the user
		RespondJSON(w, http.StatusInternalServerError, nil, "Failed to fetch user", err)
		return
	}

	RespondJSON(w, http.StatusOK, user, "User fetched successfully", nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Decode the request body to get the new user information
	var newUser database.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// if the request body is invalid
		RespondJSON(w, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	// Create the user in DB
	createdUser, err := database.CreateUser(newUser)
	if err != nil {
		// if creating the user fails from DB
		RespondJSON(w, http.StatusInternalServerError, nil, "Failed to create user", err)
		return
	}

	RespondJSON(w, http.StatusCreated, createdUser, "User created successfully", nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		// if the request id is invalid
		RespondJSON(w, http.StatusBadRequest, nil, "Invalid user ID", err)
		return
	}

	var updatedUser database.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		// if the request body is invalid
		RespondJSON(w, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	updatedUser.ID = uint(userID)
	// Update the user in DB
	updatedUser, err = database.UpdateUser(updatedUser)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, nil, "Failed to update user", err)
		return
	}

	RespondJSON(w, http.StatusOK, updatedUser, "User updated successfully", nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		// if the request id is invalid
		RespondJSON(w, http.StatusBadRequest, nil, "Invalid user ID", err)
		return
	}

	// Delete the user in DB
	err = database.DeleteUser(uint(userID))
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, nil, "Failed to delete user", err)
		return
	}

	// if we user http.StatusNoContent no response is retured
	RespondJSON(w, http.StatusOK, nil, "User deleted successfully", nil)
}
