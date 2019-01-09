package horde

import (
	"net/http"
)

// Boot starts up the Horde and returns the horde Manager
func Boot() *Manager {
	manager := new(Manager)

	go func() {
		defer manager.RemoveSelfFromHorde()

		http.HandleFunc("/", managerHandler(manager))

		http.ListenAndServe(":9742", nil)

	}()

	return manager
}

// managerHandler is the http interface to the horde
func managerHandler(manager *Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Horde is listening!"))
	}
}
