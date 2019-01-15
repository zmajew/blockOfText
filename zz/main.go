package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type word struct {
	strOfInterest string
	blocs         []block
	fileName      string
}

type block struct {
	blckUri string
	repeats int
}

var Word word

func main() {

	filePath := "tekst.txt"

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	wor := "This"

	re := regexp.MustCompile(`(?i)\b` + wor + `\b`)

	scanner := bufio.NewScanner(f)
	var rjec string
	scanner.Split(crunchSplitFunc)
	for scanner.Scan() {
		rjec = scanner.Text()
		match := re.FindAllStringSubmatch(rjec, -1)
		fmt.Println(rjec)
		fmt.Println(len(match))
		fmt.Println("----")

		// if rjec == "FILE 1.0" {

		// }
	}

}
func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	re := regexp.MustCompile(`\nFILE`)
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := re.FindStringIndex(string(data)); len(i) > 0 && i[0] >= 0 {
		fmt.Println(i[0])
		return i[0] + 1, data[0:i[0]], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
