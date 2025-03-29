package models

import "time"

type Match struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	HomeTeam    string    `json:"homeTeam" gorm:"not null"`
	AwayTeam    string    `json:"awayTeam" gorm:"not null"`
	MatchDate   time.Time `json:"matchDate" gorm:"not null"`
	HomeGoals   int       `json:"homeGoals" gorm:"default:0"`
	AwayGoals   int       `json:"awayGoals" gorm:"default:0"`
	YellowCards int       `json:"yellowCards" gorm:"default:0"`
	RedCards    int       `json:"redCards" gorm:"default:0"`
	ExtraTime   int       `json:"extraTime" gorm:"default:0"` // en minutos
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
