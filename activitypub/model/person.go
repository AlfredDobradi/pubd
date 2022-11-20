package model

import (
	"fmt"

	"github.com/alfreddobradi/activitypub/config"
)

/*
{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "https://w3id.org/security/v1"
  ],
  "id": "https://mastodon.social/users/Gargron",
  "type": "Person",
  "publicKey": {
    "id": "https://mastodon.social/users/Gargron#main-key",
    "owner": "https://mastodon.social/users/Gargron",
    "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvXc4vkECU2/CeuSo1wtn\nFoim94Ne1jBMYxTZ9wm2YTdJq1oiZKif06I2fOqDzY/4q/S9uccrE9Bkajv1dnkO\nVm31QjWlhVpSKynVxEWjVBO5Ienue8gND0xvHIuXf87o61poqjEoepvsQFElA5ym\novljWGSA/jpj7ozygUZhCXtaS2W5AD5tnBQUpcO0lhItYPYTjnmzcc4y2NbJV8hz\n2s2G8qKv8fyimE23gY1XrPJg+cRF+g4PqFXujjlJ7MihD9oqtLGxbu7o1cifTn3x\nBfIdPythWu5b4cujNsB3m3awJjVmx+MHQ9SugkSIYXV0Ina77cTNS0M2PYiH1PFR\nTwIDAQAB\n-----END PUBLIC KEY-----\n"
  }
}
*/

type Person struct {
	Context   []string  `json:"@context"`
	ID        string    `json:"id"`
	Account   string    `json:"-"`
	Type      string    `json:"type"`
	PublicKey PublicKey `json:"publicKey"`
}

type PublicKey struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Key   string `json:"publicKeyPem"`
}

func NewPerson(username string, pem string) *Person {
	id := fmt.Sprintf("https://%s/~%s", config.BaseURL(), username)
	return &Person{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		ID:      id,
		Account: fmt.Sprintf("%s@%s", username, config.BaseURL()),
		Type:    "Person",
		PublicKey: PublicKey{
			ID:    fmt.Sprintf("%s#main-key", id),
			Owner: id,
			Key:   pem,
		},
	}
}
