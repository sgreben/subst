package main

import (
	"fmt"
	"strings"
)

type enumVar struct {
	Choices []string
	Value   string
}

// Set implements the flag.Value interface.
func (so *enumVar) Set(v string) error {
	for _, c := range so.Choices {
		if strings.EqualFold(c, v) {
			so.Value = c
			return nil
		}
	}
	return fmt.Errorf(`"%s" must be one of [%s]`, v, strings.Join(so.Choices, " "))
}

func (so *enumVar) String() string {
	return so.Value
}
