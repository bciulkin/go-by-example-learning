package main

import (
  "github.com/bciulkin/go-test-value-provider"
  "strconv"
  "fmt"
)

type Player struct {
  Id string     `json:"id"`
  Rating int    `json:"rating"`
  Role string   `json:"role"`
}

type Match struct {
	Team1 []Player `json:"team1"`
	Team2 []Player `json:"team2"`
}

func NewPlayer() Player {
  return Player{Id: value_provider.String(), Rating: value_provider.IntNM(1100, 1600), Role: randomRole()}
}

func NewPlayerJsonString() []byte {
  return []byte(fmt.Sprintf(`{"id": %s, "rating": %s, "role": %s}`, value_provider.String(), strconv.Itoa(value_provider.IntNM(1100, 1600)), randomRole()))
}

type Role int

const (
  Tank = iota
  Dps
  Support
)

func randomRole() string {
  roles := [3]string{"Tank", "Dps", "Support"}
  return roles[value_provider.IntN(3)]
}

func roleToString(role Role) string {
  switch role {
  case Tank:
    return "Tank"
  case Dps:
    return "Dps"
  case Support:
    return "Support"
  }
  return "unknown"
}
