package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestBroadcast(t *testing.T) {
	go room.run()

	server := httptest.NewServer(
		http.HandlerFunc(wsHandler),
	)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	c1, _, err := websocket.DefaultDialer.Dial(
		wsURL+"?username=alice",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer c1.Close()

	c2, _, err := websocket.DefaultDialer.Dial(
		wsURL+"?username=bob",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close()

	time.Sleep(100 * time.Millisecond)

	err = c1.WriteMessage(
		websocket.TextMessage,
		[]byte("hello"),
	)
	if err != nil {
		t.Fatal(err)
	}

	c2.SetReadDeadline(time.Now().Add(2 * time.Second))

	_, msg, err := c2.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(msg), "hello") {
		t.Fatalf("unexpected msg: %s", string(msg))
	}
}
