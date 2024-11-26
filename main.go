package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// user struct
type user struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type getParam struct {
	Search string `json:"search"`
	SortBy string `json:"sort_by"`
	Order  string `json:"order"`
}

func getAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// panic("Error")
	user := []user{
		{
			Name: "John",
			Age:  25,
		},
		{
			Name: "Doe",
			Age:  30,
		},
		{
			Name: "Smith",
			Age:  35,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func getUserDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user := user{
		Id:   id,
		Name: "John",
		Age:  25,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var user user
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func editUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	decoder := json.NewDecoder(r.Body)
	var user user
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}

	user.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func getAllUsersWithParam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var param getParam
	err := decoder.Decode(&param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error : %v", err)
		return
	}

	query := "SELECT * FROM users WHERE name like '%" + param.Search + "%' ORDER BY " + param.SortBy + " " + param.Order
	fmt.Println(query)

	userList := []user{
		{
			Name: "Budi",
			Age:  12,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)
}

func main() {
	router := httprouter.New()
	router.POST("/all", getAllUsersWithParam)
	router.GET("/all", getAllUsers)
	router.GET("/detail/:id", getUserDetail)
	router.POST("/add", addUser)
	router.POST("/update/:id", editUser)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, c interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Got panic error : %v", c)
	}

	http.ListenAndServe(":8080", router)
}
