package main

import (
	"fmt"
	"net/http"

	"github.com/AlbanAGBOTA/applicationWebGO_1/internal/handlers"
	//"github.com/AlbanAGBOTA/applicationWebGO_1/internal/handlers"
)

const port = "8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Serveur lancé sur http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)

	//Importer le nom le projet
	//applicationWebGO_1.Config()

}
