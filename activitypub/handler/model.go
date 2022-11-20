package handler

import "github.com/alfreddobradi/activitypub/activitypub/model"

type WebfingerResponse struct {
	Subject string   `json:"subject"`
	Aliases []string `json:"aliases"`
	Links   []model.Link
	/*
	    {
	  "subject": "acct:axdx@boltcutter.network",
	  "aliases": [
	    "https://boltcutter.network/@axdx",
	    "https://boltcutter.network/users/axdx"
	  ],
	  "links": [
	    {
	      "rel": "http://webfinger.net/rel/profile-page",
	      "type": "text/html",
	      "href": "https://boltcutter.network/@axdx"
	    },
	    {
	      "rel": "self",
	      "type": "application/activity+json",
	      "href": "https://boltcutter.network/users/axdx"
	    },
	    {
	      "rel": "http://ostatus.org/schema/1.0/subscribe",
	      "template": "https://boltcutter.network/authorize_interaction?uri={uri}"
	    }
	  ]
	}
	*/
}
