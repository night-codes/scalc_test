package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	ret, err := parseExpression(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Println(err)
	}

	sort.Ints(ret)
	for _, v := range ret {
		fmt.Println(v)
	}
}

// sum returns a slice that contains the values of all the input slices
func sum(data [][]int) []int {
	m := map[int]bool{}
	ret := []int{}
	for k := range data {
		for _, v := range data[k] {
			if !m[v] {
				ret = append(ret, v)
				m[v] = true
			}
		}
	}
	return ret
}

// intersect returns a slice of values that are present in all of the input slices
func intersect(data [][]int) []int {
	m := map[int]bool{}
	for k := range data {
		tmp := map[int]bool{}
		for _, v := range data[k] {
			if m[v] || k == 0 {
				tmp[v] = true
			}
		}
		m = tmp
	}

	ret := []int{}
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

// difference of first set and the rest ones
func difference(data [][]int) []int {
	m := map[int]bool{}
	for k := range data {
		for _, v := range data[k] {
			if k == 0 {
				m[v] = true
			} else if m[v] {
				delete(m, v)
			}
		}
	}

	ret := []int{}
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

// parse expression from the string and return result as int set
func parseExpression(s string) ([]int, error) {
	var fn func(data [][]int) []int
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s := strings.TrimSpace(s[1 : len(s)-1])
		if strings.HasPrefix(s, "INT ") {
			fn = intersect
		} else if strings.HasPrefix(s, "SUM ") {
			fn = sum
		} else if strings.HasPrefix(s, "DIF ") {
			fn = difference
		}
		if fn != nil {
			ret, err := parseFuncArgs(s[4:])
			if err != nil {
				return nil, err
			}
			return fn(ret), nil
		}
	}
	return nil, errors.New("Parser error")
}

// parseFuncArgs parse arguments next of command INT, SUM or DIF and return as slice of int sets
func parseFuncArgs(s string) ([][]int, error) {
	s = strings.TrimSpace(s)
	ret := [][]int{}
	open := 0
	arg := ""

	for k, v := range s {
		if v == '[' {
			open++
		} else if v == ']' {
			open--
			if open == 0 && len(arg) > 0 {
				cr, err := parseExpression("[" + arg + "]")
				if err != nil {
					return nil, err
				}
				ret = append(ret, cr)
				arg = ""
			} else if open < 0 {
				return nil, errors.New("Parser error")
			}
		} else if open == 0 && (v == ' ' && len(arg) > 0 || k == len(s)-1) {
			if v != ' ' {
				arg += string(v)
			}
			cr, err := getFromFile(arg)
			if err != nil {
				return nil, err
			}
			ret = append(ret, cr)
			arg = ""
		} else if v != ' ' || open > 0 {
			arg += string(v)
		}
	}

	return ret, nil
}

// get set from file by fileName
func getFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var perline int
	var nums []int
	for {
		if _, err := fmt.Fscanf(file, "%d\n", &perline); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		nums = append(nums, perline)
	}
	return nums, nil
}
