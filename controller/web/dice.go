package web

import (
	"net/http"
	"fmt"
)

var pr = fmt.Println

func DicePlay(w http.ResponseWriter, r *http.Request){
	pr("Dice Play!")
}


