package gomsg

import (
	"testing"
)

func TestEmptyMessage(t *testing.T) {
	invalidJSON := []byte(`{"data": [{}]}`)

	_, errs := ParseInteractionSchema(invalidJSON)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "FromRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "TypeRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}
func TestMessageWithoutFrom(t *testing.T) {
	invalidJSON := []byte(`{"data": [{"type": "plain"}]}`)

	_, errs := ParseInteractionSchema(invalidJSON)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "FromRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "BodyRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}
func TestMessageWithoutType(t *testing.T) {
	invalidJSON := []byte(`{"data": [{"from": "user"}]}`)

	_, errs := ParseInteractionSchema(invalidJSON)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) != 1 {
		t.Fatalf("Expected 1 errors to be returned")
	}
	if (*errs)[0].Type != "TypeRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}

func TestPlainMessageWithoutBody(t *testing.T) {
	invalidJSON := []byte(`{"data": [{"from": "user", "type": "plain"}]}`)

	_, errs := ParseInteractionSchema(invalidJSON)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) == 0 {
		t.Fatalf("Expected errors to be returned")
	}
	if (*errs)[0].Type != "BodyRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}
func TestPlainMessageWithBody(t *testing.T) {
	invalidJSON := []byte(`{"data": [{"from": "user", "type": "plain", "body": "some body"}]}`)

	_, errs := ParseInteractionSchema(invalidJSON)
	if errs != nil {
		t.Fatalf("Expected no errors to be returned")
	}
}

func TestActionMessageWithoutMeta(t *testing.T) {
	data := []byte(`{"data": [{"from": "user", "type": "setStringValue"}]}`)

	_, errs := ParseInteractionSchema(data)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "MetaRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "ParametersRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}

func TestActionMessageWithMeta(t *testing.T) {
	data := []byte(`{"data": [{"from": "user", "type": "setStringValue", "meta": {}}]}`)

	_, errs := ParseInteractionSchema(data)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "ActionIdRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "ParametersRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}

func TestActionMessageWithoutParameters(t *testing.T) {
	data := []byte(`{"data": [{"from": "user", "type": "setStringValue", "meta": {"action_id": "somerandomid"}}]}`)

	_, errs := ParseInteractionSchema(data)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if len(*errs) != 1 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "ParametersRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}

func TestActionMessageWithParameters(t *testing.T) {
	data := []byte(`{"data": [{"from": "user", "type": "setStringValue", "meta": {"action_id": "somerandomid"}, "parameters": {"string_value": "somestrinvalue", "record_key": "somekey"}}]}`)

	_, errs := ParseInteractionSchema(data)
	if errs != nil {
		t.Fatalf("Expected no errors to be returned")
	}
}
func TestActionMessageWithWrongParameters(t *testing.T) {
	data := []byte(`{"data": [{"from": "user", "type": "setStringValue", "meta": {"action_id": "somerandomid"}, "parameters": {"string_value": "somestrinvalue"}}]}`)

	_, errs := ParseInteractionSchema(data)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty schema")
	}
	if (*errs)[0].Type != "RecordKeyRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
}

func TestActionResultWithoutResult(t *testing.T) {
	message := Message{
		From: "archivist",
		Type: "actionResult",
	}
	errs := validateMessage(&message, 0)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty action result")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "MetaRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "ResultRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}

func TestActionResultWithoutActionId(t *testing.T) {
	message := Message{
		From: "archivist",
		Type: "actionResult",
		Meta: &Meta{},
	}
	errs := validateMessage(&message, 0)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty action result")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "ActionIdRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "ResultRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}

func TestActionResultWithoutStatus(t *testing.T) {
	message := Message{
		From: "archivist",
		Type: "actionResult",
		Meta: &Meta{
			ActionId: "randomid",
		},
		Result: &ActionResult{},
	}
	errs := validateMessage(&message, 0)
	if errs == nil {
		t.Fatalf("Expected an error when parsing empty action result")
	}
	if len(*errs) != 2 {
		t.Fatalf("Expected 2 errors to be returned")
	}
	if (*errs)[0].Type != "StatusRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[0].Type)
	}
	if (*errs)[1].Type != "MessageRequired" {
		t.Fatalf("Unexpected error type: %s", (*errs)[1].Type)
	}
}
