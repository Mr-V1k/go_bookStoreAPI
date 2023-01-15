package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"bookstore.com/go-bookstore/pkg/models"
	"bookstore.com/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	//gets the books from the db
	newBooks := models.GetAllBooks()
	//converts to json
	res, _ := json.Marshal(newBooks)
	//setting headers
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	//sending response back
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//to get the id
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	//convert to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}
	//look for the books with the id
	bookdetails, _ := models.GetBookById(ID)
	//convert to json
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	//send json
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	//get the model
	CreateBook := &models.Book{}
	//parse the body
	utils.ParseBody(r, CreateBook)
	//create the book
	b := CreateBook.CreateBook()
	//convert to json
	res, _ := json.Marshal(b)
	//send status
	w.WriteHeader(http.StatusOK)
	//setting the content type
	w.Header().Set("Content-Type", "pkglication/json")
	//response
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//get id
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	//parse id to int
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error Parsing!")
	}
	//look for the book with the id and delete
	deletedBook := models.DeleteBook(id)
	//convert ot json
	dBook, _ := json.Marshal(deletedBook)
	//prompt the it is completed
	w.WriteHeader(http.StatusOK)
	//send the message
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(dBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//get the book model
	var updateBook = &models.Book{}
	//parse it
	utils.ParseBody(r, updateBook)
	//get the id
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	//parse it into int
	id, err := strconv.ParseInt(bookId, 0, 0)
	//update he details
	if err != nil {
		fmt.Println("Error parsing to int")
	}
	bookdetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookdetails.Publication = updateBook.Publication
	}
	//save it to db
	db.Save(&bookdetails)
	//convert to json
	res, _ := json.Marshal(bookdetails)
	//set the headers and status
	w.Header().Set("Content_Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	//send the updated book
	w.Write(res)
}
