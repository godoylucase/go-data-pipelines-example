package stream

import (
	"testing"

	"github.com/godoylucase/go-data-pipelines-example/business"
	"github.com/stretchr/testify/assert"
)

func Test_generator(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	ids := []business.UserID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	idsChan := UserIDs(done, ids...)

	for id := range idsChan {
		assert.LessOrEqual(t, id, uint(10))
	}

	_, ok := <-idsChan
	assert.False(t, ok)
}
