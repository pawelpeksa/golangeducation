package main

import (
	"./controllers"
	"encoding/json"
	"fmt"
	"goserver/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func basicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ping!\n")
}

func protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "I'm protected here!\n")
}

func register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Register me please!\n")
	decoder := json.NewDecoder(r.Body)
	var p models.Profile
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%v\n", p.Username)
	fmt.Printf("\n\n%v\n", p.Password)
	fmt.Printf("\n\n%v\n", p.Email)
}

func main() {
	fmt.Println("I'm working 0.1")

	r := httprouter.New()

	rc := controllers.RegistrationController{}
	rc.PrintMePlease()

	r.POST("/login", login)

	r.POST("/logout", logout)

	r.POST("/register", register)

	r.GET("/ping", ping)

	r.GET("/protected", basicAuth(protected, "testingo1", "testingo1"))

	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", r)
	err := http.ListenAndServe(":8083", r)

	log.Fatal(err)
}
