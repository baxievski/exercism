// Package erratum implements a solution for the exercism error-handling challenge
package erratum

import "fmt"

// Use opens a resource, retries on transient errors, takes care to always close the resource
func Use(o ResourceOpener, s string) (err error) {
	res, err := o()
	for err != nil {
		if _, ok := err.(TransientError); ok {
			res, err = o()
			continue
		}
		return err
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	res.Frob(s)
	return nil
}
