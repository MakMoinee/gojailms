package models

type Users struct {
	UserID       int    `json:"userID"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword,omitempty"`
	UserType     int    `json:"userType"`
}
