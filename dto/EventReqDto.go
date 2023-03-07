package dto

import "encoding/json"

type EventReqDto struct {
	EventType string `json:"eventType"`
	EventData string `json:"eventData"`
}

func NewEventReqDto(eventCode string, eventData interface{}) *EventReqDto {
	data, err := json.Marshal(eventData)
	if err != nil {
		panic(err)
	}
	return &EventReqDto{
		EventType: eventCode,
		EventData: string(data),
	}
}
