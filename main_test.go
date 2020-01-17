package main

import (
	"fmt"
	"sort"
	"testing"
)

var (
	tcs = []struct {
		data       [][]int
		intersect  []int
		sum        []int
		difference []int
	}{{
		[][]int{},
		[]int{},
		[]int{},
		[]int{},
	}, {
		[][]int{{1, 2, 3}},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}, {
		[][]int{{1, 2, 3}, {2, 3, 4}},
		[]int{2, 3},
		[]int{1, 2, 3, 4},
		[]int{1},
	}, {
		[][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		[]int{3},
		[]int{1, 2, 3, 4, 5},
		[]int{1},
	}, {
		[][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}},
		[]int{},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1},
	}, {
		[][]int{{1, 1, 1, 3}, {2, 2, 2, 3}, {3, 3, 3, 3}, {4, 4, 4, 3}},
		[]int{3},
		[]int{1, 2, 3, 4},
		[]int{1},
	}, {
		[][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {7, 8, 9, 10}},
		[]int{},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3},
	}}

	calcTcs = []struct {
		input  string
		result []int
		err    bool
	}{{
		"[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]",
		[]int{1, 3, 4},
		false,
	}, {
		"[ SUM [ DIF a.txt b.txt c.txt ] ]",
		[]int{1},
		false,
	}, {
		"[ SUM a.txt b.txt c.txt ]",
		[]int{1, 2, 3, 4, 5},
		false,
	}, {
		"[ DIF [ SUM a.txt b.txt c.txt ] b.txt ]",
		[]int{1, 5},
		false,
	}, {
		"[ SUM a.txt b.txt c.txt ] ]",
		nil,
		true,
	}, {
		"[ SUM [ SUM wrongName.txt a.txt ] b.txt ]",
		nil,
		true,
	}, {
		"[ a.txt b.txt ]",
		nil,
		true,
	}, {
		"SUM a.txt b.txt",
		nil,
		true,
	}}
)

func TestIntersect(t *testing.T) {
	for _, v := range tcs {
		result := intersect(v.data)
		sort.Ints(result)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", v.intersect) {
			t.Errorf("TestIntersect fail: got <%v> want <%v>", result, v.intersect)
		}
	}
}
func TestSum(t *testing.T) {
	for _, v := range tcs {
		result := sum(v.data)
		sort.Ints(result)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", v.sum) {
			t.Errorf("TestSum fail: got <%v> want <%v>", result, v.sum)
		}
	}
}
func TestDifference(t *testing.T) {
	for _, v := range tcs {
		result := difference(v.data)
		sort.Ints(result)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", v.difference) {
			t.Errorf("TestDifference fail: got <%v> want <%v>", result, v.difference)
		}
	}
}
func TestParseExpression(t *testing.T) {
	for _, v := range calcTcs {
		result, err := parseExpression(v.input)
		if err != nil && !v.err {
			t.Errorf("TestCalcMoc fail: unexpected error: %v", err)
		}

		sort.Ints(result)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", v.result) {
			t.Errorf("TestCalcMoc fail: got <%v> want <%v>", result, v.result)
		}
	}

}
