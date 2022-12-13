package art

import (
	"fmt"
	"os"
	"strings"
)

func Startascii(str, file string) string {
	as := Start(str, file)
	return as
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckFile(s string) string {
	sep := "--output="
	var res string
	for i, v := range s {
		if v == '-' {
			f := s[i : i+len(sep)]
			if sep == f {
				for _, j := range s[i+len(sep):] {
					res += string(j)
				}
				break
			}
		}
	}
	return res
}

func Maps(s []string) map[rune][]string {
	res := make(map[rune][]string)
	count := 1
	for i := ' '; i <= '~'; i++ {
		res[i] = s[count : count+8]
		count = count + 9
	}
	return res
}

func Massiv(s string) []string {
	a := strings.ReplaceAll(s, "\r", "\n")
	res := strings.Split(a, "\n")
	return res
}

func Print(s []string, maps map[rune][]string) string {
	// var res string
	var res string
	for _, val := range s {
		runes := []rune(val)
		str := []string{}
		for i := 0; i < 8; i++ {
			for _, v := range runes {
				str = append(str, maps[v][i])
			}
		}
		for i := 0; i < len(str); i++ {
			if i == 0 {
				res += string(str[i])
				// fmt.Print(str[i])
			} else if i%len(runes) == 0 && i != 0 {
				// fmt.Println()
				// fmt.Print(str[i])
				res += "\n"
				res += string(str[i])
			} else {
				// fmt.Print(str[i])
				res += string(str[i])
			}
		}
		// fmt.Println()
		res += "\n"
	}
	// fmt.Println(res)
	return res
}

func Scanfile(s string) []string {
	textres := []string{}
	var textword string
	for i, v := range s {
		if v == '\n' {
			// textword += string(v)
			textres = append(textres, textword)
			textword = ""
		} else if i == 0 || v == '\r' {
			continue
		} else {
			textword += string(v)
		}
	}
	if textword != "" {
		textres = append(textres, textword)
	}
	return textres
}

func Start(s, file string) string {
	// fmt.Println(s)
	if file == "standard" {
		text, err := os.ReadFile("standard.txt")
		// fmt.Println(text)
		Check(err)
		textStr := string(text)
		// fmt.Println(textStr)
		ScannerFile := Scanfile(textStr)
		// fmt.Println(1)
		// fmt.Println(s)
		maps := Maps(ScannerFile)
		// fmt.Println(s)
		// fmt.Println(1)
		slice := Massiv(s)
		// fmt.Println(1)
		a := Print(slice, maps)
		return a
	} else if file == "shadow" {
		text1, err1 := os.ReadFile("shadow.txt")
		// fmt.Println(text)
		Check(err1)
		textStr1 := string(text1)
		// fmt.Println(textStr)
		ScannerFile1 := Scanfile(textStr1)
		// fmt.Println(1)
		maps1 := Maps(ScannerFile1)
		// fmt.Println(1)
		slice1 := Massiv(s)
		a := Print(slice1, maps1)
		return a
	} else if file == "thinkertoy" {
		text2, err2 := os.ReadFile("thinkertoy.txt")
		// fmt.Println(text)
		Check(err2)
		textStr2 := string(text2)
		// fmt.Println(textStr)
		ScannerFile2 := Scanfile(textStr2)
		// fmt.Println(1)functions
		maps2 := Maps(ScannerFile2)
		// fmt.Println(1)
		slice2 := Massiv(s)
		a := Print(slice2, maps2)
		return a
	} else {
		return "Incorect: Banner. standard, shadow, thinkertoy"
	}
}

func verify(s string) string {
	if s == "shadow" {
		return "shadow.txt"
	} else if s == "standard" {
		return "standart.txt"
	} else if s == "thinkertoy" {
		return "thinkertoy.txt"
	}
	return s
}
