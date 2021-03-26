package map_from

import (
	"fmt"

	"github.com/godoylucase/go-data-pipelines-example/business"
)

func UPAggToPlainStruct(done <-chan interface{}, upAggregation <-chan business.UserProfileAggregation) <-chan business.PlainStruct {
	psStream := make(chan business.PlainStruct)
	go func() {
		defer close(psStream)
		for upa := range upAggregation {
			fmt.Printf("[map_from] UserAggregation to PlainStruct for UserID %v\n", upa.User.ID)

			select {
			case <-done:
				return
			case psStream <- upa.ToPlainStruct():
			}
		}
	}()

	return psStream
}
