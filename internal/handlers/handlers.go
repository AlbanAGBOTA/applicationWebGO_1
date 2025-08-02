package handlers

import (
	"html/template"
	"net/http"
)

// Accueille
func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Bonjour a tous!")
	renderTemplate(w, "home")
}

// A propos
func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "A propos de moi!")
	renderTemplate(w, "about")
}

// Associe les deux fichier templates
func renderTemplate(w http.ResponseWriter, tmpl string) {
	//recuperer le contenu des fichiers home.page.tmpl et about.page.tmpl
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	//gérer les erreurs
	if err != nil {
		//fmt.Printf("erreur")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}
