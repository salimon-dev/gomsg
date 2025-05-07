package gomsg

func ValidateActionParameters(message *Message, index int) *[]ValidationError {
	switch message.Type {
	case "setStringValue":
		return ValidateSetStringValue(message, index)
	case "getStringValue":
		return ValidateGetStringValue(message, index)
	case "removeStringValue":
		return ValidateRemoveStringValue(message, index)
	default:
		return nil
	}
}

func ValidateGetStringValue(message *Message, index int) *[]ValidationError {
	results := make([]ValidationError, 1)
	errIndex := 0
	if message.Parameters.RecordKey == nil {
		results[errIndex] = ValidationError{
			Type:    "RecordKeyRequired",
			Message: "Missing record key in parameter",
			Index:   index,
		}
		errIndex++
	}
	return recapErrors(&results, errIndex)
}
func ValidateRemoveStringValue(message *Message, index int) *[]ValidationError {
	results := make([]ValidationError, 1)
	errIndex := 0
	if message.Parameters.RecordKey == nil {
		results[errIndex] = ValidationError{
			Type:    "RecordKeyRequired",
			Message: "Missing record key in parameter",
			Index:   index,
		}
		errIndex++
	}
	return recapErrors(&results, errIndex)
}
func ValidateSetStringValue(message *Message, index int) *[]ValidationError {
	results := make([]ValidationError, 2)
	errIndex := 0
	if message.Parameters.RecordKey == nil {
		results[errIndex] = ValidationError{
			Type:    "RecordKeyRequired",
			Message: "Missing record key in parameter",
			Index:   index,
		}
		errIndex++
	}
	if message.Parameters.StringValue == nil {
		results[errIndex] = ValidationError{
			Type:    "StringValueRequired",
			Message: "Missing string value in parameter",
			Index:   index,
		}
		errIndex++
	}
	return recapErrors(&results, errIndex)
}
