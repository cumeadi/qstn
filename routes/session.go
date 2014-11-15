package routes

import (
	"github.com/daryl/sketchy-api/models"
	"github.com/daryl/sketchy-api/utils"
	"github.com/daryl/sketchy-api/utils/session"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func sessionGet(w http.ResponseWriter, r *http.Request) {
	s := session.Start(w, r)
	t := s.Get("token")

	if t == nil {
		utils.JSON(w, nil)
		return
	}

	var user *models.User
	models.Find(user, bson.M{
		"token": t,
	}).One(&user)

	utils.JSON(w, user)
}
