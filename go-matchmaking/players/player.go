package main

import (
  "github.com/bciulkin/go-test-value-provider"
  "net/url"
  "strconv"
)

type Player struct {
  Id string     `json:"id"`
  Rating int    `json:"rating"`
  Role string   `json:"role"`
}

func NewPlayer() Player {
  return Player{Id: value_provider.String(), Rating: value_provider.IntNM(1100, 1600), Role: randomRole()}
}

func NewPlayerParams() url.Values {
  params := url.Values{}
  params.Add("id", value_provider.String())
  params.Add("rating", strconv.Itoa(value_provider.IntNM(1100, 1600)))
  params.Add("role", randomRole())

  return params
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
