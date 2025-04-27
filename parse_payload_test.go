package gomsg

import "testing"

func TestParseEmptySchema(t *testing.T) {
	invalidJSON := []byte(``)

	_, errs := ParseInteractionSchema(invalidJSON)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) == 0 {
		t.Fatalf("Expected errors to be returned")
	}
	if (*errs)[0].Type != "InvalidJson" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}

func TestParseNoDataSchema(t *testing.T) {
	payload := []byte(`{}`)
	_, errs := ParseInteractionSchema(payload)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) == 0 {
		t.Fatalf("Expected errors to be returned")
	}
	if (*errs)[0].Type != "DataRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}
func TestParseEmptyDataSchema(t *testing.T) {
	payload := []byte(`{ "data": [] }`)
	_, errs := ParseInteractionSchema(payload)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) == 0 {
		t.Fatalf("Expected errors to be returned")
	}
	if (*errs)[0].Type != "DataNotEmpty" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}

func TestParseInvalidaData(t *testing.T) {
	payload := []byte(`{"data": {}}`)
	_, errs := ParseInteractionSchema(payload)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) == 0 {
		t.Fatalf("Expected errors to be returned")
	}
	if (*errs)[0].Type != "InvalidJson" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}

func TestParseInvalidaDataEntry(t *testing.T) {
	payload := []byte(`{"data": ["string"]}`)
	_, errs := ParseInteractionSchema(payload)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) == 0 {
		t.Fatalf("Expected errors to be returned")
	}
	if (*errs)[0].Type != "InvalidJson" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}
