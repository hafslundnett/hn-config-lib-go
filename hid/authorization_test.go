package hid

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// at overrides time value for tests and restores default value after
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
} //time.Unix(0, 0)
