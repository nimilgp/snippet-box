package main

import (
		"fmt"
		"net/http"
		"strconv"
		"html/template"
		"log"
)

//handler definitions
func getHome(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "net/http Golang") //for adding additional header details
		w.Header().Add("Author", "nimilgp") 
		w.Header().Add("Program", "Snippet-Box") 
		w.Header().Add("Field-Name", "Can't have spaces in them")

		//slice containing files to build the final home page
		//base template must be 1st
		files := []string{
				"./ui/html/base.tmpl.html",
				"./ui/html/pages/home.tmpl.html",
		}

		ts, err := template.ParseFiles(files...)//pass contents of file as variadic args
		if err != nil {
				log.Print(err.Error()) //log err on server system
				http.Error(w, "Internal Server ERROR(ParseFiles)", http.StatusInternalServerError) //notify user of err
				return
		}

		err = ts.ExecuteTemplate(w, "base", nil)//specifically tells to respond with content 'base', dynamic data -> nil
		if err != nil {
				log.Print(err.Error())
				http.Error(w, "Internal Server ERROR(Execute)", http.StatusInternalServerError)
		}
}


func getSnippetView(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id<1 {
				http.NotFound(w,r)
				return
		}
		fmt.Fprintf(w, "Display Snippet no: %d", id)
}

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create snippet")
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated) //201 created status code
		fmt.Fprintf(w, "POST part of create snippet")
}

