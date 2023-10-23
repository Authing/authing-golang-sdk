package dto


type PubEventDto struct{
    EventType  string `json:"eventType"`
    EventData  interface{} `json:"eventData"`
}

