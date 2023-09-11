package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Username string
	TeamName string
	Ranking  int
	Score    int
}
