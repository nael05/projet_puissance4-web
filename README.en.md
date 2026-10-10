# Connect 4 Web (Puissance 4 Web)

**Individual School Project**
This project was completed individually as part of my computer science studies at Ynov Campus.

## Description
This is a lightweight web implementation of the classic strategy game **Connect 4**, developed with **Go** (Golang) for the backend and **HTML/CSS** for the frontend.

The project provides a playable grid where two players can compete locally. The server handles all game logic: move validation, win detection (horizontal, vertical, diagonal), and draws.

## Features
- **Classic Gameplay**: 6x7 grid respecting standard rules.
- **Two-Player Mode**: Local multiplayer with alternating turns (🔴 Red vs 🟡 Yellow).
- **Game Logic (Server)**:
  - Automatic win detection (4 tokens in a row).
  - Draw detection (full board).
  - Invalid move prevention (full columns).
- **Interactive User Interface**:
  - Hover effects on selection buttons.
  - Visual indicators for the current turn.
  - **Fireworks animation** upon winning.
  - Animated background.
- **Reset Function**: Ability to restart the game immediately at the end of a match.

## Tech Stack
- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3
- **Templating**: Go `html/template` package

## Project Structure
```text
projet_puissance4-web/
├── main.go            # Entry point: HTTP server and game logic
├── go.mod             # Go module definition
├── templates/         # HTML templates (home.html, game.html)
├── style/             # CSS files (style.css)
└── static/            # Static resources (images, gifs)
```

## Installation and Launch
1. Make sure you have Go installed.
2. Clone this repository.
3. Run the server from the root directory:
   ```bash
   go run main.go
   ```
4. Open your browser and go to `http://localhost:8080`.
