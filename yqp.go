package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"mime/quotedprintable"
	"mime"
	"os"
	"strings"
)

func toQuotedPrintable(s string) (string, error) {
	var ac bytes.Buffer
	w := quotedprintable.NewWriter(&ac)
	_, err := w.Write([]byte(s))
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}
	return ac.String(), nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputFile := flag.String("i", "", "input file name. Add path to read from any location "+
		"than the script. Otherwise file name")
	outputFile := flag.String("o", "", "output file name with .txt extension")
	flag.Parse()
	if *outputFile == "" || *inputFile == "" {
		log.Fatal("file name cannot be empty")
	}
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	body := ""
	subject := []string{}
	finalText := ""
	scanner := bufio.NewScanner(file)
	str := []string{}
	for scanner.Scan() {
		str = strings.Split(scanner.Text(), ":")

		if len(str) > 1 {
			if str[0] == "To" {
				finalText += scanner.Text()
				finalText += "\n"
			} else if str[0] == "Subject" {
				subject = str
			} else {
				if len(subject) > 0 {
					encoded, _ := toQuotedPrintable(scanner.Text())
					body += encoded + "\n"
				} else {
					finalText += scanner.Text() + "\n"
				}
			}
		} else {
			encoded, _ := toQuotedPrintable(scanner.Text())
			body += encoded + "\n"
		}
	}
	
	finalText += subject[0] + ": "
	subEncoded := mime.BEncoding.Encode("UTF-8", strings.Join(subject[1:], ":"))
	if len(strings.Split(subEncoded, " ")) > 1 && strings.Join(subject[1:], ":") != subEncoded {
		for i, se := range strings.Split(subEncoded, " ") {
			if i > 0 {
				finalText += "  " + se
			} else {
				finalText += se
			}
			finalText += "\n"
		}
	} else {
		finalText += subEncoded + "\n"
	}

	finalText += "MIME-Version: 1.0\nContent-Type: text/plain; charset=UTF-8\nContent-Transfer-Encoding: quoted-printable\n"
	finalText += body
	fmt.Println(finalText)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*outputFile)
	check(err)
	defer f.Close()
	f.WriteString(finalText)
}
