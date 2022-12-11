package pgn

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tag [2]string

func Split(input io.Reader) ([]Tag, []string, error) {

	s := bufio.NewScanner(input)

	s.Split(bufio.ScanLines)
	ln := 1

	tags := make([]Tag, 0)
	moves := make([]string, 0)
	var moveCtx = false //whether the scanner is in tags or in moves
	for s.Scan() {
		l := s.Text()

		if l != "" {

			if strings.HasPrefix(l, "[") {
				moveCtx = false
				tag := Tag{}
				i := strings.Index(l, " ")
				tag[0] = l[1:i]
				tag[1] = l[i : len(l)-1]
				tags = append(tags, tag)
			} else {
				moveCtx = true
				moves = append(moves, l)
			}

		} else {
			if moveCtx {
				tags = append(tags, Tag{})
				moves = append(moves, "---")
			}
		}

		ln++
	}

	return tags, moves, nil
}

// SplitPoints detects the line breaks in a pgn file and reports.
func SplitPoints(input io.Reader) ([]int, error) {

	s := bufio.NewScanner(input)

	s.Split(bufio.ScanLines)
	ln := 1

	splitPoints := make([]int, 0)
	var moveCtx = false //whether the scanner is in tags or in moves
	for s.Scan() {
		l := s.Text()

		if l != "" {

			if strings.HasPrefix(l, "[") {
				moveCtx = false
			} else {
				moveCtx = true

			}

		} else {
			if moveCtx {
				splitPoints = append(splitPoints, ln)
			}
		}

		ln++
	}

	return splitPoints, nil
}

// SplitAndSave splits file and save in the directory.
// New files named sequentially. a.txt -> [a.txt.1, a.txt.2 ... ]
func SplitAndSave(fileToBeChunked string) error {

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := "somebigfile_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			return err
		}

		// write/save buffer to disk
		err = ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		if err != nil {
			return err
		}

	}
	return nil
}

//https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
//	procstr("HELLO PEOPLE!", LOWER|REV|CAP)

const (
	UPPER = 1 // upper case
	LOWER = 2 // lower case
	CAP   = 4 // capitalizes
	REV   = 8 // reverses
)

func procstr(str string, conf byte) string {
	// reverse string
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	// query configs
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}
