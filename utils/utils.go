package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func ParseIDs(ids []string) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, 0, len(ids))

	for _, id := range ids {
		parsedID, err := uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("invalid id: %s", id)
		}

		result = append(result, parsedID)
	}

	return result, nil
}
