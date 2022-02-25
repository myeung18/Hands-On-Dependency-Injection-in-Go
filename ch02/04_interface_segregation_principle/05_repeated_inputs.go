package isp

import (
	"context"
	"time"
)

func UseEncryptV2() {
	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// store the key
	ctx = context.WithValue(ctx, "encryption-key", "-secret-")

	// call the function
	_, _ = EncryptV2(ctx, ctx, []byte("my data"))


	testCxtCall(ctx)
}


type newDeadLine interface {
	Deadline() (deadline time.Time, ok bool)
}


func testCxtCall(dl newDeadLine) {
	dl.Deadline()

}