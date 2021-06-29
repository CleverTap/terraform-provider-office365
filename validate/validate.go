package validate

import (
	"fmt"
	"regexp"
)

func ValidateEmail(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value := v.(string)
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("Expected EmailId is not valid  %s", k))
		return warns, errs
	}
	return
}
