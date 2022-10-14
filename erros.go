package validate_nudity

import "fmt"

func NewError(msn string, arg interface{}) (err error) {
	_, err = fmt.Printf("msn: %v arg: %v", msn, arg)
	return
}
