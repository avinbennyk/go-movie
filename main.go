package main

import{
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
}

type movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	Firstname string `json:"fname"`
	Lastname string `json:"lname"`
}

var movies []Movie



