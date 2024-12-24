package pkg

import (
	"github.com/google/uuid"
	"strings"
)

func CreateUuid() string {
	id := uuid.New()
	hexString := strings.ReplaceAll(id.String(), "-", "")
	return hexString
}
