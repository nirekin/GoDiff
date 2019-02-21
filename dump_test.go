package godiff

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDump(t *testing.T) {

	f1, e := readFile("./testdata/f1.yaml")
	assert.Nil(t, e)
	o := IniDoc(f1, "f1.yaml")

	f2, e := readFile("./testdata/f2.yaml")
	assert.Nil(t, e)

	f3, e := readFile("./testdata/f3.yaml")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, f2, "f2.yaml")
	assert.Nil(t, e)

	d, e = ProcessDiff(d, f3, "f3.yaml")
	assert.Nil(t, e)

	e = d.dump("./", "out.txt")
	assert.Nil(t, e)

	log.Println("--------------------------------------------------------------")
	d.dumpOut()

}
