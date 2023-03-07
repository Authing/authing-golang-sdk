package util

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketEventHub struct {
	Connections map[string]*websocket.Conn
	receiveMap  map[string]bool
	funcMap     sync.Map
}

type EventReceiver interface {
	OnSuccess(msg []byte)
	OnError(err error)
}

type EventReceives struct {
	Successes []func(msg []byte)
	Errors    []func(err error)
}

// NewWebSocketEvent
func NewWebSocketEvent() *WebSocketEventHub {
	return &WebSocketEventHub{
		Connections: make(map[string]*websocket.Conn),
		receiveMap:  make(map[string]bool),
	}
}

func NewEventReceives() *EventReceives {
	return &EventReceives{
		Successes: make([]func(msg []byte), 0),
		Errors:    make([]func(err error), 0),
	}
}

func (eventHub *WebSocketEventHub) CreateManagement(eventCode string, websocketHost string, token string) bool {
	if eventHub.Connections[eventCode] == nil {
		var socketUri = fmt.Sprintf("%s/events/v1/management/sub?code=%s", websocketHost, eventCode)
		// fmt.Println(socketUri)
		conn, _, err := websocket.DefaultDialer.Dial(
			socketUri,
			http.Header{
				"Authorization": []string{token},
			},
		)
		if err != nil {
			log.Printf("dail: %s", err.Error())
			return false
		}
		eventHub.Connections[eventCode] = conn
	}
	return true
}

func (eventHub *WebSocketEventHub) CreateAuthentication(eventCode string, websocketHost string, token string) bool {
	if eventHub.Connections[eventCode] == nil {
		var socketUri = fmt.Sprintf("%s/events/v1/authentication/sub?code=%s&token=%s", websocketHost, eventCode, token)
		// fmt.Println(socketUri)
		conn, _, err := websocket.DefaultDialer.Dial(
			socketUri,
			http.Header{},
		)
		if err != nil {
			log.Printf("dail: %s", err.Error())
			return false
		}
		eventHub.Connections[eventCode] = conn
	}
	return true
}

func (eventHub *WebSocketEventHub) AddReceiver(eventCode string, onSuccess func(msg []byte), onError func(err error)) {
	receivers := NewEventReceives()
	if funcMap, ok := eventHub.funcMap.Load(eventCode); ok {
		receivers.Successes = append(funcMap.(*EventReceives).Successes, onSuccess)
		receivers.Errors = append(funcMap.(*EventReceives).Errors, onError)
	} else {
		receivers.Successes = []func(msg []byte){onSuccess}
		receivers.Errors = []func(err error){onError}
	}
	eventHub.funcMap.Store(eventCode, receivers)
}

func (eventHub *WebSocketEventHub) StartReceive(eventCode string) {
	started, ok := eventHub.receiveMap[eventCode]
	if ok && started {
		return
	}
	eventHub.receiveMap[eventCode] = true
	for {
		_, message, err := eventHub.Connections[eventCode].ReadMessage()
		if funcMap, ok := eventHub.funcMap.Load(eventCode); ok {
			funcs := funcMap.(*EventReceives)
			if err != nil {
				for _, onError := range funcs.Errors {
					onError(err)
				}
			} else {
				for _, onSuccess := range funcs.Successes {
					onSuccess(message)
				}
			}
		}

		time.Sleep(time.Microsecond * 500)
	}
}
