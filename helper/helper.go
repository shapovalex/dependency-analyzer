package helper

import (
	"bufio"
	"fmt"
	"os"
)

func DeleteIntersectionElements(baseElements []string, searchElements []string) []string {
	intermediateResult := map[string]bool{}
	for _, v := range baseElements {
		intermediateResult[v] = true
	}
	for _, v := range searchElements {
		delete(intermediateResult, v)
	}
	var result []string
	for k, _ := range intermediateResult {
		result = append(result, k)
	}
	return result
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
