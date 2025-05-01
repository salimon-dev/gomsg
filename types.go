package gomsg

// generic reference type for all key/value pairs can be used as parameters.
type Parameters struct {
	// unique record key to identify an object that needs to be acted upon
	RecordKey   *string `json:"record_key"`
	StringValue *string `json:"string_value,omitempty"`
}

type ActionResult struct {
	Status  string `json:"status"`  // status of action, e.g. success or failure
	Message string `json:"message"` // additional information about result
}

// meta is the general info about a message when it's not just a plain interaction.
type Meta struct {
	ActionId string `json:"action_id" validate:"required"`
}

// message structurs in each interaction signal
type Message struct {
	From       string        `json:"from" validate:"required"`
	Type       string        `json:"type" validate:"required"`
	Body       *string       `json:"body,omitempty"`
	Meta       *Meta         `json:"meta,omitempty"`
	Parameters *Parameters   `json:"parameters,omitempty"`
	Result     *ActionResult `json:"result,omitempty"`
}

type InteractionSchema struct {
	Data []Message `json:"data" validate:"required,dive"`
}
