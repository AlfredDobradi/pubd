package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alfreddobradi/activitypub/activitypub/model"
	"github.com/alfreddobradi/activitypub/database/memory"
)

const (
	RelProfile string = "http://webfinger.net/rel/profile-page"
	RelSelf    string = "self"

	TypeHtml     string = "text/html"
	TypeActivity string = "application/activity+json"
)

func WebfingerHandler(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Query().Get("resource")
	if resource == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	query := strings.TrimPrefix(resource, "acct:")
	user, err := memory.GetStore().Users.FindByAccount(query)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	response := WebfingerResponse{
		Subject: resource,
		Aliases: []string{
			user.ID,
		},
		Links: []model.Link{
			{
				Rel:  RelProfile,
				Type: "text/html",
				Href: user.ID,
			},
			{
				Rel:  RelSelf,
				Type: "application/activity+json",
				Href: user.ID,
			},
		},
	}

	raw, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(raw) // nolint
}
