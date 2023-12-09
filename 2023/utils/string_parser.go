package utils

import (
	"errors"
	"strconv"
	"strings"
)

type IStringParser interface {
	Trim(cutset string)
	Split(sep string) []*StringParser
	ToInt() (int, error)
}

type StringParser string

func NewStringParser(s string) IStringParser {
	sp := StringParser(s)
	return &sp
}

func (s *StringParser) Trim(cutset string) {
	*s = StringParser(strings.Trim(string(*s), cutset))
}

func (s *StringParser) ToInt() (int, error) {
	if num, err := strconv.Atoi(string(*s)); err == nil {
		return num, nil
	}

	return 0, errors.New("could not convert string to int")
}

func (s *StringParser) Split(sep string) []*StringParser {
	substrings := strings.Split(string(*s), sep)
	sSplit := make([]*StringParser, len(substrings))

	for i, ss := range substrings {
		ssAsStringParser := StringParser(ss)
		sSplit[i] = &ssAsStringParser
	}

	return sSplit

}
