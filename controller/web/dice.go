package web

import (
	"net/http"
	"fmt"
	"html/template"
)

var pr = fmt.Println

func DicePlay(w http.ResponseWriter, r *http.Request){


	pr("Dice Play!")

	indexPage, _ := template.New("webpage").Parse("view/dice/main.html")
	indexPage.Execute(w, nil)

}


