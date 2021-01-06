package uuid_identify

import "github.com/google/uuid"

func NewUuid() (newUuid string) {
	id := uuid.New()
	return id.String()
}
