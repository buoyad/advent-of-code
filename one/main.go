package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/* Merry Christmas! */

func main() {
	filename := flag.String("captcha", "input", "Path to input file")
	flag.Parse()
	if filename == nil || *filename == "" {
		fmt.Println("Error: must supply a captcha!\nSee 'one --help' for usage")
		return
	}

	captcha, err := readCaptcha(*filename)
	if err != nil {
		fmt.Printf("Error in reading captcha: %s", err)
		return
	}

	nums, err := processCaptcha(captcha)
	if err != nil {
		fmt.Printf("Error in processing captcha: %s", err)
	}

	answer := accumulate(nums)
	fmt.Printf("The answer to the captcha is %d\n", answer)
}

// Get the entire captcha as a string
func readCaptcha(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", nil
	}
	defer file.Close()
	res, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// Process the captcha string into an array of integers
func processCaptcha(captcha string) ([]int, error) {
	res := []int{}
	for _, char := range captcha {
		if char == '\n' {
			continue
		}
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}

// Sum all numbers that equal the next element
// wrap so the last num is next to the first
func accumulate(nums []int) int {
	answer := 0
	totalLength := len(nums)
	for i := 0; i < totalLength; i++ {
		if nums[i] == nums[(i+1)%totalLength] {
			answer += nums[i]
		}
	}
	return answer
}
