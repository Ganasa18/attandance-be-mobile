package helper

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func FormatArraySplitByComma(values []interface{}) string {
	var valueStrings []string
	for _, value := range values {
		valueStrings = append(valueStrings, fmt.Sprintf("%v", value))
	}
	return strings.Join(valueStrings, ", ")
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
