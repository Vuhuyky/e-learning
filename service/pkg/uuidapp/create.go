package uuidapp

import "github.com/google/uuid"

func Create() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
