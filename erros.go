package validate_nudity

import "fmt"

func NewError(msn string, arg interface{}) (err error) {
	err = fmt.Errorf("msn: %v arg: %v", msn, arg)
	return
}
