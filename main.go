package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run(in io.Reader, out io.Writer) error {
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
