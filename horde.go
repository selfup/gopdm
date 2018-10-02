package horde

import (
	"net/http"
)

// Boot starts up the Horde and returns the horde Manager
func Boot() Manager {
	manager := Manager{}

	go func() {
		http.ListenAndServe(":9742", &hordeHandler{})
	}()

	return manager
}

type hordeHandler struct {
}

func (m *hordeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Horde is Listening on 9742"))
}
