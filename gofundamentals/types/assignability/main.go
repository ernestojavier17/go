package assignability

import (
	"time"
)

type (
	Map map[string]interface{}

	Context struct {
		store Map
	}
)



type User struct {
	ID                     int64
	CurrentGroupIdentifier string
	Email                  string
	ReferenceIdentifier    string
	WhenCreated            time.Time
	WhenCleared            *time.Time
	VerifiedEmail          bool
	IsSuspended            bool
	WhenFlagged            *time.Time
}



func (c *Context) Set(key string, val interface{}) {
	if c.store == nil {
		c.store = make(Map)
	}

	c.store[key] = val
}

func (c *Context) Get(key string) interface{} {
	return c.store[key]
}
