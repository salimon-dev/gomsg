package gomsg

import "testing"

func TestReactErrors(t *testing.T) {
	errs := make([]ValidationError, 2)
	recapped := recapErrors(&errs, 0)
	if recapped != nil {
		t.Fatal("recapErrors should return nil if there are no errors")
	}
}

func TestRecapWithWrongSize(t *testing.T) {
	errs := make([]ValidationError, 1)
	errs[0] = ValidationError{
		Type:    "wrong",
		Index:   0,
		Message: "error message",
	}
	recapped := recapErrors(&errs, 1)
	if recapped == nil {
		t.Fatal("recapErrors should return an error if the size is wrong")
	}
	if len(*recapped) != 1 {
		t.Fatalf("recapErrors should return a single error")
	}
	if (*recapped)[0].Error() != "error message" {
		t.Fatalf("recapErrors should return the correct error")
	}
}

func TestNilErrors(t *testing.T) {
	recapped := recapErrors(nil, 0)
	if recapped != nil {
		t.Fatal("recapErrors should return nil if there are no errors")
	}
}
