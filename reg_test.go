package main

import (
	"log"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {
	s := []byte(`abcs
	asjdi
	asjdiojsd
	asjdio`)

	m, _ := regexp.Match("www", s)
	log.Println(m)
}
