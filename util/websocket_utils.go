package util

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait = 15 * time.Second
)

type WebSocketEventHub struct {
	Connections map[string]*websocket.Conn
	listenTag   sync.Map
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
	}
}

func NewEventReceives() *EventReceives {
	return &EventReceives{
		Successes: make([]func(msg []byte), 0),
		Errors:    make([]func(err error), 0),
	}
}

func createConnection(uri string, headers http.Header) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(uri, headers)
	if err != nil {
		log.Fatal("dial:", err)
		return nil
	}
	// conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	return conn
}

func (eventHub *WebSocketEventHub) CreateManagement(eventCode string, websocketHost string, token string) bool {
	if eventHub.Connections[eventCode] == nil {
		var socketUri = fmt.Sprintf("%s/events/v1/management/sub?code=%s", websocketHost, eventCode)
		// fmt.Println(socketUri)
		conn := createConnection(socketUri, http.Header{
			"Authorization": []string{token},
		})
		if conn == nil {
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
		conn := createConnection(socketUri, http.Header{})
		if conn == nil {
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

func connPong(conn *websocket.Conn) {
	ticker := time.NewTicker(pongWait)
	defer ticker.Stop()
	for range ticker.C {
		err := conn.WriteMessage(websocket.PongMessage, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (eventHub *WebSocketEventHub) StartReceive(eventCode string) {
	started, loaded := eventHub.listenTag.LoadOrStore(eventCode, true)
	if loaded && started.(bool) {
		return
	}
	log.Println("start connection receive")

	conn := eventHub.Connections[eventCode]
	defer conn.Close()
	go connPong(conn)
	ticker := time.NewTicker(pongWait)
	defer ticker.Stop()
	begin_time := time.Now()
	count := 0
	for {
		select {
		case <-ticker.C:
			log.Printf("received %v messages, has been running for %v seconds", count, time.Since(begin_time).Seconds())
		default:
			_, message, err := conn.ReadMessage()
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
			count += 1
		}
	}
}
