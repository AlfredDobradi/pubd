package memory

import (
	"fmt"
	"sync"

	"github.com/alfreddobradi/activitypub/activitypub/model"
	"github.com/alfreddobradi/activitypub/config"
)

var store *Store

type Store struct {
	Users *UserStore
}

type UserStore struct {
	mx *sync.Mutex

	users []*model.Person
}

func (u *UserStore) Add(p *model.Person) {
	u.mx.Lock()
	defer u.mx.Unlock()

	update := false
	for i, user := range u.users {
		if user.ID == p.ID {
			u.users[i] = p
			update = true
			break
		}
	}
	if !update {
		u.users = append(u.users, p)
	}
}

func (u *UserStore) FindByAccount(resource string) (*model.Person, error) {
	u.mx.Lock()
	defer u.mx.Unlock()

	for _, user := range u.users {
		if user.Account == resource {
			return user, nil
		}
	}

	return nil, fmt.Errorf("Not found")
}

func GetStore() *Store {
	if store == nil {
		userStore := &UserStore{
			mx:    &sync.Mutex{},
			users: make([]*model.Person, 0),
		}

		store = &Store{
			Users: userStore,
		}

		store.Users.Add(model.NewPerson(config.User(), config.KeyPem()))
	}

	return store
}
