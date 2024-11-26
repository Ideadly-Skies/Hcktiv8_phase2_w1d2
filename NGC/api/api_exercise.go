package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/go-sql-driver/mysql"
)

type Hero struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"image_url"`
}

type Villain struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageURL string `json:"image_url"`
}

var db *sql.DB

// Initialize Database
func initDB() {
	var err error
	dsn := "root:11111111@tcp(127.0.0.1:3306)/ftgo_phase2_w1d2"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Database connected!")
}

// Get All Heroes
func getAllHeroes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := db.Query("SELECT Id, Name, Universe, Skill, ImageURL FROM Heroes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []Hero
	for rows.Next() {
		var hero Hero
		if err := rows.Scan(&hero.Id, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		heroes = append(heroes, hero)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

// Get All Villains
func getAllVillains(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := db.Query("SELECT Id, Name, Universe, ImageURL FROM Villain")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []Villain
	for rows.Next() {
		var villain Villain
		if err := rows.Scan(&villain.Id, &villain.Name, &villain.Universe, &villain.ImageURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		villains = append(villains, villain)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villains)
}

func API() {
	initDB()

	router := httprouter.New()
	router.GET("/heroes", getAllHeroes)
	router.GET("/villains", getAllVillains)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}