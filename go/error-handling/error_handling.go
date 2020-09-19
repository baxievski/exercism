// Package erratum implements a solution for the exercism error-handling challenge
package erratum

// Use opens a resource, retries on transient errors, takes care to always close the resource
func Use(o ResourceOpener, s string) (err error) {
	res, err := o()
	for err != nil {
		switch err.(type) {
		case TransientError:
			res, err = o()
		default:
			return err
		}
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
				break
			}
		}
	}()

	res.Frob(s)
	return nil
}
