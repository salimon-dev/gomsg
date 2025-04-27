package gomsg

import "encoding/json"

func ParseInteractionSchema(data []byte) (*InteractionSchema, error) {
	var interactionSchema *InteractionSchema
	err := json.Unmarshal(data, &interactionSchema)
	return interactionSchema, err
}
