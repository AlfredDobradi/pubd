package activitypub

import (
	"net/http"

	"github.com/alfreddobradi/activitypub/activitypub/model"
	"github.com/alfreddobradi/activitypub/config"
)

func Sign(r *http.Request) error {
	return nil
}

var owner *model.Person

func Owner() *model.Person {
	if owner == nil {
		owner = model.NewPerson(config.User(), config.KeyPem())
	}

	return owner
}
