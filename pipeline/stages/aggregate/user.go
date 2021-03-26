package aggregate

import (
	"fmt"

	"github.com/godoylucase/go-data-pipelines-example/business"
)

type UserFetcher interface {
	Get(ID business.UserID) (business.User, error)
}

func User(done <-chan interface{}, uids <-chan business.UserID) <-chan business.User {
	uStream := make(chan business.User)
	go func() {
		defer close(uStream)
		for id := range uids {
			fmt.Printf("[aggregate] User for UserID %v\n", id)

			user, err := getUser(id)
			if err != nil {
				// TODO address errors in a better way
				fmt.Println("some error ocurred")
			}

			select {
			case <-done:
				return
			case uStream <- user:
			}
		}
	}()

	return uStream
}

// getUser dummy function to simulate some fetching action on some user repository
func getUser(ID business.UserID) (business.User, error) {
	username := fmt.Sprintf("username_%v", ID)
	user := business.User{
		ID:       ID,
		Username: username,
		Email:    fmt.Sprintf("%v@pipeliner.com"),
		IsActive: true,
	}

	if ID%3 == 0 {
		user.IsActive = false
	}

	return user, nil
}
