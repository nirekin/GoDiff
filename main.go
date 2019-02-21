package godiff

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
)

type (
	ResolvedDoc []ResolvedLine

	ResolvedLine struct {
		Content string
		Origin  string
	}
)

func (r ResolvedDoc) dump(path string, name string) error {

	maxLength := 0
	for _, v := range r {
		l := len(v.Content)
		if l > maxLength {
			maxLength = l
		}
	}

	f := filepath.Join(path, name)
	fout, err := os.Create(f)
	if err != nil {
		return err
	}
	defer fout.Close()

	w := bufio.NewWriter(fout)

	padLength := len(strconv.Itoa(maxLength))
	for i, v := range r {
		pad := strconv.Itoa(padLength - len(strconv.Itoa(i+1)) + maxLength + 3)
		s := fmt.Sprintf("%d|%-"+pad+"v|from: %s", i+1, v.Content, v.Origin)
		fmt.Fprintln(w, s)
	}
	return w.Flush()
}

func (r ResolvedDoc) dumpOut() {

	baseColors := []color.Attribute{color.FgRed, color.FgGreen, color.FgYellow, color.FgBlue, color.FgMagenta, color.FgCyan, color.FgWhite}
	colors := baseColors

	maxLength := 0
	sources := make(map[string]color.Attribute)

	for _, v := range r {
		l := len(v.Content)

		if l > maxLength {
			maxLength = l
		}

		if _, ok := sources[v.Origin]; !ok {
			if len(colors) == 0 {
				colors = baseColors
			}
			sources[v.Origin] = colors[0]
			colors = colors[1:]
		}
	}

	padLength := len(strconv.Itoa(maxLength))
	for i, v := range r {
		pad := strconv.Itoa(padLength - len(strconv.Itoa(i+1)) + maxLength + 3)
		s := fmt.Sprintf("%d|%-"+pad+"v|from: %s", i+1, v.Content, v.Origin)
		if c, ok := sources[v.Origin]; ok {
			color.New(c).Println(s)
		} else {
			fmt.Println(s)
		}
	}

}

// IniDoc creates the base version of a resolved document
func IniDoc(origin []string, originPath string) ResolvedDoc {
	r := make([]ResolvedLine, 0)

	for _, v := range origin {
		r = append(r, ResolvedLine{v, originPath})
	}
	return r
}

func ProcessDiff(origin ResolvedDoc, update []string, updatePath string) (ResolvedDoc, error) {
	r := make([]ResolvedLine, 0)

	lenO := len(origin)
	lenV := len(update)
	if lenO > lenV {
		return r, fmt.Errorf("The update (%d lines) cannot be shorter than the origin (%d lines)", lenV, lenO)
	}

	cLine := make(chan ResolvedLine)
	exit := make(chan bool)

	go merge(cLine, exit, origin, update, updatePath)

	for {
		select {
		case <-exit:
			return r, nil
		case l := <-cLine:
			r = append(r, l)
		}
	}
}

func merge(cLine chan ResolvedLine, exit chan bool, origin ResolvedDoc, update []string, updatePath string) {
	// No more lines to process, then we exit
	if len(update) == 0 {
		exit <- true
	}

	if len(origin) > 0 {
		// No changes, matching lines
		if origin[0].Content == update[0] {
			// The source of the line stays the same
			cLine <- origin[0]
			merge(cLine, exit, origin[1:], update[1:], updatePath)
		} else {
			// check if the original line can be located further  into the version
			located := make(chan bool)
			go checkNext(located, origin[0], update)

			for {
				select {
				case l := <-located:
					if l {
						// The original line has been located further
						// It's an insertion
						cLine <- ResolvedLine{update[0], updatePath}

						merge(cLine, exit, origin, update[1:], updatePath)
					} else {
						// The original line has not been located further
						// It's a replacement
						cLine <- ResolvedLine{update[0], updatePath}

						merge(cLine, exit, origin[1:], update[1:], updatePath)
					}
				}
			}
		}
	} else {
		// We just have remaining line into the version
		// Then they should be added
		for _, v := range update {
			cLine <- ResolvedLine{v, updatePath}
		}
		// Everything has been added then all is done
		exit <- true
	}
}

func checkNext(located chan bool, origin ResolvedLine, update []string) {

	for _, v := range update {
		if origin.Content == v {
			located <- true
		}
	}
	located <- false
}

func readFile(path string) ([]string, error) {

	r := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		return r, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		return r, err
	}
	return r, nil
}
