package godiff

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkLines(t *testing.T, d ResolvedDoc, from int, to int, o string) {

	if from == to {
		f := from - 1
		if !assert.Equal(t, d[f].Origin, o) {
			log.Printf("Failed on line %d, expected %s", f, o)
			log.Printf("Line content \"%s\", from \"%s\"", d[f].Content, d[f].Origin)
		}
	}

	for i := from - 1; i < to-1; i++ {
		if !assert.Equal(t, d[i].Origin, o) {
			log.Printf("Failed on line %d, expected %s", i, o)
			log.Printf("Line content \"%s\", from \"%s\"", d[i].Content, d[i].Origin)
		}
	}
}

func TestYaml(t *testing.T) {

	f, e := readFile("./testdata/f1.yaml")
	assert.Nil(t, e)
	o := IniDoc(f, "f1")

	fo, e := readFile("./testdata/f2.yaml")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f2")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 192)

	checkLines(t, d, 1, 20, "f1")
	checkLines(t, d, 21, 23, "f2")
	checkLines(t, d, 24, 91, "f1")
	checkLines(t, d, 92, 92, "f2")
	checkLines(t, d, 93, 173, "f1")
	checkLines(t, d, 174, 181, "f2")
	checkLines(t, d, 182, 192, "f1")

	fo, e = readFile("./testdata/f3.yaml")
	assert.Nil(t, e)

	d, e = ProcessDiff(d, fo, "f3")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 201)

	checkLines(t, d, 1, 20, "f1")
	checkLines(t, d, 21, 23, "f2")
	checkLines(t, d, 24, 26, "f3")
	checkLines(t, d, 27, 40, "f1")
	checkLines(t, d, 41, 43, "f3")
	checkLines(t, d, 44, 97, "f1")
	checkLines(t, d, 98, 98, "f2")
	checkLines(t, d, 99, 128, "f1")
	checkLines(t, d, 129, 131, "f3")
	checkLines(t, d, 132, 182, "f1")
	checkLines(t, d, 183, 190, "f2")
	checkLines(t, d, 191, 201, "f1")

}
