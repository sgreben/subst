package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var definitions = make(map[string]string)
var nonzeroExit = false
var newline = []byte{'\n'}

var quiet bool
var unknownWarn = true
var unknownKey = enumVar{
	Choices: []string{
		unknownValueEmpty,
		unknownValueIgnore,
		unknownValueError,
	},
}

const (
	unknownValueEmpty  = "empty"
	unknownValueIgnore = "ignore"
	unknownValueError  = "error"
)

func parseDefinition(d string) (key, value string, err error) {
	i := strings.IndexRune(d, '=')
	if i < 0 {
		err = fmt.Errorf(`"%s" should have the format KEY=VALUE`, d)
		return
	}
	key, value = d[:i], d[i+1:]
	return
}

func init() {
	log.SetOutput(os.Stderr)
	flag.BoolVar(&quiet, "q", false, "suppress all logs")
	flag.Var(&unknownKey, "unknown", "handling of unknown keys, one of [ignore empty error] (default ignore)")
	flag.Parse()
	if quiet {
		log.SetOutput(ioutil.Discard)
	}
	for _, d := range flag.Args() {
		k, v, err := parseDefinition(d)
		if err != nil {
			log.Println(err)
			nonzeroExit = true
			continue
		}
		definitions[k] = v
	}
}

func subst(k string) (out string) {
	if v, ok := definitions[k]; ok {
		return v
	}
	switch unknownKey.Value {
	case unknownValueEmpty:
		out = ""
	case unknownValueIgnore:
		out = "$" + k
	case unknownValueError:
		log.Printf(`undefined key: $%s`, k)
		os.Exit(1)
	default:
		out = "$" + k
	}
	if unknownWarn {
		log.Printf(`undefined key: $%s, using value "%s"`, k, out)
	}
	return
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		_, err := os.Stdout.WriteString(os.Expand(line, subst))
		if err != nil {
			log.Println(err)
		}
		_, err = os.Stdout.Write(newline)
		if err != nil {
			log.Println(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Println(err)
		nonzeroExit = true
	}

	if nonzeroExit {
		os.Exit(1)
	}
}
