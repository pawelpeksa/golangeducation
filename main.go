package main

import (
	"./controllers"
	"./db"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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

func ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	count, _ := r.URL.Query()["count"]

	n := "BLA"

	if len(count) > 0 {
		n = count[0]
	}

	fmt.Printf("ping! param:%v\n", n)
	fmt.Fprintf(w, "ping! param:%v\n", n)
}

func protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "I'm protected here!\n")
}

func main() {
	fmt.Println("I'm working 0.1")

	r := httprouter.New()

	da := db.NewDataAccess()
	rc := controllers.NewRegistrationController(da)

	r.POST("/login", login)

	r.POST("/logout", logout)

	r.POST("/register", rc.Register)

	r.GET("/ping", ping)

	r.GET("/protected", basicAuth(protected, "testingo1", "testingo1"))

	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", r)
	err := http.ListenAndServe(":8083", r)

	log.Fatal(err)
}
