package godiff

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func check(t *testing.T, d ResolvedDoc, i int, c string, o string) {
	assert.Equal(t, d[i].Content, c)
	if o == "" {
		o = strings.ToLower(c[:2])
		assert.Equal(t, d[i].Origin, o)
	} else {
		assert.Equal(t, d[i].Origin, o)
	}
}

func TestSimpleUpdate(t *testing.T) {

	f, e := readFile("./testdata/f1.txt")
	assert.Nil(t, e)
	o := IniDoc(f, "f1")

	fo, e := readFile("./testdata/f2.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f2")

	assert.Nil(t, e)
	assert.Equal(t, len(d), 5)
	checkF1ToF2(t, d, "f1", "f2")
}

func checkF1ToF2(t *testing.T, d ResolvedDoc, o string, tr string) {
	check(t, d, 0, "F1_line_1", o)
	check(t, d, 1, "F2_line_2_updated", tr)
	check(t, d, 2, "F1_line_3", o)
	check(t, d, 3, "F1_line_4", o)
	check(t, d, 4, "F1_line_5", o)
}

func TestSimpleAddition(t *testing.T) {

	f, e := readFile("./testdata/f2.txt")
	assert.Nil(t, e)
	o := IniDoc(f, "f2")

	fo, e := readFile("./testdata/f3.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f3")

	assert.Nil(t, e)
	assert.Equal(t, len(d), 7)
	checkF2ToF3(t, d, "f2", "f3")
}

func checkF2ToF3(t *testing.T, d ResolvedDoc, o string, tr string) {
	check(t, d, 0, "F1_line_1", o)
	check(t, d, 1, "F2_line_2_updated", o)
	check(t, d, 2, "F1_line_3", o)
	check(t, d, 3, "F1_line_4", o)
	check(t, d, 4, "F1_line_5", o)
	check(t, d, 5, "F3_line_6_added", tr)
	check(t, d, 6, "F3_line_7_added", tr)
}

func TestSimpleUpdateAddAddition(t *testing.T) {

	f, e := readFile("./testdata/f3.txt")
	assert.Nil(t, e)
	o := IniDoc(f, "f3")

	fo, e := readFile("./testdata/f4.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f4")

	assert.Nil(t, e)
	assert.Equal(t, len(d), 9)
	checkF3ToF4(t, d, "f3", "f4")
}

func checkF3ToF4(t *testing.T, d ResolvedDoc, o string, tr string) {
	check(t, d, 0, "F1_line_1", o)
	check(t, d, 1, "F2_line_2_updated", o)
	check(t, d, 2, "F4_line_3_updated", tr)
	check(t, d, 3, "F1_line_4", o)
	check(t, d, 4, "F1_line_5", o)
	check(t, d, 5, "F3_line_6_added", o)
	check(t, d, 6, "F3_line_7_added", o)
	check(t, d, 7, "F4_line_8_added", tr)
	check(t, d, 8, "F4_line_9_added", tr)
}

func TestSimpleInsert(t *testing.T) {

	f, e := readFile("./testdata/f4.txt")
	assert.Nil(t, e)
	o := IniDoc(f, "f4")

	fo, e := readFile("./testdata/f5.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f5")

	assert.Nil(t, e)
	assert.Equal(t, len(d), 10)
	checkF4ToF5(t, d, "f4", "f5")
}

func checkF4ToF5(t *testing.T, d ResolvedDoc, o string, tr string) {
	check(t, d, 0, "F1_line_1", o)
	check(t, d, 1, "F2_line_2_updated", o)
	check(t, d, 2, "F4_line_3_updated", o)
	check(t, d, 3, "F1_line_4", o)
	check(t, d, 4, "F5_inserted", tr)
	check(t, d, 5, "F1_line_5", o)
	check(t, d, 6, "F3_line_6_added", o)
	check(t, d, 7, "F3_line_7_added", o)
	check(t, d, 8, "F4_line_8_added", o)
	check(t, d, 9, "F4_line_9_added", o)

}

func TestSimpleLongInsert(t *testing.T) {

	f, e := readFile("./testdata/f5.txt")
	assert.Nil(t, e)
	o := IniDoc(f, "f5")

	fo, e := readFile("./testdata/f6.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f6")

	assert.Nil(t, e)
	assert.Equal(t, len(d), 13)
	checkF5ToF6(t, d, "f5", "f6")
}

