# Changelog

## 2026-05-15 Initial MVP: websocket chatroom

### Core Implementation

`main.go`: A minimal websocket chatroom implementation.

- Uses a global `Room` to manage all connected clients
- Clients are stored in a map
- Messages are broadcast through a channel (`room.broadcast`)
- Broadcasting is handled in a single goroutine (`room.run`)
- Each message is delivered via `client.conn.WriteMessage`

### Test Implementation

`main_test.go`: Integration test for WebSocket communication (written with AI assistance).

#### What the test does:
The test simulates a live chatroom environment using Go's built-in `httptest` package:
1. **Spins up an in-memory test server** running `wsHandler`.
2. **Establishes two client connections**: `alice` and `bob`.
3. **Simulates communication**: `alice` sends a `"hello"` message.
4. **Validates broadcasting**: Verifies that `bob` successfully receives Alice's message within a 2-second deadline.

### Post-Test Analysis & Learning Notes

> ⚠️ **Self-Correction & Note:** The following technical analysis was generated with AI assistance. I haven't had the time to deeply verify every single detail here yet. However, I am documenting this as my current working hypothesis, which I will personally verify in the next sprint.

#### What the AI Explains About These "Errors":
When running `go test`, the console sometimes outputs:
- `read error: read tcp... connection reset by peer`
- `read error: websocket: close 1006 (abnormal closure): unexpected EOF`

According to the architectural review:
1. **The Ghost Disconnection:** After the test successfully verifies that Bob received the message, the test script abruptly kills the client processes.
2. **Protocol Violation (1006):** Because the test clients disappear instantly without sending a standard WebSocket close frame, the `gorilla/websocket` library flags this as an `Abnormal Closure (1006)`.
3. **Resilience Verification:** Even with these brutal dropouts, the server's `defer` blocks successfully caught the anomalies, broke the loops safely, and executed client cleanup routines without freezing or panicking.
