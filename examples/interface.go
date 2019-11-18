package examples

import "time"

type MyGeniusIdea interface {
	Smart(awesome string, worst interface{}) (greatReturn time.Time, err error)
}
