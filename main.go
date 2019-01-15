package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type word struct {
	strOfInterest string
	blocks        []block
	fileName      string
}

type block struct {
	blckURI string
	repeats int
}

var Word word

func main() {
	var filePath string
	switch len(os.Args) {
	case 1:
		fmt.Println("type a file name: ")
		fmt.Scan(&filePath)
		fmt.Println("type the word to find in file: ")
		fmt.Scan(&Word.strOfInterest)
	case 2:
		filePath = os.Args[1]
		fmt.Println("type the word to find in file: ")
		fmt.Scan(&Word.strOfInterest)
	case 3:
		filePath = os.Args[1]
		Word.strOfInterest = os.Args[2]
	default:
		fmt.Println(`Number of imput arguments not 3.`)
		fmt.Println("Format should be: app_name <file> <word>")
		fmt.Println(`If <word> has spaces put the backslash before space: 'open\ the\ door'`)
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	destFilePath := "recLog.txt"
	fOut, err := os.OpenFile(destFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error creating file:", destFilePath)
		log.Fatalln(err)
	}
	defer fOut.Close()

	m := findWord(f, Word)

	err = record(fOut, m)
	if err != nil {
		log.Println("unable to record in to the file", destFilePath)
		log.Fatalln(err)
	}
}

func findWord(r io.Reader, w word) word {
	reWord := regexp.MustCompile(`(?i)\b` + w.strOfInterest + `\b`)
	reURI := regexp.MustCompile(`\w+:(\/?\/?)[^\s]+`)
	scanner := bufio.NewScanner(r)
	var text string
	scanner.Split(crunchSplitFunc)

	for scanner.Scan() {
		var b block
		text = scanner.Text()
		mtchWord := reWord.FindAllStringSubmatch(text, -1)
		b.blckURI = reURI.FindString(text)
		b.repeats = len(mtchWord)
		w.blocks = append(w.blocks, b)
	}

	return w
}

func record(r io.Writer, w word) error {
	log1 := "-------------------------\n" + time.Now().Format(time.UnixDate) + "\nResults fo finding of word '" + w.strOfInterest + "':\n"
	_, err := fmt.Fprintf(r, log1)
	if err != nil {
		return err
	}

	for i, v := range w.blocks {
		a := i + 1
		log2 := "Blokc number " + strconv.Itoa(a)
		_, err := fmt.Fprintln(r, log2)
		if err != nil {
			return err
		}

		log3 := "At URI: " + v.blckURI
		_, err = fmt.Fprintln(r, log3)
		if err != nil {
			return err
		}

		log4 := "Word '" + w.strOfInterest + "' was found " + strconv.Itoa(v.repeats) + " times."
		_, err = fmt.Fprintln(r, log4)
		if err != nil {
			return err
		}
	}
	return nil
}

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	re := regexp.MustCompile(`\nFILE [\d]+\.[\d]+`)
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := re.FindStringIndex(string(data)); len(i) > 0 && i[0] >= 0 {
		return i[0] + 1, data[0:i[0]], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
