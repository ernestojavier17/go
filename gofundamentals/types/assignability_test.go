package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/ernestoca/go-fundamentals/dummy"
	"github.com/magiconair/properties/assert"
)

func TestAssignmentRules(t *testing.T) {

	ctx := &Context{
		store:make(Map),
	}

	user := &User{
		ID:                     0,
		CurrentGroupIdentifier: "1",
		Email:                  "",
		ReferenceIdentifier:    "",
		WhenCreated:            time.Time{},
		WhenCleared:            nil,
		VerifiedEmail:          false,
		IsSuspended:            false,
		WhenFlagged:            nil,
	}

	key := "user-entity"
	ctx.Set(key, user)
	//Print the type and values
	fmt.Printf("type: %T :: val:%+v", ctx.Get(key), ctx.Get(key))
	//Assign generic type interface to a different type with same properties
	//Type assertion
	//val, ok := ctx.Get(key).(User)
	val, ok := ctx.Get(key).(*dummy.User)
	assert.Equal(t, ok, true)
	fmt.Printf("Type Assertion result: %v", val)

}
