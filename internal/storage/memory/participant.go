package memory

import (
	"grail-participant-registry/internal/domain"
	"sync"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type ParticipantRepository struct {
	mu           sync.RWMutex
	participants map[domain.ParticipantReference]domain.Participant
}

func NewParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{
		participants: make(map[domain.ParticipantReference]domain.Participant),
	}
}

func (repo *ParticipantRepository) FindByRef(ref domain.ParticipantReference) (domain.Participant, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	participant, ok := repo.participants[ref]
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

	_, ok := repo.participants[participant.Reference]
	if ok {
		logrus.Warnf("participant already exists with the following ref: [%s]", participant.Reference)

		return domain.ErrParticipantAlreadyExists
	}

	if participant.ID == "" {
		participant.ID = domain.ParticipantID(uuid.New().String())
	}

	repo.participants[participant.Reference] = participant

	return nil
}

func (repo *ParticipantRepository) Update(participant domain.Participant) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.participants[participant.Reference]
	if !ok {
		logrus.Warnf("participant does not exists with the following ref: [%s]", participant.Reference)

		return domain.ErrParticipantNotFound
	}

	repo.participants[participant.Reference] = participant

	return nil
}

func (repo *ParticipantRepository) Remove(ref domain.ParticipantReference) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.participants[ref]
	if !ok {
		logrus.Warnf("participant does not exists with the following ref: [%s]", ref)

		return domain.ErrParticipantNotFound
	}

	delete(repo.participants, ref)

	return nil
}
