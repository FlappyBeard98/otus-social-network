package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProfilesPageReturnsNotNil(t *testing.T) {
	sut := ProfilesRequest{}
	act := sut.ReadProfilesPage()
	assert.NotNil(t, act)
}

func TestReadProfilesTotalReturnsNotNil(t *testing.T) {
	sut := ProfilesRequest{}
	act := sut.ReadProfilesTotal()
	assert.NotNil(t, act)
}

func TestProfilesRequestCreateRequestReturnsNotNil(t *testing.T) {
	sut := ProfilesRequest{}
	act, _ := sut.CreateRequest(host)
	assert.NotNil(t, act)
}
