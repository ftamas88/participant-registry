package service

import "grail-participant-registry/internal/domain"

type Participant struct {
	repository domain.ParticipantRepository
}

func NewParticipantService(rep domain.ParticipantRepository) *Participant {
	return &Participant{
		repository: rep,
	}
}

func (p *Participant) Repository() domain.ParticipantRepository {
	return p.repository
}
