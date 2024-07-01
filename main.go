package main

import "net/http"

func main() {
	http.HandleFunc("/", streamBackup)

	http.ListenAndServe(":8080", nil)
}

func streamBackup(w http.ResponseWriter, r *http.Request) {

}
