package utils

import (
	"os"
	"strings"
)

func LoadSchemaLocation() string {
	schemaLocation := os.Getenv("SCHEMA_LOCATION")
	if !strings.HasSuffix(schemaLocation, "/") {
		schemaLocation = schemaLocation + "/"
	}

	return schemaLocation
}
