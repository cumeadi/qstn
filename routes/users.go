package routes

import (
	"fmt"
	"net/http"
)

func usersGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users!")
}
