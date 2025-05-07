package gomsg

import "testing"

func TestGetActionResultNotFound(t *testing.T) {
	data := []byte(`{"data":[{"from":"user","type":"plain","body":"set my name to ali"},{"from":"tina","type":"setStringValue","meta":{"action_id":"action_id_00"},"parameters":{"record_key":"name","string_value":"ali"}}]}`)
	payload, errs := ParseInteractionSchema(data)

	if errs != nil {
		t.Fatal("failed to parse data")
	}
	if payload == nil {
		t.Fatal("payload is nil")
	}
	if len(payload.Data) != 2 {
		t.Fatalf("expected %d messages but got %d", 2, len(payload.Data))
	}

	actions := ExtractUnresolvedActionMessages(&payload.Data)
	if len(actions) != 1 {
		t.Fatalf("expected %d unresolved actions but got %d", 1, len(actions))
	}
}

func TestGetActionResultFound(t *testing.T) {
	data := []byte(`{"data":[{"from":"user","type":"plain","body":"set my name to ali"},{"from":"tina","type":"setStringValue","meta":{"action_id":"action_id_00"},"parameters":{"record_key":"name","string_value":"ali"}},{"from":"archivist","type":"actionResult","meta":{"action_id":"action_id_00"},"result":{"status":"success","message":"name set to ali"}}]}`)
	payload, errs := ParseInteractionSchema(data)

	if errs != nil {
		t.Fatal("failed to parse data")
	}
	if payload == nil {
		t.Fatal("payload is nil")
	}
	if len(payload.Data) != 3 {
		t.Fatalf("expected %d messages but got %d", 2, len(payload.Data))
	}

	actions := ExtractUnresolvedActionMessages(&payload.Data)
	if len(actions) != 0 {
		t.Fatalf("expected %d unresolved actions but got %d", 0, len(actions))
	}
}

func TestGetActionsExtra(t *testing.T) {
	data := []byte(`{"data":[{"from":"user","type":"plain","body":"set my name to ali and my birth year to 1996"},{"from":"tina","type":"setStringValue","meta":{"action_id":"action_id_00"},"parameters":{"record_key":"name","string_value":"ali"}},{"from":"tina","type":"setStringValue","meta":{"action_id":"action_id_01"},"parameters":{"record_key":"birth_year","string_value":"1996"}},{"from":"archivist","type":"actionResult","meta":{"action_id":"action_id_00"},"result":{"status":"success","message":"name set to ali"}}]}`)
	payload, errs := ParseInteractionSchema(data)

	if errs != nil {
		t.Fatal("failed to parse data")
	}
	if payload == nil {
		t.Fatal("payload is nil")
	}
	if len(payload.Data) != 4 {
		t.Fatalf("expected %d messages but got %d", 4, len(payload.Data))
	}

	actions := ExtractUnresolvedActionMessages(&payload.Data)
	if len(actions) != 1 {
		t.Fatalf("expected %d unresolved actions but got %d", 1, len(actions))
	}
}
