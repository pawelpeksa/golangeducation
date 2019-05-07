package controllers

import (
	"fmt"
	"goserver/db"
)

type RegistrationController struct {
	da db.DataAccess
}

func (rc RegistrationController) PrintMePlease() {
	fmt.Println("co za zycie")
	fmt.Printf("%v", rc.da)
}
