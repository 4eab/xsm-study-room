# Go Online Study Room

A multi-player online study room server built from scratch to explore Golang backend development in depth.

> **Note:** I entered this project with **zero prior experience in Go**. This repository documents my journey of learning the language's core syntax, runtime mechanics, and concurrency patterns entirely from scratch, using an iterative, hands-on approach.

This project is currently in a very early MVP stage and is being used as a continuous learning sandbox for Go concurrency, networking, and system design.

---

## Learning Philosophy

This project is part of a deliberate AI-assisted learning experiment. 

The entire journey of designing, coding, and debugging this project is done through intentional, high-friction collaboration with AI.

The goal is to explore how far a developer can go in learning a new language and building systems while using AI as a *thinking partner*, rather than an automation tool.

### Interaction Model

- ❌ No Copilot-style autocomplete or inline code generation
- ❌ No CLI agents generating or modifying the codebase
- ✔ Only conversational interaction via a standard chat interface

AI is used for:
- Explaining concepts
- Discussing design decisions
- Helping debug issues
- Suggesting possible implementations

However, I take responsibility for:
- Understanding the code
- Debugging runtime issues
- Making design decisions
- Maintaining the mental ownership of the system
- ...

Because Go is entirely new to me, I force myself to dismantle, analyze, and figure out exactly *what* each keyword achieves and *why* it works before integrating it.

---

## Quick Start

### 1. Initialize dependencies

```
go mod tidy
```

### 2. Run the server

```
go run main.go
```

Server will start on:
```
http://localhost:8080
```

WebSocket endpoint:

```
ws://localhost:8080/ws?username=YourName
```

---

## Current Features (MVP)

- [x] Basic WebSocket chatroom infrastructure
- [x] Multi-client connection management
- [x] Global message broadcasting

This MVP intentionally contains unresolved design and concurrency issues. These issues will be addressed in future iterations.

---

## Deep Dives & Learning Notes

As development progresses, I document key learnings and design decisions in:

* `CHANGELOG.md` (system evolution and technical changes)
* Personal blog (concept explanations and deeper reflections)
