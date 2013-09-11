package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObj(t *testing.T) {

	obj := "hello objx"
	o := New(obj)
	assert.Equal(t, o.Obj(), obj)

}
