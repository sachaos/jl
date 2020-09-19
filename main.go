package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	hasHeader bool
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func Run() error {
	flag.BoolVar(&hasHeader, "header", false, "header")
	flag.Parse()

	if hasHeader {
		if err := WithHeader(os.Stdin, os.Stdout); err != nil {
			return err
		}

		return nil
	}

	if err := HeaderLess(os.Stdin, os.Stdout); err != nil {
		return err
	}

	return nil
}

func WithHeader(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	encoder := json.NewEncoder(out)

	var headers []string
	if scanner.Scan() {
		headers = Split(scanner.Text())
	} else {
		return fmt.Errorf("cannot scan")
	}

	for scanner.Scan() {
		obj := map[string]interface{}{}
		line := scanner.Text()
		values := Split(line)

		for i, header := range headers {
			obj[header] = values[i]
		}

		if err := encoder.Encode(obj); err != nil {
			return err
		}
	}

	return nil
}

func HeaderLess(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	encoder := json.NewEncoder(out)

	for scanner.Scan() {
		obj := map[string]interface{}{}
		line := scanner.Text()
		values := Split(line)

		obj["column"] = values

		if err := encoder.Encode(obj); err != nil {
			return err
		}
	}

	return nil
}
