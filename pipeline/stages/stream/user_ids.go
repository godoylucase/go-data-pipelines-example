package stream

import (
	"fmt"

	"github.com/godoylucase/go-data-pipelines-example/business"
)

func UserIDs(done <-chan interface{}, uids ...business.UserID) <-chan business.UserID {
	uidsStream := make(chan business.UserID)
	go func() {
		defer close(uidsStream)
		for _, uid := range uids {
			fmt.Printf("[stream] UserID %v\n", uid)

			select {
			case <-done:
				return
			case uidsStream <- uid:
			}
		}
	}()
	return uidsStream
}
