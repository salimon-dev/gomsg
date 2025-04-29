package gomsg

func GetActionResult(actionMessage *Message, messages *[]Message) *Message {
	for _, message := range *messages {
		if message.Type != "actionResult" {
			continue
		}
		if message.Meta.ActionId == actionMessage.Meta.ActionId {
			return &message
		}
	}
	return nil
}

func IsActionResolved(actionMessage *Message, actionResults *[]Message) bool {
	return GetActionResult(actionMessage, actionResults) != nil
}

func ExtractUnresolvedActionMessages(messages *[]Message) []*Message {
	actionMessages := make([]*Message, len(*messages))
	index := 0
	for _, message := range *messages {
		switch message.Type {
		case "getStringValue", "removeStringValue", "setStringValue":
			if IsActionResolved(&message, messages) {
				continue
			}
			actionMessages[index] = &message
			index++
			break
		default:
			break
		}
	}
	result := make([]*Message, index)
	copy(result, actionMessages[:index])
	return result
}
