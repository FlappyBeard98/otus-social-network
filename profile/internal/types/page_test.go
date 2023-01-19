package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPageResponseReturnsNotNil(t *testing.T) {
	pr := PageRequest{}
	items := make([]int, 0)
	act := NewPageResponse(&pr, items, 1)

	assert.NotNil(t, act)
}
