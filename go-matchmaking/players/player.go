package main

import (
  "github.com/bciulkin/go-test-value-provider"
)

type Player struct {
  Id string     `json:"id"`
  Rating int    `json:"rating"`
  Role string   `json:"role"`
}

func NewPlayer() Player {
  return Player{Id: value_provider.String(), Rating: value_provider.IntNM(1100, 1600), Role: randomRole()}
}

type Role int

const (
  Tank = iota
  Dps
  Support
)

func randomRole() string {
  roles := [3]string{"Tank", "Dps", "Support"}
  return roles[value_provider.IntN(2)]
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
