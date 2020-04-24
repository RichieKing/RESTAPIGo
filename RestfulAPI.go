package main

//Hello! This Go is about RESTful APIs, unfortunately I had no idea how to run this entirely in file so I used Postman to make sure that the REST commands worked properly
//So! This program is a simple RESTful API that works with 4 commands: GET, POST, PUT, and DELETE, all while running on a local server; localhost8080/dictionary.
//This is where the JSON library is and the mock data that has been put into it to test and see that the code is working properly.
//While going over this file, you will find brief descriptions of each function and what they are doing.
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//Constructing the JSON library called "dictionary" to allow input
type dictionary struct {
	ID      string `json:"id"`
	Message string `json:"Message"`
}

//Splits the dictionary to have an undefined amount of space/memory
var dictionary_content []dictionary

//Getfun is meant to bring back all the data stored at the address, in this case, it will be the localhost:8080/dictionary
func getfun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dictionary_content) //This is found and every function and makes it so that the JSON ID/Message can be read from the splitting dictionary
}

//Postfun is able to take in information and post the data onto the RESTful API
//The point of this function is to be able input an ID and Message of your own through Postman.
func postfun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dictionary_contents dictionary
	_ = json.NewDecoder(r.Body).Decode(&dictionary_contents)
	dictionary_content = append(dictionary_content, dictionary_contents)
	json.NewEncoder(w).Encode(dictionary_contents)
}

//Putfun updates an EXISTING dictionary data, you do NOT change the ID, but the message can be changed.
func putfun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range dictionary_content {
		if item.ID == param["id"] {
			dictionary_content = append(dictionary_content[:index], dictionary_content[index+1:]...)
			var dictionary_contents dictionary
			_ = json.NewDecoder(r.Body).Decode(&dictionary_contents) //Decode takes away the existing Message on a targetted ID, making it ready to be filled by the new message
			dictionary_contents.ID = param["id"]
			dictionary_content = append(dictionary_content, dictionary_contents)
			json.NewEncoder(w).Encode(dictionary_contents)
			return
		}
	}
}

//Deletefun deletes an entry in the JSON dictionary library by indicating what the ID of a message is.
func deletefun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range dictionary_content {
		if item.ID == param["id"] {
			dictionary_content = append(dictionary_content[:index], dictionary_content[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(dictionary_content)
}

//Where the magic happens
func main() {
	//init router
	rout := mux.NewRouter()
	//Dummy data to get started and show that the ID and Message are together in Postman and can be manipulated from there.
	dictionary_content = append(dictionary_content, dictionary{ID: "1", Message: "Hello World"})
	dictionary_content = append(dictionary_content, dictionary{ID: "2", Message: "Goodbye World"})
	dictionary_content = append(dictionary_content, dictionary{ID: "3", Message: "Dark Souls 2"})
	//Route Handles for GET,POST,PUT,DELETE
	//The reason that PUT and DELETE have "{id}" attached to them is because they are targetting existing entries in the dummy data put in
	//GET only brings back the existing data and POST will be able to create its' own ID and Message that DELETE and PUT will then have a target ID.
	rout.HandleFunc("/dictionary", getfun).Methods("GET")
	rout.HandleFunc("/dictionary", postfun).Methods("POST")
	rout.HandleFunc("/dictionary/{id}", putfun).Methods("PUT")
	rout.HandleFunc("/dictionary/{id}", deletefun).Methods("DELETE")

	//Start up the localhost server on port 8080
	http.ListenAndServe(":8080", rout)
}
