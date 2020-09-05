package service

import (
	"encoding/json"
	"grail-participant-registry/internal/domain"

	"github.com/pkg/errors"
)

type Participant struct {
	repository domain.ParticipantRepository
}

func NewParticipantService(rep domain.ParticipantRepository) *Participant {
	return &Participant{
		repository: rep,
	}
}

func (srv *Participant) GetParticipant(ref string) (domain.Participant, error) {
	p, err := srv.repository.FindByID(domain.ParticipantReference(ref))
	if err != nil {
		return p, errors.Wrap(err, "unable to find participant")
	}

	return p, nil
}

func (srv *Participant) CreateParticipant(p domain.Participant) error {
	return srv.repository.Add(p)
}

func (srv *Participant) UpdateParticipant(ref string, p domain.Participant) error {
	pInDB, err := srv.repository.FindByID(domain.ParticipantReference(ref))
	if err != nil {
		return errors.Wrap(err, "unable to find participant")
	}

	// Overwrite the existing data if not all fields are sent back
	pb, _ := json.Marshal(p)
	err = json.Unmarshal(pb, &pInDB)
	if err != nil {
		return errors.Wrap(err, "unable to overwrite existing participant data")
	}

	return srv.repository.Update(pInDB)
}

func (srv *Participant) RemoveParticipant(ref string) error {
	return srv.repository.Remove(
		domain.ParticipantReference(ref),
	)
}
