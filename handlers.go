package RestAPI_go

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Домашняя страница")
	fmt.Println("homePage")
}

func returnAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllData")
	err := json.NewEncoder(w).Encode(db)
	if err != nil {
		return
	}
}

func returnSingleData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnSingleData")
	id := r.URL.Query().Get("id")

	for _, v := range db {

		if v.Id == id {
			err := json.NewEncoder(w).Encode(v)

			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
	http.Error(w, "Data not found", http.StatusNotFound)
}

func createNewData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var data Data
	if err := json.Unmarshal(reqBody, &data); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	db = append(db, data)

	json.NewEncoder(w).Encode(db)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for index, v := range db {
		if v.Id == id {
			db = append(db[:index], db[index+1:]...)
			fmt.Fprintf(w, "Data with ID:%s sucesfully deleted\n", id)
		}
	}

	http.Error(w, "Data not found", http.StatusNotFound)
}
