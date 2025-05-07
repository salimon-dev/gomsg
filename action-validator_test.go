package gomsg

import "testing"

func TestValidateSetStringValue(t *testing.T) {
	message := Message{
		From: "entity",
		Type: "setStringValue",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Parameters: &Parameters{},
	}

	errs := ValidateSetStringValue(&message, 0)
	if len(*errs) != 2 {
		t.Errorf("Expected %d errors but got %d", 2, len(*errs))
	}
	if (*errs)[0].Type != "RecordKeyRequired" {
		t.Errorf("Expected error type RecordKeyRequired but got %v", (*errs)[0])
	}
	if (*errs)[1].Type != "StringValueRequired" {
		t.Errorf("Expected error type StringValueRequired but got %v", (*errs)[1])
	}

	recordKey := "somekey"
	recordValue := "somevalue"

	message.Parameters.RecordKey = &recordKey
	errs = ValidateSetStringValue(&message, 0)
	if len(*errs) != 1 {
		t.Errorf("Expected %d errors but got %d", 1, len(*errs))
	}
	if (*errs)[0].Type != "StringValueRequired" {
		t.Errorf("Expected error type StringValueRequired but got %v", (*errs)[0])
	}

	message.Parameters.RecordKey = nil
	message.Parameters.StringValue = &recordValue
	errs = ValidateSetStringValue(&message, 0)
	if len(*errs) != 1 {
		t.Errorf("Expected %d errors but got %d", 1, len(*errs))
	}
	if (*errs)[0].Type != "RecordKeyRequired" {
		t.Errorf("Expected error type RecordKeyRequired but got %v", (*errs)[0])
	}

	message.Parameters.RecordKey = &recordKey
	message.Parameters.StringValue = &recordValue
	errs = ValidateSetStringValue(&message, 0)
	if errs != nil {
		t.Errorf("Expected %d errors but got %d", 0, len(*errs))
	}
}

func TestValidateGetStringValue(t *testing.T) {
	message := Message{
		From: "entity",
		Type: "getStringValue",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Parameters: &Parameters{},
	}

	errs := ValidateGetStringValue(&message, 0)
	if len(*errs) != 1 {
		t.Errorf("Expected %d errors but got %d", 2, len(*errs))
	}
	if (*errs)[0].Type != "RecordKeyRequired" {
		t.Errorf("Expected error type RecordKeyRequired but got %v", (*errs)[0])
	}

	recordKey := "somekey"

	message.Parameters.RecordKey = &recordKey
	errs = ValidateGetStringValue(&message, 0)
	if errs != nil {
		t.Errorf("Expected %d errors but got %d", 0, len(*errs))
	}
}
func TestValidateRemoveStringValue(t *testing.T) {
	message := Message{
		From: "entity",
		Type: "removeStringValue",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Parameters: &Parameters{},
	}

	errs := ValidateRemoveStringValue(&message, 0)
	if len(*errs) != 1 {
		t.Errorf("Expected %d errors but got %d", 2, len(*errs))
	}
	if (*errs)[0].Type != "RecordKeyRequired" {
		t.Errorf("Expected error type RecordKeyRequired but got %v", (*errs)[0])
	}

	recordKey := "somekey"

	message.Parameters.RecordKey = &recordKey
	errs = ValidateRemoveStringValue(&message, 0)
	if errs != nil {
		t.Errorf("Expected %d errors but got %d", 0, len(*errs))
	}
}

func TestValidationActionParameters(t *testing.T) {
	recordKey := "somekey"
	stringValue := "somevalue"
	message := Message{
		From: "entity",
		Type: "removeStringValue",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Parameters: &Parameters{
			RecordKey: &recordKey,
		},
	}
	errors := ValidateActionParameters(&message, 0)
	if errors != nil {
		t.Errorf("Expected no validation errors")
	}
	message = Message{
		From: "entity",
		Type: "getStringValue",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Parameters: &Parameters{
			RecordKey: &recordKey,
		},
	}
	errors = ValidateActionParameters(&message, 0)
	if errors != nil {
		t.Errorf("Expected no validation errors")
	}
	message = Message{
		From: "entity",
		Type: "stringStringValue",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Parameters: &Parameters{
			RecordKey:   &recordKey,
			StringValue: &stringValue,
		},
	}
	errors = ValidateActionParameters(&message, 0)
	if errors != nil {
		t.Errorf("Expected no validation errors")
	}
}
