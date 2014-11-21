package routes

import (
	"github.com/daryl/skatchy/utils"
	"net/http"
)

func sessionGet(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, nil)
}
