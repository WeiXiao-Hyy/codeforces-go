// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/500/C
// https://codeforces.com/problemset/status/500/problem/C
func Test_cf500C(t *testing.T) {
	testCases := [][2]string{
		{
			`3 5
1 2 3
1 3 2 3 1`,
			`12`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf500C)
}
