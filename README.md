# leetcode.go

> My LeetCode Problem Solutions(Golang).

[![Build Status](https://travis-ci.org/WindomZ/leetcode.go.svg?branch=master)](https://travis-ci.org/WindomZ/leetcode.go)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/leetcode.go/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/leetcode.go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/leetcode.go)](https://goreportcard.com/report/github.com/WindomZ/leetcode.go)

## Algorithms

| # | Problem | Solution | Difficulty | Total Repetitions | Single Repetition Duration | LeetCode Run Time |
|---| ----- | :--------: | :----------: | ----------: | ----------: | ----------: |
|6|[ZigZag Conversion][Algorithms-6]|[WindomZ][Algorithms-6-Go]|Medium|30000000|51.5 ns/op| 9 ms |
|5|[Longest Palindromic Substring][Algorithms-5]|[WindomZ][Algorithms-5-Go]|Medium|50000000|28.7 ns/op| 9 ms |
|4|[Median of Two Sorted Arrays][Algorithms-4]|[WindomZ][Algorithms-4-Go]|Hard|300000000|5.74 ns/op| 32 ms |
|3|[Longest Substring Without Repeating Characters][Algorithms-3]|[WindomZ][Algorithms-3-Go]|Medium|50000000|36.3 ns/op| 6 ms |
|2|[Add Two Numbers][Algorithms-2]|[WindomZ][Algorithms-2-Go]|Medium|20000000|79.4 ns/op| 29 ms |
|1|[Two Sum][Algorithms-1]|[WindomZ][Algorithms-1-Go]|Easy|20000000|113 ns/op| 6 ms |

## Contributing

Welcome to pull requests, report bugs, suggest ideas and discuss on 
[issues page](https://github.com/WindomZ/leetcode.go/issues).

```bash
git clone https://github.com/WindomZ/leetcode.go "$PROJECT_PATH"
cd "$PROJECT_PATH"
go test -v --run=".*" ./algorithms/...
go test -test.bench=".*" ./algorithms/...
```

If you like it then you can put a :star: on it.

[Algorithms-6-Go]:algorithms/zigzag_conversion/convert.go
[Algorithms-6]:https://leetcode.com/problems/zigzag-conversion/
[Algorithms-5-Go]:algorithms/longest_palindromic_substring/longestPalindrome.go
[Algorithms-5]:https://leetcode.com/problems/longest-palindromic-substring/
[Algorithms-4-Go]:algorithms/median_of_two_sorted_arrays/findMedianSortedArrays.go
[Algorithms-4]:https://leetcode.com/problems/median-of-two-sorted-arrays/
[Algorithms-3-Go]:algorithms/longest_substring_without_repeating_characters/lengthOfLongestSubstring.go
[Algorithms-3]:https://leetcode.com/problems/longest-substring-without-repeating-characters/
[Algorithms-2-Go]:algorithms/add_two_numbers/addTwoNumbers.go
[Algorithms-2]:https://oj.leetcode.com/problems/add-two-numbers/
[Algorithms-1-Go]:algorithms/two_sum/twoSum.go
[Algorithms-1]:https://oj.leetcode.com/problems/two-sum/
