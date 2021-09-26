package models

// Error is a general error response struct
type Error struct {
	Code      int
	ErrorCode string
	Message   string
	Status    string
}
