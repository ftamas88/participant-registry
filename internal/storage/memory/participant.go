package memory

import (
	"grail-participant-registry/internal/domain"
	"sort"
	"sync"
)

type ParticipantRepository struct {
	mu           sync.RWMutex
	participants map[domain.ParticipantID]domain.Participant
}

func NewParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{
		participants: make(map[domain.ParticipantID]domain.Participant),
	}
}

func (repo *ParticipantRepository) FindByID(participantID domain.ParticipantID) (domain.Participant, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	participant, ok := repo.participants[participantID]

	if !ok {
		return domain.Participant{}, domain.ErrParticipantNotFound
	}

	return participant, nil
}

func (repo *ParticipantRepository) FindAll() ([]domain.Participant, error) {
	participants := make([]domain.Participant, 0)

	repo.mu.RLock()
	for _, participant := range repo.participants {
		participants = append(participants, participant)
	}
	repo.mu.RUnlock()

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Name < participants[j].Name
	})

	return participants, nil
}

func (repo *ParticipantRepository) Add(participant domain.Participant) error {
	repo.mu.Lock()
	repo.participants[domain.ParticipantID(participant.ID)] = participant
	repo.mu.Unlock()

	return nil
}

func (repo *ParticipantRepository) Update(participant domain.Participant) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.participants[domain.ParticipantID(participant.ID)]
	if !ok {
		return domain.ErrParticipantNotFound
	}

	repo.participants[domain.ParticipantID(participant.ID)] = participant

	return nil
}

func (repo *ParticipantRepository) Remove(participantID domain.ParticipantID) error {
	repo.mu.Lock()
	_, ok := repo.participants[participantID]

	if !ok {
		repo.mu.Unlock()

		return domain.ErrParticipantNotFound
	}

	delete(repo.participants, participantID)
	repo.mu.Unlock()

	return nil
}
