package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

type Users struct {
	gorm.Model

	Name     string
	Password string
	Email    string
}

func InitialMigration() {

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	} else {
		fmt.Println("Connected successfully")
	}
	defer db.Close()
	db.AutoMigrate(&Users{})
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helloworld")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/signup/googleSignup/{name}/{password}/{email}", signupByGoogle).Methods("POST")
	myRouter.HandleFunc("/signup/{email}/{password}", signup).Methods("POST")
	myRouter.HandleFunc("/googleLogin/{name}/{email}", loginByGoogle).Methods("POST")
	myRouter.HandleFunc("/login/{email}", login).Methods("POST")
	myRouter.HandleFunc("/mail/{email}/{rand}", mail).Methods("POST")
	myRouter.HandleFunc("/UpdatePassword/{email}/{password}", PasswordChange).Methods("POST")
	myRouter.HandleFunc("/otp/{otp}/{number}", otpVerify).Methods("POST")
	log.Fatal(http.ListenAndServe(":8100", cors.Default().Handler(myRouter)))
}

func main() {
	InitialMigration()
	handleRequests()
}
