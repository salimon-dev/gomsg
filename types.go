package gomsg

// generic reference type for all key/value pairs can be used as parameters.
type Parameters struct {
	StringValue string `json:"key"`
	StringKey   string `json:"value"`
}

// meta is the general info about a message when it's not just a plain interaction.
type Meta struct {
	ActionId string `json:"action_id"`
}

// message structurs in each interaction signal
type Message struct {
	From       string     `json:"from"`
	Type       string     `json:"type"`
	Body       string     `json:"body"`
	Meta       Meta       `json:"meta"`
	Parameters Parameters `json:"parameters"`
}

type InteractionSchema struct {
	Data []Message `json:"data"`
}
