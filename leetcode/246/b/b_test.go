// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`"12:01"`, `"12:44"`, 
			`1`,
		},
		{
			`"20:00"`, `"06:00"`, 
			`40`,
		},
		{
			`"00:00"`, `"23:59"`, 
			`95`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, numberOfRounds, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-246/problems/the-number-of-full-rounds-you-have-played/
