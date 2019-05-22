package main

import (
	"fmt"
	"log"
	"net/http"

	"goserver/controllers"
	"goserver/db"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

const dbAddress = "10.16.22.198"
//const dbAddress = "localhost"

func ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	count, _ := r.URL.Query()["count"]

	n := "BLA"

	if len(count) > 0 {
		n = count[0]
	}

	fmt.Printf("ping! param:%v\n", n)
	fmt.Fprintf(w, "ping! param:%v\n", n)
}

func main() {
	fmt.Println("Starting server 0.001 . . .")

	r := httprouter.New()

	session, err := mgo.Dial(dbAddress)
	defer session.Close()

	if err != nil {
		fmt.Printf("Can not start server because of no connection to database:%v \n", err)
		return
	}

	da := db.NewDataAccess(session)

	rc := controllers.NewRegistrationController(da)
	lc := controllers.NewLoginController(da)

	r.POST("/login", lc.Login)

	r.POST("/logout", lc.Authorize(lc.Logout))

	r.POST("/register", rc.Register)

	r.GET("/ping", ping)

	port := "8083"
	fmt.Printf("I'm listening on %v . . .\n", port)
	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", r)
	err = http.ListenAndServe(":"+port, r)

	log.Fatal(err)
}
