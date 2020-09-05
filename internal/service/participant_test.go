package service

import (
	"fmt"
	"grail-participant-registry/internal/domain"
	"grail-participant-registry/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParticipantService(t *testing.T) {
	repo := &mocks.ParticipantRepository{}

	p := Participant{repository: repo}
	if got := NewParticipantService(repo); !reflect.DeepEqual(got, &p) {
		t.Errorf("NewParticipantService() = %v, want %v", got, &p)
	}
}

func TestParticipant_Create(t *testing.T) {
	repo := &mocks.ParticipantRepository{}
	repo.On(
		"Add",
		domain.Participant{
			Reference: "mocked-ref",
			Name:      "mocked-name",
		},
	).
		Once().
		Return(nil)

	srv := &Participant{
		repository: repo,
	}
	err := srv.Create(
		domain.Participant{
			Reference: "mocked-ref",
			Name:      "mocked-name",
		},
	)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestParticipant_Delete(t *testing.T) {
	repo := &mocks.ParticipantRepository{}
	repo.On(
		"Remove",
		domain.ParticipantReference("delete-ref-123"),
	).
		Once().
		Return(nil)

	srv := &Participant{
		repository: repo,
	}
	err := srv.Delete("delete-ref-123")
	repo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestParticipant_Get(t *testing.T) {
	repo := &mocks.ParticipantRepository{}
	repo.On(
		"FindByRef",
		domain.ParticipantReference("get-ref-123"),
	).
		Once().
		Return(
			domain.Participant{
				ID:        "mocked id",
				Name:      "mocked name",
				Reference: "get-ref-123",
				Phone:     "mocked phone",
				Address: domain.ParticipantAddress{
					StreetName: "mocked street",
				},
			},
			nil,
		)

	srv := &Participant{
		repository: repo,
	}
	p, err := srv.Get("get-ref-123")
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, domain.ParticipantID("mocked id"), p.ID)
	assert.Equal(t, domain.ParticipantReference("get-ref-123"), p.Reference)
	assert.Equal(t, "mocked name", p.Name)
	assert.Equal(t, "mocked phone", p.Phone)
	assert.Equal(
		t,
		domain.ParticipantAddress{
			StreetName: "mocked street",
		},
		p.Address,
	)
}

func TestParticipant_Update(t *testing.T) {
	repo := &mocks.ParticipantRepository{}
	repo.
		On("FindByRef", domain.ParticipantReference("update-ref-123")).
		Return(
			domain.Participant{
				ID:        "mocked id",
				Name:      "mocked name",
				Reference: "updated-ref-123",
				Phone:     "mocked phone",
				Address: domain.ParticipantAddress{
					StreetName: "mocked street",
				},
			},
			nil,
		)
	repo.On(
		"Update",
		domain.Participant{
			ID:        "mocked id",
			Name:      "updated name",
			Reference: "updated-ref-123",
			Phone:     "updated phone",
			Address: domain.ParticipantAddress{
				StreetName: "mocked street",
			},
		},
	).
		Once().
		Return(nil)

	srv := &Participant{
		repository: repo,
	}
	err := srv.Update(
		"update-ref-123",
		domain.Participant{
			Name:  "updated name",
			Phone: "updated phone",
		},
	)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestParticipant_Update_Not_Found(t *testing.T) {
	repo := &mocks.ParticipantRepository{}
	repo.
		On("FindByRef", domain.ParticipantReference("update-bad-ref-123")).
		Return(
			domain.Participant{},
			fmt.Errorf("mocked error"),
		)

	srv := &Participant{
		repository: repo,
	}
	err := srv.Update(
		"update-bad-ref-123",
		domain.Participant{},
	)
	repo.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mocked error")
}
