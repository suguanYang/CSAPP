package number_presentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.Equal(t, Umax64()>>1, uint64(Tmax64()))
	assert.Equal(t, Tmin64(), Tmax64()+1)
	assert.Equal(t, uint64(Tmin64()<<1), Umin64())
}
