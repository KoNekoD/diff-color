package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiff(t *testing.T) {
	const a = `{
 "id": 1,
 "name": "foo",
 "created": "2021-01-01T00:00:00Z",
}`

	const b = `{
 "id": 2,
 "name": "foo",
              "created": "2021-01-01T00:00:00Z",
}`

	result := Diff(a, b)

	excepted := " {\n\x1b[0;31m- \"id\": 1,\x1b[0m\n\x1b[0;32m+ \"id\": 2,\x1b[0m\n  \"name\": \"foo\",\n\x1b[0;31m- \"created\": \"2021-01-01T00:00:00Z\",\x1b[0m\n\x1b[0;32m+              \"created\": \"2021-01-01T00:00:00Z\",\x1b[0m\n }"

	assert.Equal(t, excepted, result)
}
