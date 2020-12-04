package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Field struct {
	name  string
	regex *regexp.Regexp
}

func Day4Part2() {
	var counter int
	content, err := ioutil.ReadFile("assets/input-4")
	if err != nil {
		log.Fatal(err)
	}
	fields := []Field{
		{"byr", nil},
		{"iyr", nil},
		{"eyr", nil},
		{"hgt", nil},
		{"hcl", regexp.MustCompile(`^#[a-f0-9]{6}$`)},
		{"ecl", regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)$`)},
		{"pid", regexp.MustCompile(`^\d{9}$`)},
	}

	// split into string passport
	passports := strings.Split(string(content), "\n\n")

	for _, passport := range passports {

		isValid := true
		// replace new lines with space
		passport := strings.ReplaceAll(passport, "\n", " ")


		// split into key value pairs "eyr:2020"
		splitBySpace := strings.Split(passport, " ")

		// if entries is less than 7 we can skip
		if len(splitBySpace) < 7 {
			continue
		}

		//split key value into array where a[0] = key and a[1] = value
		passportFields := make(map[string]string)
		for _, v := range splitBySpace {
			keyVal := strings.Split(v, ":")
			passportFields[keyVal[0]] = strings.ReplaceAll(keyVal[1], " ", "")
		}

		for _, field := range fields {

			val, isPresent := passportFields[field.name]

			if !isPresent || !isValidField(val, field) {
				// passport does not contain field or does not match regular expression
				isValid = false
				break
			}
		}

		if isValid {
			counter++
		}

	}
	fmt.Printf("solution day 4 part 2: %d\n", counter)
}

func isValidField(val string, with Field) bool {
	switch with.name {
	case "byr": // check if birth year is between 1920 and 2002
		{
			val, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			return val >= 1920 && val <= 2002
		}
	case "iyr": // check if issue year is between 2010 and 2020
		{
			val, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			return val >= 2010 && val <= 2020
		}
	case "eyr": // check if expiration year is between 2020 and 2030
		{
			val, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			return val >= 2020 && val <= 2030
		}
	case "hgt": // check if height as suffix cm (between 150 and 193) or in (between 59 and 76)
		{
			if strings.HasSuffix(val, "cm") {
				replaced := strings.ReplaceAll(val, "cm", "")
				val, err := strconv.Atoi(replaced)
				if err != nil {
					log.Fatal(err)
				}
				return val >= 150 && val <= 193
			} else if strings.HasSuffix(val, "in") {
				replaced := strings.ReplaceAll(val, "in", "")
				val, err := strconv.Atoi(replaced)
				if err != nil {
					log.Fatal(err)
				}
				return val >= 59 && val <= 76
			}
			return false
		}
	default: // match value with regular expression
		return with.regex.MatchString(val)
	}
}
