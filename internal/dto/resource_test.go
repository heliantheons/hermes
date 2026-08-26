package dto

import (
	"testing"
	"time"

	"github.com/heliantheon/hermes/internal/models"
)

func TestNewRelationshipResponseExposesMutationIdentifier(t *testing.T) {
	t.Parallel()

	createdAt := time.Date(2026, time.August, 26, 12, 0, 0, 0, time.UTC)
	response := NewRelationshipResponse(&models.Relationship{
		ID:          42,
		ServiceID:   "hermes",
		SubjectType: "user",
		SubjectID:   "user-1",
		Relation:    "viewer",
		ObjectType:  "resource",
		ObjectID:    "resource-1",
		CreatedAt:   createdAt,
	})

	if response.RelationshipID != 42 {
		t.Fatalf("relationship id = %d, want 42", response.RelationshipID)
	}
}
