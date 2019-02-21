package godiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelfDiff(t *testing.T) {

	f1, e := readFile("./testdata/f1.txt")
	assert.Nil(t, e)

	o := IniDoc(f1, "f1.txt")

	d, e := ProcessDiff(o, f1, "f1.txt")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 5)
	assert.Equal(t, d[0].Content, "F1_line_1")
	assert.Equal(t, d[1].Content, "F1_line_2")
	assert.Equal(t, d[2].Content, "F1_line_3")
	assert.Equal(t, d[3].Content, "F1_line_4")
	assert.Equal(t, d[4].Content, "F1_line_5")
	assert.Equal(t, d[0].Origin, "f1.txt")
	assert.Equal(t, d[1].Origin, "f1.txt")
	assert.Equal(t, d[2].Origin, "f1.txt")
	assert.Equal(t, d[3].Origin, "f1.txt")
	assert.Equal(t, d[4].Origin, "f1.txt")
}
