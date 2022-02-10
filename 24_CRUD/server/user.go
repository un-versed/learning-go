package server

import (
	"CRUD/db"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading request body"))
		return
	}

	var user user

	if error = json.Unmarshal(body, &user); error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading request body"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error inserting user"))
		return
	}

	defer statement.Close()

	insertedRow, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error inserting user"))
		return
	}

	id, error := insertedRow.LastInsertId()
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error getting user id"))
		return
	}

	user.ID = id

	userJSON, error := json.Marshal(user)
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in JSON parse"))
		return
	}

	rw.Write(userJSON)
}

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	db, error := db.Connect()
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	rows, error := db.Query("SELECT * FROM users")
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error executing query"))
		return
	}

	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if error := rows.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Error scanning user"))
			return
		}

		users = append(users, user)
	}

	usersJSON, error := json.Marshal(users)
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in JSON parse"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(usersJSON)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseInt(params["id"], 10, 64)
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error parsing ID"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	row, error := db.Query("SELECT * FROM users WHERE ID = ?", ID)
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error executing query"))
		return
	}

	defer row.Close()

	var user user
	if row.Next() {
		if error := row.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Error scanning user"))
			return
		}
	}

	usersJSON, error := json.Marshal(user)
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error in JSON parse"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(usersJSON)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading request body"))
		return
	}

	ID, error := strconv.ParseInt(params["id"], 10, 64)
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error parsing ID"))
		return
	}

	var user user

	if error = json.Unmarshal(body, &user); error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading request body"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error updating user"))
		return
	}

	defer statement.Close()

	if _, error := statement.Exec(user.Name, user.Email, ID); error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error updating user"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseInt(params["id"], 10, 64)
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error parsing ID"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("DELETE FROM users WHERE id = ?")
	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error deleting user"))
		return
	}

	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error deleting user"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
