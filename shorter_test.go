package godiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToShort(t *testing.T) {

	f1, e := readFile("./testdata/f1.txt")
	assert.Nil(t, e)

	f1Shorter, e := readFile("./testdata/f1_shorter.txt")
	assert.Nil(t, e)

	o := IniDoc(f1, "f1.txt")
	_, e = ProcessDiff(o, f1Shorter, "f1_shorter.txt")
	assert.NotNil(t, e)
}
