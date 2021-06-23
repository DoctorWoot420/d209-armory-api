package domain

import (
	"time"

	"github.com/doctorwoot420/d209-armory-api/internal/domain"
)

// Character represents a Diablo II character.
type Character struct {
	ID         string               `json:"d2s_id"`
	D2s        *domain.D2sCharacter `json:"d2s"`
	LastParsed time.Time            `json:"last_parsed"`
}
