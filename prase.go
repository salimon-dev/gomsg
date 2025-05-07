package gomsg

import (
	"encoding/json"
	"fmt"
)

func ParseInteractionSchema(data []byte) (*InteractionSchema, *[]ValidationError) {
	var result *InteractionSchema
	err := json.Unmarshal(data, &result)
	if err != nil {
		errs := make([]ValidationError, 1)
		errs[0] = ValidationError{
			Type:    "InvalidJson",
			Message: err.Error(),
			Index:   -1,
		}
		return nil, &errs
	}

	dataError := validateData(result)
	if dataError != nil {
		return nil, &[]ValidationError{*dataError}
	}

	validationErrors := make([]ValidationError, 128)
	errsIndex := 0

	for i, message := range result.Data {
		messageErrs := validateMessage(&message, i)
		if messageErrs == nil {
			continue
		}
		copy(validationErrors[errsIndex:], *messageErrs)
		errsIndex += len(*messageErrs)
	}

	return result, recapErrors(&validationErrors, errsIndex)

}

func validateData(payload *InteractionSchema) *ValidationError {
	if payload.Data == nil {
		return &ValidationError{
			Type:    "DataRequired",
			Message: "data is required",
			Index:   -1,
		}
	}
	if len(payload.Data) == 0 {
		return &ValidationError{
			Type:    "DataNotEmpty",
			Message: "data must not be empty",
			Index:   -1,
		}
	}
	return nil
}

func validateMessage(message *Message, index int) *[]ValidationError {
	errs := make([]ValidationError, 128)
	errIndex := 0

	if message.From == "" {
		errs[errIndex] = ValidationError{
			Type:    "FromRequired",
			Index:   index,
			Message: fmt.Sprintf("from is required in message %d", index),
		}
		errIndex++
	}

	if message.Type == "" {
		errs[errIndex] = ValidationError{
			Type:    "TypeRequired",
			Index:   index,
			Message: fmt.Sprintf("type is required in message %d", index),
		}
		errIndex++
		cappedErrors := recapErrors(&errs, errIndex)
		return cappedErrors
	}

	if message.Type == "plain" {
		plainErrors := validatePlainMessage(message, index)
		if plainErrors == nil {
			return recapErrors(&errs, errIndex)
		}
		copy(errs[errIndex:], *plainErrors)
		errIndex += len(*plainErrors)
	} else if message.Type == "actionResult" {
		plainErrors := validateActionResultMessage(message, index)
		if plainErrors == nil {
			return recapErrors(&errs, errIndex)
		}
		copy(errs[errIndex:], *plainErrors)
		errIndex += len(*plainErrors)
	} else {
		plainErrors := validateActionMessage(message, index)
		if plainErrors == nil {
			return recapErrors(&errs, errIndex)
		}
		copy(errs[errIndex:], *plainErrors)
		errIndex += len(*plainErrors)
	}

	cappedErrors := recapErrors(&errs, errIndex)
	return cappedErrors
}

func validatePlainMessage(message *Message, index int) *[]ValidationError {
	errs := make([]ValidationError, 1)
	if message.Body == nil || *message.Body == "" {
		errs[0] = ValidationError{
			Type:    "BodyRequired",
			Index:   index,
			Message: fmt.Sprintf("body is required in message %d", index),
		}
		return &errs
	}
	return nil

}

func validateActionMessage(message *Message, index int) *[]ValidationError {
	errs := make([]ValidationError, 32)
	errsIndex := 0
	if message.Meta == nil {
		errs[errsIndex] = ValidationError{
			Type:    "MetaRequired",
			Index:   index,
			Message: fmt.Sprintf("meta is required in message %d", index),
		}
		errsIndex++
	} else {
		if message.Meta.ActionId == "" {
			errs[errsIndex] = ValidationError{
				Type:    "ActionIdRequired",
				Index:   index,
				Message: fmt.Sprintf("action id is required in message %d", index),
			}
			errsIndex++
		}
	}
	if message.Parameters == nil {
		errs[errsIndex] = ValidationError{
			Type:    "ParametersRequired",
			Index:   index,
			Message: fmt.Sprintf("parameters are required in message %d", index),
		}
		errsIndex++
	} else {
		actionErrs := ValidateActionParameters(message, index)
		if actionErrs != nil {
			copy(errs[errsIndex:], *actionErrs)
			errsIndex += len(*actionErrs)
		}

	}
	cappedErrors := recapErrors(&errs, errsIndex)
	return cappedErrors
}

func validateActionResultMessage(message *Message, index int) *[]ValidationError {
	errs := make([]ValidationError, 32)
	errsIndex := 0
	if message.Meta == nil {
		errs[errsIndex] = ValidationError{
			Type:    "MetaRequired",
			Index:   index,
			Message: fmt.Sprintf("meta is required in message %d", index),
		}
		errsIndex++
	} else {
		if message.Meta.ActionId == "" {
			errs[errsIndex] = ValidationError{
				Type:    "ActionIdRequired",
				Index:   index,
				Message: fmt.Sprintf("action id is required in message %d", index),
			}
			errsIndex++
		}
	}
	if message.Result == nil {
		errs[errsIndex] = ValidationError{
			Type:    "ResultRequired",
			Index:   index,
			Message: fmt.Sprintf("result is required in message %d", index),
		}
		errsIndex++
	} else {
		if message.Result.Status == "" {
			errs[errsIndex] = ValidationError{
				Type:    "StatusRequired",
				Index:   index,
				Message: fmt.Sprintf("status is required in message %d", index),
			}
			errsIndex++
		}
		if message.Result.Message == "" {
			errs[errsIndex] = ValidationError{
				Type:    "MessageRequired",
				Index:   index,
				Message: fmt.Sprintf("message is required in message %d", index),
			}
			errsIndex++
		}
	}
	cappedErrors := recapErrors(&errs, errsIndex)
	return cappedErrors
}