func checkF5ToF6(t *testing.T, d ResolvedDoc, o string, tr string) {
	check(t, d, 0, "F1_line_1", o)
	check(t, d, 1, "F2_line_2_updated", o)
	check(t, d, 2, "F4_line_3_updated", o)
	check(t, d, 3, "F1_line_4", o)
	check(t, d, 4, "F5_inserted", o)
	check(t, d, 5, "F1_line_5", o)
	check(t, d, 6, "F6_insertion_l1", tr)
	check(t, d, 7, "F6_insertion_l2", tr)
	check(t, d, 8, "F6_insertion_l3", tr)
	check(t, d, 9, "F3_line_6_added", o)
	check(t, d, 10, "F3_line_7_added", o)
	check(t, d, 11, "F4_line_8_added", o)
	check(t, d, 12, "F4_line_9_added", o)
}

func TestSuperDiff(t *testing.T) {

	f, e := readFile("./testdata/f6.txt")
	assert.Nil(t, e)
	o := IniDoc(f, "f6")

	fo, e := readFile("./testdata/f7.txt")
	assert.Nil(t, e)

	d, e := ProcessDiff(o, fo, "f7")

	assert.Nil(t, e)
	assert.Equal(t, len(d), 27)
	checkF6ToF7(t, d, "f6", "f7")
}

func checkF6ToF7(t *testing.T, d ResolvedDoc, o string, tr string) {
	check(t, d, 0, "F7_insertion_l1", tr)
	check(t, d, 1, "F7_insertion_l2", tr)
	check(t, d, 2, "F7_insertion_l3", tr)
	check(t, d, 3, "F7_insertion_l4", tr)
	check(t, d, 4, "F7_insertion_l5", tr)
	check(t, d, 5, "F7_insertion_l6", tr)
	check(t, d, 6, "F1_line_1", o)
	check(t, d, 7, "F2_line_2_updated", o)
	check(t, d, 8, "F4_line_3_updated", o)
	check(t, d, 9, "F1_line_4", o)
	check(t, d, 10, "F5_inserted", o)
	check(t, d, 11, "F1_line_5", o)
	check(t, d, 12, "F6_insertion_l1", o)
	check(t, d, 13, "F7_insertion_l7", tr)
	check(t, d, 14, "F7_insertion_l8", tr)
	check(t, d, 15, "F7_insertion_l9", tr)
	check(t, d, 16, "F6_insertion_l2", o)
	check(t, d, 17, "F6_insertion_l3", o)
	check(t, d, 18, "F3_line_6_added", o)
	check(t, d, 19, "F7_updated", tr)
	check(t, d, 20, "F4_line_8_added", o)
	check(t, d, 21, "F4_line_9_added", o)
	check(t, d, 22, "F7_addition_l1", tr)
	check(t, d, 23, "F7_addition_l2", tr)
	check(t, d, 24, "F7_addition_l3", tr)
	check(t, d, 25, "F7_addition_l4", tr)
	check(t, d, 26, "F7_addition_l5", tr)
}

func TestIncrementalMegaDiff(t *testing.T) {

	f1, _ := readFile("./testdata/f1.txt")
	f2, _ := readFile("./testdata/f2.txt")
	f3, _ := readFile("./testdata/f3.txt")
	f4, _ := readFile("./testdata/f4.txt")
	f5, _ := readFile("./testdata/f5.txt")
	f6, _ := readFile("./testdata/f6.txt")
	f7, _ := readFile("./testdata/f7.txt")

	d := IniDoc(f1, "f1")
	d, e := ProcessDiff(d, f2, "f2")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 5)
	checkF1ToF2(t, d, "", "")

	d, e = ProcessDiff(d, f3, "f3")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 7)
	checkF2ToF3(t, d, "", "")

	d, e = ProcessDiff(d, f4, "f4")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 9)
	checkF3ToF4(t, d, "", "")

	d, e = ProcessDiff(d, f5, "f5")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 10)
	checkF4ToF5(t, d, "", "")

	d, e = ProcessDiff(d, f6, "f6")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 13)
	checkF5ToF6(t, d, "", "")

	d, e = ProcessDiff(d, f7, "f7")
	assert.Nil(t, e)
	assert.Equal(t, len(d), 27)
	checkF6ToF7(t, d, "", "")

}
