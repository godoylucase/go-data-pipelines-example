package aggregate

import (
	"fmt"

	"github.com/godoylucase/go-data-pipelines-example/business"
)

type ProfileFetcher interface {
	GetByUserID(uid business.UserID) (business.Profile, error)
}

func Profile(done <-chan interface{}, users <-chan business.User) <-chan business.UserProfileAggregation {
	upaStream := make(chan business.UserProfileAggregation)
	go func() {
		defer close(upaStream)
		for user := range users {
			fmt.Printf("[aggregate] Profile for UserID %v\n", user.ID)

			profile, err := getByUserID(user.ID)
			if err != nil {
				// TODO address errors in a better way
				fmt.Println("some error ocurred")
			}

			select {
			case <-done:
				return
			case upaStream <- business.UserProfileAggregation{
				User:    user,
				Profile: profile,
			}:
			}
		}
	}()

	return upaStream
}

// getByUserID dummy function to simulate some fetching action on some profile repository
func getByUserID(uids business.UserID) (business.Profile, error) {
	p := business.Profile{
		ID:       business.ProfileID(uint(uids) + 100),
		PhotoURL: fmt.Sprintf("https://some-storage-url/%v-photo", uids),
	}

	return p, nil
}
