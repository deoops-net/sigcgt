package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var pid string

func init() {
	flag.StringVar(&pid, "p", "1", "pid")
}

func main() {
	flag.Parse()
	out, err := SigCgtCMD(pid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		parseSigCgt(out),
	)

}

func SigCgtCMD(p string) ([]byte, error) {
	cmd := exec.Command("cat", fmt.Sprintf("/proc/%s/status", p))
	return cmd.Output()
}

func parseSigCgt(s []byte) string {
	cursor := 0
	cgt := []string{}
	line := SigCgtLine(s)
	sigBits := bytes.TrimSpace([]byte(
		strings.Split(line, ":")[1],
	))

	mm, _ := hex.DecodeString(string(sigBits))
	// decimal to binary bits
	// then range from low to find singal mappings
	for i := len(mm) - 1; i >= 0; i-- {
		bits := leftFullBinBits(
			fmt.Sprintf("%b", mm[i]),
		)
		for n := len(bits) - 1; n >= 0; n-- {
			cursor++
			if string(bits[n]) == "1" && Signal[strconv.Itoa(cursor)] != "" {
				cgt = append(cgt, Signal[strconv.Itoa(cursor)])
			}
		}

	}
	return strings.Join(cgt, ",")
}

func SigCgtLine(s []byte) string {
	lines := strings.Split(string(s), "\n")
	for _, v := range lines {
		m, _ := regexp.MatchString("SigCgt", v)
		if m {
			return v
		}
	}
	return ""
}

func trimZeroLeft(s []byte) []byte {
	for i, b := range s {
		if b != '0' {
			return s[i:]
		}
	}
	return []byte("")
}

func leftFullBinBits(s string) string {
	if len(s) == 8 {
		return s
	}
	return strings.Repeat("0", 8-len(s)) + s

}
