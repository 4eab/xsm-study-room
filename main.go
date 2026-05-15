package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	username string
}

type Room struct {
	clients   map[*Client]bool
	broadcast chan string
	mu        sync.Mutex
}

var room = Room{
	clients:   make(map[*Client]bool),
	broadcast: make(chan string), // A broadcast channel is unbuffered, which is dangerous.
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Should close conn
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}

	username := r.URL.Query().Get("username")

	if username == "" {
		username = "Guest"
	}

	client := &Client{
		conn:     conn,
		username: username,
	}

	// In the nebulous boundary where the current coroutine logic has not yet fully exited
	// —or while a lock is still being held—
	// allowing synchronous Channel sends to become intertwined with lock operations creates a precarious situation. Unbuffered Channels, in particular, act as "time bombs" in high-concurrency environments.
	defer func() {
		room.mu.Lock()
		delete(room.clients, client)
		room.mu.Unlock()

		room.broadcast <- client.username + " left" // Performing a channel send within a `defer` statement carries risks.
	}()

	// This involves purely in-memory operations.
	// Writing a pointer to a Go Map takes the CPU merely a few nanoseconds—one-thousandth of a microsecond—to complete.
	// Gemini: If you wish to completely eliminate the potential for queuing bottlenecks associated with locks in the future,
	// the current industry standard is to use Channels as a substitute for locks.
	fmt.Printf("%s connected\n", client.username)
	room.mu.Lock()
	room.clients[client] = true
	room.mu.Unlock()

	// Receive and broadcast messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		fmt.Printf("%s: %s\n", client.username, string(msg))
		// Broadcast the message
		room.broadcast <- client.username + ": " + string(msg)
	}
}

func (r *Room) run() {
	fmt.Println("ROOM RUN STARTED")
	for {
		msg := <-r.broadcast
		fmt.Printf("broadcast: \"%s\"\n", string(msg))

		// A lock is used to iterate through all clients and send messages,
		// the loop will stall if even a single client has an extremely slow network connection or a deadlocked connection.
		// Instead, each `Client` struct should be allocated its own dedicated send channel;
		// the `room.run` routine is then responsible solely for dropping messages into each individual's channel,
		// while each client's own goroutine handles the actual writing.
		r.mu.Lock()
		for client := range r.clients {
			client.conn.WriteMessage(websocket.TextMessage, []byte(msg)) // Critical Point: Network I/O Operations
		}
		r.mu.Unlock()
	}
}

func main() {
	go room.run()

	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8080", nil)
}
