package main

import(
	"fmt"
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

// GET CLASSIFIED CTRL
func GetClassifiedsCtrl (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	enoder := json.NewEncoder(w)

	enoder.Encode(DBGetClassifiedsCtrl ())
} 

// GET ONCE CLASSIFIED CTRL
func GetClassifiedCtrl (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	enoder := json.NewEncoder(w)

	id := mux.Vars(r) // type is: map[string]string

	cId, _ := strconv.Atoi(id["id"])

	enoder.Encode(DBGetClassifiedCtrl (cId))
} 

// POST CLASSIFIED CTRL
func PostClassifiedsCtrl (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	enoder := json.NewEncoder(w)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)
	}

	var newClassified PostClassifed

	json.Unmarshal(body, &newClassified)

	fmt.Println(newClassified)

	enoder.Encode(DBPostClassifiedCtrl(newClassified))
} 

// PUT CLASSIFIED CTRL
func PutClassifiedsCtrl (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	enoder := json.NewEncoder(w)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)
	}

	var changedClassified PutClassified

	json.Unmarshal(body, &changedClassified)

	fmt.Println(changedClassified)

	enoder.Encode(DBPutClassifiedCtrl(changedClassified))
	//enoder.Encode(DBPutClassifiedCtrl(changedClassified))
} 


// DELETE CLASSIFIED CTRL
func DelClassifiedCtrl (w http.ResponseWriter, r *http.Request) {

	type DelClsfdId struct {
		Id uint16	`json:"id"`
	}

	w.Header().Set("Content-Type", "application/json")

	enoder := json.NewEncoder(w)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("ioutil readall error %v", err)
	}

	var id DelClsfdId

	json.Unmarshal(body, &id)

	enoder.Encode(DBDeleteClassifiedCtrl(id.Id))
} 