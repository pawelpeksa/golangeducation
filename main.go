package main

import (
	"./controllers"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goserver/models"
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

	session, err := mgo.Dial("localhost")

	if err != nil {
		fmt.Printf("\n\nCouldn't open session with mongo with error:%v\n", err)
		return
	}

	c := session.DB("db").C("users")

	count, err := c.Find(bson.M{"username": p.Username}).Limit(1).Count()

	fmt.Printf("\nTEST: %v", count)
	fmt.Printf("\nEND")

	err = c.Insert(p)

	if err != nil {
		if mgo.IsDup(err) {
			fmt.Printf("\nYou're trying to give me duplicate:%v\n", p)
		}
	} else {
		fmt.Printf("\n Looks like success:%v\n", p)
	}

	fmt.Printf("\n\n%v\n", p.Username)
	fmt.Printf("\n%v\n", p.Password)
	fmt.Printf("\n%v\n\n", p.Email)

	session.Close()
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
