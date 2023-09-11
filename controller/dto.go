package controller

type UserRequestAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PlayerRequest struct {
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	Ranking  int    `json:"ranking"`
	Score    int    `json:"score"`
}
