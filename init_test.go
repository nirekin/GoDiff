package godiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	f1, e := readFile("./testdata/f1.txt")
	assert.Nil(t, e)

	o := IniDoc(f1, "f1.txt")
	assert.Equal(t, len(o), 5)
	assert.Equal(t, o[0].Content, "F1_line_1")
	assert.Equal(t, o[1].Content, "F1_line_2")
	assert.Equal(t, o[2].Content, "F1_line_3")
	assert.Equal(t, o[3].Content, "F1_line_4")
	assert.Equal(t, o[4].Content, "F1_line_5")
}
