// Package dynowebportal handles http requests.
package dynowebportal

import (
	"fmt"
	"net/http"
)

// RunWebPortal runs the dino web portal on the given address argument.
func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	// The line below is blocking code.
	// If successful, the program will stop and listen for incoming requests.
	// If it fails the, the code will no longer block and an error will be returned.
	return http.ListenAndServe(addr, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}
