package horde

import (
	"net/http"
)

// Boot starts up the Horde and returns the horde Manager
func Boot() *Manager {
	manager := new(Manager)

	go func() {
		http.HandleFunc("/", ManagerHandler(manager))

		http.ListenAndServe(":9742", nil)

		defer manager.RemoveSelfFromHorde()
	}()

	return manager
}

// ManagerHandler is the http interface to the horde
func ManagerHandler(manager *Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Horde is listening!"))
	}
}
