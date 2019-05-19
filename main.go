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
	fmt.Println("Starting server 0.001 . . .")

	r := httprouter.New()

	da, err := db.NewDataAccess()

	if err != nil {
		fmt.Printf("Can not start server because of no connection to database:%v \n", err)
		return
	}

	rc := controllers.NewRegistrationController(da)
	la := controllers.NewLoginController(da)

	r.POST("/login", la.Login)

	r.POST("/logout", la.Logout)

	r.POST("/register", rc.Register)

	r.GET("/ping", ping)

	r.GET("/protected", basicAuth(protected, "testingo1", "testingo1"))

	port := "8083"
	fmt.Printf("I'm listening on %v . . .\n", port)
	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", r)
	err = http.ListenAndServe(":"+port, r)

	log.Fatal(err)
}
