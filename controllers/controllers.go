package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, GetView("index"))
}

func ReactPage(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, GetView("reactPage"))
}
