package erratum

import "errors"

func Use(opener ResourceOpener, input string) error {
	resource, err := opener()
	if err != nil && errors.Is(err, TransientError{}) {
		Use(opener, input)		
	} else if err != nil {
		return errors.Unwrap(Use(opener, input))
	}
	resource.Frob(input)
	resource.Close()
	return nil
}
