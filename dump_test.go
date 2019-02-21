package godiff

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDumpYaml(t *testing.T) {

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

func TestDumpNeruda(t *testing.T) {

	f1, e := readFile("./testdata/v1.txt")
	assert.Nil(t, e)
	o := IniDoc(f1, "v1.txt")

	f2, e := readFile("./testdata/v2.txt")
	assert.Nil(t, e)

	f3, e := readFile("./testdata/v3.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, f2, "v2.txt")
	assert.Nil(t, e)

	d, e = ProcessDiff(d, f3, "v3.txt")
	assert.Nil(t, e)

	e = d.dump("./", "neruda.txt")
	assert.Nil(t, e)

	log.Println("--------------------------------------------------------------")
	d.dumpOut()

}
