package memory

import (
	"grail-participant-registry/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestItStoresParticipantsInMemory(t *testing.T) {
	repo := NewParticipantRepository()

	uuid := uuid.New().String()

	participant := domain.Participant{
		ID: domain.ParticipantID(uuid),
	}

	_ = repo.Add(participant)
	result, err := repo.FindByRef(domain.ParticipantReference(uuid))

	require.NoError(t, err)
	assert.Equal(t, participant.ID, result.ID)
}

func TestErrorsWhenFindingAParticipantThatIsNotStored(t *testing.T) {
	repo := NewParticipantRepository()
	_, err := repo.FindByRef("does_not_exist")
	assert.Equal(t, err, domain.ErrParticipantNotFound)
}

func TestItRemovesStoredParticipants(t *testing.T) {
	repo := NewParticipantRepository()
	participant := domain.Participant{
		ID: "test",
	}

	_ = repo.Add(participant)
	_ = repo.Remove("test")
	_, err := repo.FindByRef("test")

	require.Equal(t, domain.ErrParticipantNotFound, err)
}

func TestErrorsRemovingNonExistentParticipant(t *testing.T) {
	repo := NewParticipantRepository()
	err := repo.Remove("test")

	require.Error(t, err)
	require.Equal(t, err, domain.ErrParticipantNotFound)
}

func TestFindsAllStoredParticipants(t *testing.T) {
	repo := NewParticipantRepository()

	_ = repo.Add(domain.Participant{
		ID:   "participant_1",
		Name: "participant a",
	})
	_ = repo.Add(domain.Participant{
		ID:   "participant_2",
		Name: "participant c",
	})
	_ = repo.Add(domain.Participant{
		ID:   "participant_3",
		Name: "participant b",
	})

	result, _ := repo.FindAll()

	assert.Len(t, result, 3)
	assert.Exactly(t, result[0].Name, "participant a")
	assert.Exactly(t, result[1].Name, "participant b")
	assert.Exactly(t, result[2].Name, "participant c")
}

func TestParticipantRepository_Update(t *testing.T) {
	repo := NewParticipantRepository()

	p1 := domain.Participant{
		Reference: "participant_1",
		Name:      "participant a",
	}

	_ = repo.Add(p1)

	newDoB := time.Now()
	p1.Name = "updated name"
	p1.Phone = "updated phone"
	p1.DateOfBirth = newDoB
	err := repo.Update(p1)

	p2, err2 := repo.FindByRef("participant_1")

	assert.NoError(t, err)
	assert.NoError(t, err2)
	assert.Equal(t, p2.Name, "updated name")
	assert.Equal(t, p2.Phone, "updated phone")
	assert.Equal(t, p2.DateOfBirth, newDoB)
}

func TestParticipantRepository_Update_FailsWithInvalidRef(t *testing.T) {
	repo := NewParticipantRepository()

	err := repo.Update(
		domain.Participant{
			Reference: "invalid ref",
		},
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "participant not found")
}
