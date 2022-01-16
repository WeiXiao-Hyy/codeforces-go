// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`["lc","cl","gg"]`, 
			`6`,
		},
		{
			`["ab","ty","yt","lc","cl","ab"]`, 
			`8`,
		},
		{
			`["cc","ll","xx"]`, 
			`2`,
		},
		{
			`["dd","aa","bb","dd","aa","dd","bb","dd","aa","cc","bb","cc","dd","cc"]`,
			`22`,
		},
		
	}
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, longestPalindrome, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-69/problems/longest-palindrome-by-concatenating-two-letter-words/