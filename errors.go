package gomsg

type ValidationError struct {
	Type    string
	Message string
	Index   int
}

func (e *ValidationError) Error() string { return e.Message }

func recapErrors(errs *[]ValidationError, size int) *[]ValidationError {
	if errs == nil {
		return nil
	}
	if size == 0 {
		return nil
	}
	cappedErrors := make([]ValidationError, size)
	copy(cappedErrors, *errs)
	return &cappedErrors
}
