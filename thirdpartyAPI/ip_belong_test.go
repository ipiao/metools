package thirdapi

import "testing"
import "github.com/stretchr/testify/assert"

func TestIpBelong(t *testing.T) {
	res, err := AliIPBelongToJSON("183.162.11.243")
	t.Log(res)
	assert.Nil(t, err)
}
