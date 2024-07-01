package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", streamBackup)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

var streamCount = 0
var lock = sync.Mutex{}

func create() {
	lock.Lock()
	defer lock.Unlock()

	streamCount++
	if streamCount > 1 {
		// already active
		return
	}

	// create snapshot
	// TODO
}

func release() {
	lock.Lock()
	defer lock.Unlock()

	streamCount--
	if streamCount > 0 {
		// still active backup streams
		return
	}

	// remove snapshot
	// TODO
}

func streamBackup(w http.ResponseWriter, r *http.Request) {
	create()
	defer release()
}
