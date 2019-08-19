package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"html/template"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/subosito/twilio"
)

func signupByGoogle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	password := vars["password"]
	email := vars["email"]

	db.Create(&Users{Name: name, Password: password, Email: email})

	fmt.Println(w, "New user created successfully")
	fmt.Fprintf(w, "New user created successfully")

}

func signup(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)

	password := vars["password"]
	email := vars["email"]

	db.Create(&Users{Password: password, Email: email})

	fmt.Println(w, "New user created successfully")
	fmt.Fprintf(w, "New user created successfully")

	welcomeMail(email)

}

func loginByGoogle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	var user Users

	name := vars["name"]
	email := vars["email"]

	db.Model(&user).Where("Name= ? AND Email=?", name, email).Find(&user)
	json.NewEncoder(w).Encode(user)

}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	var user Users

	email := vars["email"]

	db.Model(&user).Where("Email=?", email).Find(&user)
	json.NewEncoder(w).Encode(user)

}

func mail(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	email := vars["email"]
	rand := vars["rand"]

	fmt.Println(w, "reason ")
	fmt.Fprintf(w, "reason")
	send(email, rand)
}

func send(email string, rand string) {
	from := "rajuart678@gmail.com"
	pass := "veluprabha21669"
	to := "rajuart678@gmail.com"

	msg := "From:  DReam 11 \n" +
		"To: " + to + "\n" +
		"Subject: Password Change\n\n" +
		"Respected sir/Mam\n\n" +
		"http://localhost:8080/changePass/" + email + "/" + rand

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

func PasswordChange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	var user Users

	email := vars["email"]
	password := vars["password"]

	db.Model(&user).Where("Email=?", email).Update("Password", password)
	json.NewEncoder(w).Encode(user)

}

func otpVerify(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=dream password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)

	otp := vars["otp"]
	to := vars["number"]

	sendSms(otp, to)

}

var (
	AccountSid = "AC7e841768526cb8a12b2fa3a685199cfa"
	AuthToken  = "65187b21d33927bdd438367b93403447"
	From       = "+12055461950"
	To         = "+9177080 22202"
)

func sendSms(otp string, to string) {
	// Initialize twilio Client
	c := twilio.NewClient(AccountSid, AuthToken, nil)

	// Send Message
	params := twilio.MessageParams{
		Body: otp,
	}

	s, resp, err := c.Messages.Send(From, to, params)
	log.Println("Send:", s)
	log.Println("Response:", resp)
	log.Println("Err:", err)
}

/*-------------welcome mail-------------*/
var auth smtp.Auth

func welcomeMail(toEmail string) {
	auth = smtp.PlainAuth("", "rajuart678@gmail.com", "veluprabha21669", "smtp.gmail.com")
	templateData := struct {
		Name string
		URL  string
		Src  string
	}{
		Name: "Raj",
		URL:  "https://Localhost:8080",
		Src:  "./dream.jpeg",
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r := NewRequest([]string{toEmail}, "Welcome to DREAM11", "Hello")

	if err := r.ParseTemplate("template.html", templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}

}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "dhanush@geektrust.in", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
