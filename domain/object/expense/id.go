package expense

import (
	"github.com/google/uuid"
)

type ID string

func newID() (ID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return ID(""), err
	}

	return ID(uid.String()), nil
}
