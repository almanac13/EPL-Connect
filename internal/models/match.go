package models

import "time"

type Match struct {
  ID         string    `json:"id"`
  HomeTeamID string    `json:"homeTeamId"`
  AwayTeamID string    `json:"awayTeamId"`
  MatchDate  time.Time `json:"matchDate"`
  Score      string    `json:"score"`
}