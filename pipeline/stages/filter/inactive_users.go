package filter

import (
	"fmt"

	"github.com/godoylucase/go-data-pipelines-example/business"
)

func InactiveUsers(done <-chan interface{}, users <-chan business.User) <-chan business.User {
	activeUsers := make(chan business.User)
	go func() {
		defer close(activeUsers)
		for user := range users {
			if !user.IsActive {
				fmt.Printf("[filter] Inactive user with UserID %v\n", user.ID)
				continue
			}

			select {
			case <-done:
				return
			case activeUsers <- user:
			}
		}
	}()

	return activeUsers
}
