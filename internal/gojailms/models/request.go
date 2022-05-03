package models

type CreateRequest struct {
	LocalUser    Users   `json:"user"`
	LocalVisitor Visitor `json:"visitor"`
}
