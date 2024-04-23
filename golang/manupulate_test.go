package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestManipulate(t *testing.T) {
	case1 := Manipulate("a")
	case2 := Manipulate("ab")
	case3 := Manipulate("abc")
	case4 := Manipulate("aabb")

	fmt.Println(case1, case2, case3, case4)

	if !slices.Equal(case1, []string{"a"}) {
		t.Errorf("Manipulate(\"a\") != %s", case1)
	}

	if !slices.Equal(case2, []string{"ab", "ba"}) {
		t.Errorf("Manipulate(\"ab\") != %s", case2)
	}

	if !slices.Equal(case3, []string{"abc", "acb", "bac", "bca", "cab", "cba"}) {
		t.Errorf("Manipulate(\"abc\") != %s", case3)
	}

	if !slices.Equal(case4, []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}) {
		t.Errorf("Manipulate(\"aabb\") != %s", case4)
	}

}
