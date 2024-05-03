package RestAPI_go

import "net/http"

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/data/all", returnAllData)
	http.HandleFunc("/data/single", returnSingleData)
	http.HandleFunc("/data", createNewData)
	http.HandleFunc("/data/delete", deleteData)
}
