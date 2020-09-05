package memory

import (
	"grail-participant-registry/internal/domain"
	"sync"

	"github.com/sirupsen/logrus"
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
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	participants := make([]domain.Participant, 0)

	for _, participant := range repo.participants {
		participants = append(participants, participant)
	}

	/* Optional
	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Name < participants[j].Name
	})
	*/

	return participants, nil
}

func (repo *ParticipantRepository) Add(participant domain.Participant) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.participants[participant.ID]
	if !ok {
		logrus.Warnf("participant already exists with the following ref: [%s]", participant.Reference)

		return domain.ErrParticipantAlreadyExists
	}

	repo.participants[participant.ID] = participant

	return nil
}

func (repo *ParticipantRepository) Update(participant domain.Participant) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.participants[participant.ID]
	if !ok {
		logrus.Warnf("participant does not exists with the following ref: [%s]", participant.Reference)

		return domain.ErrParticipantNotFound
	}

	repo.participants[participant.ID] = participant

	return nil
}

func (repo *ParticipantRepository) Remove(participantID domain.ParticipantID) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.participants[participantID]
	if !ok {
		logrus.Warnf("participant does not exists with the following id: [%s]", participantID)

		return domain.ErrParticipantNotFound
	}

	delete(repo.participants, participantID)

	return nil
}
