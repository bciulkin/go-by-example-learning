package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"sync"
	"time"
	"github.com/gorilla/websocket"
	"fmt"
)

// Upgrader to upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketClient struct {
	conn *websocket.Conn
	id string
}
// Matchmaker manages the pool of players and creates matches
type Matchmaker struct {
	PlayerPool []Player
	Clients    map[*websocket.Conn]Player
	mu         sync.Mutex
}

var matchmaker = &Matchmaker{
	PlayerPool: []Player{},
	Clients:    make(map[*websocket.Conn]Player),
}

// adds a player to the pool and creates a match if we have 10 players
func (m *Matchmaker) AddPlayer(conn *websocket.Conn, player Player) (*Match, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Add player to the pool and map it to the WebSocket connection
	m.PlayerPool = append(m.PlayerPool, player)
	m.Clients[conn] = player

	// If we have 10 players, create a match
	if len(m.PlayerPool) >= 10 {
		return m.createMatch(), true
	}
	return nil, false
}

// balances teams and creates a new match
func (m *Matchmaker) createMatch() *Match {
	// Sort players by rating (highest to lowest)
	sort.Slice(m.PlayerPool, func(i, j int) bool {
		return m.PlayerPool[i].Rating > m.PlayerPool[j].Rating
	})

	// Split players into two teams, balancing by rating
	var team1, team2 []Player
	team1Rating, team2Rating := 0, 0

	for _, player := range m.PlayerPool[:10] {
		if team1Rating <= team2Rating {
			team1 = append(team1, player)
			team1Rating += player.Rating
		} else {
			team2 = append(team2, player)
			team2Rating += player.Rating
		}
	}

	// Remove the 10 players used to create this match from the pool
	m.PlayerPool = m.PlayerPool[10:]

	return &Match{
		Team1: team1,
		Team2: team2,
	}
}

// handles new WebSocket connections
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Listen for messages from the client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		// Parse the player info from JSON
		var player Player
		if err := json.Unmarshal(message, &player); err != nil {
			log.Println("Invalid JSON:", err)
			continue
		}

		// Add the player to the pool and check if a match can be created
		match, matchCreated := matchmaker.AddPlayer(conn, player)
		if matchCreated {
			// Notify all 10 players in the match via their WebSocket connections
			matchmaker.mu.Lock()
			for clientConn, p := range matchmaker.Clients {
				if isInMatch(p, match) {
					if err := clientConn.WriteJSON(match); err != nil {
						log.Println("Error sending match to player:", err)
					}
				}
			}
			matchmaker.mu.Unlock()
		} else {
			// Notify the player they have been added to the pool
			responseMessage := fmt.Sprintf("Player %s added to the pool.", player.Id)
			if err := conn.WriteMessage(websocket.TextMessage, []byte(responseMessage)); err != nil {
				log.Println("Error sending response:", err)
			}
		}
	}
}

// checks if the player is part of the match
func isInMatch(player Player, match *Match) bool {
	for _, p := range match.Team1 {
		if p.Id == player.Id {
			return true
		}
	}
	for _, p := range match.Team2 {
		if p.Id == player.Id {
			return true
		}
	}
	return false
}

// Player represents a player info sent by client
type Player struct {
	Id string    `json:"id"`
	Rating int    `json:"rating"`
	Role   string `json:"role"`
}

// Match contains two balanced teams
type Match struct {
	Team1 []Player `json:"team1"`
	Team2 []Player `json:"team2"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/player", handleWebSocket)
	log.Println("Matchmaking WebSocket server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
