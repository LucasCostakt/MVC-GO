package view

import (
	"net/http"

	"github.com/MVC/controllers"
)

func BootAplication() {
	http.HandleFunc("/employees", controllers.GetEmployee)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
