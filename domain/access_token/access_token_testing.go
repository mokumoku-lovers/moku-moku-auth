import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationHours, "expiration time should be 24 hours")
}

