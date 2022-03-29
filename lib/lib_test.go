package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenShortUrl(t *testing.T) {
	slug := GenShortSlug()
	assert.Equal(t, 4, len(slug))
}
