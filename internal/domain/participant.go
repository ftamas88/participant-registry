package domain

import (
	"errors"
	"time"
)

//go:generate mockery -name=ParticipantRepository
type ParticipantRepository interface {
	FindByID(participantRef ParticipantReference) (Participant, error)
	FindAll() ([]Participant, error)
	Add(participant Participant) error
	Update(participant Participant) error
	Remove(participantRef ParticipantReference) error
}

type ParticipantID string
type ParticipantReference string

type Participant struct {
	ID          ParticipantID        `json:"id"`
	Reference   ParticipantReference `json:"reference"`
	Name        string               `json:"name"`
	DateOfBirth time.Time            `json:"date_of_birth"`
	Phone       string               `json:"phone"`
	Address     ParticipantAddress   `json:"address"`
}

type ParticipantAddress struct {
	AddressType        string   `json:"address_type,omitempty"`
	Department         string   `json:"department,omitempty"`
	SubDepartment      string   `json:"sub_department,omitempty"`
	StreetName         string   `json:"street_name,omitempty"`
	BuildingNumber     string   `json:"building_number,omitempty"`
	PostCode           string   `json:"post_code,omitempty"`
	TownName           string   `json:"town_name,omitempty"`
	CountrySubDivision string   `json:"country_sub_division,omitempty"`
	Country            string   `json:"country,omitempty"`
	AddressLine        []string `json:"address_line,omitempty"`
}

var (
	ErrParticipantNotFound      = errors.New("participant not found")
	ErrParticipantAlreadyExists = errors.New("participant already exists in the system")
)
