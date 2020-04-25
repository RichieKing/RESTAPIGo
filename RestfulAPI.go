package main

//Hello and thank you for your time and consideration to look over this Go file!
//This Go is about RESTful APIs, with the assistance of an application called Postman to ensure that the REST commands worked properly
//So! This program is a RESTful API that works with 4 commands: GET, POST, PUT, and DELETE, all while running on a local server; localhost:8080/dictionary.
//To run and excercise the program, I have included a mock JSON library with 3 entries that have an ID and a Message, the JSON entries can be seen in the "func Main()"
//While going over this file, you will find brief descriptions of each function and at the end, how to run each command (GET, POST, DELETE, and PUT) in Postman to
//ensure that each command is working properly.
//The Website to download Postman is as follows: https://www.postman.com/downloads/ , personally I chose the windows-64bit download to run the file.
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
	rout.HandleFunc("/dictionary", getfun).Methods("GET")            //Creates a GET request of the JSON library to show the entries inside
	rout.HandleFunc("/dictionary", postfun).Methods("POST")          //While the server runs, creates a new entry in the JSON library with an ID and Message
	rout.HandleFunc("/dictionary/{id}", putfun).Methods("PUT")       //Overrides an existing Message in the entries, it targets a message by the ID as seen on the end of the URL
	rout.HandleFunc("/dictionary/{id}", deletefun).Methods("DELETE") //Deletes an ID and Message from the JSON library, it targets the message by the ID on the end of the URL

	//Start up the localhost server on port 8080
	http.ListenAndServe(":8080", rout)
}

// INSTRUCTIONS \\
// Before attempting the commands, first open Postman, run the file to start the server, and use localhose:8080/dictionary as the URL to start.
// You can either hit the "+" to open another tab which will be to make requests, or hit "Create Request" to open a tab

//GET - Postman has a GET command as the default, type in the URL for the local host which is: localhost:8080/dictionary
//		Hit "Send" and the command will bring back the 3 entries in the library.

//POST - Click on GET and scroll to find the POST command next to the URL bar, then go to "Headers" and under "Key" type "Content-Type" which should also be an option
//		in the dropdown menu once you start typing, then under "Value" start typing "application/json".
//		Once those are in place, go to "Body" and click on the "raw" option and type in { "ID": 10, "Message": "Hello Again!" }, hit Send and it will display
//		the message at the bottom of Postman, use a GET command on the same URL and the new message will appear with the other 3 entries.

//PUT - Change the request again and find the option "PUT", this will ONLY work if there is an existing Message you want changed in the entries.
//		You will have to change the URL slightly to include an ID you want targetted, for the excercise, we change ID 3 so the URL should now look like:
//		localhost:8080/dictionary/3
//		Just as directions in POST, go to Headers, type in "Key" Content-Type, and "Value": application/json, go to the "Body", select "raw", and type:
//		{ "Message": "New Message"} , hit send, request a GET command on the Localhost:8080/dictionary and the message of ID 3 should now be changed.

//DELETE -	Request a "DELETE" command from the drop down menu, change the URL by selecting which ID you want to delete, let's delete ID 10, so the URL must be:
//		localhost:8080/dictionary/10 , hit "Send" and then request a "GET" command on localhost:8080/dictionary and notice that the ID 10 and its' message
//		has been deleted.

// 		Thank you for your time and consideration by going over this file! I had fun learning about Golang and how it works from an API perspective and I hope that
//		I have satisfied the requirements set out by this Coding Challenge!
//		- Richard Wong
