package main

import (
	"fmt"

	"github.com/godoylucase/go-data-pipelines-example/business"
	"github.com/godoylucase/go-data-pipelines-example/pipeline/stages/aggregate"
	"github.com/godoylucase/go-data-pipelines-example/pipeline/stages/filter"
	"github.com/godoylucase/go-data-pipelines-example/pipeline/stages/map_from"
	"github.com/godoylucase/go-data-pipelines-example/pipeline/stages/stream"
)

const maxUserID = 100

func main() {
	done := make(chan interface{})
	defer close(done)

	userIDs := make([]business.UserID, maxUserID)
	for i := 1; i <= maxUserID; i++ {
		userIDs = append(userIDs, business.UserID(i))
	}

	plainStructs := map_from.UPAggToPlainStruct(done,
		aggregate.Profile(done,
			filter.InactiveUsers(done,
				aggregate.User(done,
					stream.UserIDs(done, userIDs...)))),
	)

	for ps := range plainStructs {
		fmt.Printf("[result] plain struct for UserID %v is: -> %v \n", ps.UserID, ps)
	}
}
