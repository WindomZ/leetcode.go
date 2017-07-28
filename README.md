# leetcode.go

> My LeetCode Problem Solutions(Golang).

[![Build Status](https://travis-ci.org/WindomZ/leetcode.go.svg?branch=master)](https://travis-ci.org/WindomZ/leetcode.go)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/leetcode.go/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/leetcode.go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/leetcode.go)](https://goreportcard.com/report/github.com/WindomZ/leetcode.go)

## Algorithms

| # | Problem | Solution | Difficulty | Single Repetition Duration | LeetCode Run Time |
|---| ----- | :--------: | :----------: | ----------: | ----------: |
|8|[String to Integer (atoi)][Algorithms-8]|[WindomZ][Algorithms-8-Go]|Medium|[33.5 ns/op/5 test cases][Algorithms-8-Test]|3 ms|
|7|[Reverse Integer][Algorithms-7]|[WindomZ][Algorithms-7-Go]|Easy|[8.90 ns/op/3 test cases][Algorithms-7-Test]|3 ms|
|6|[ZigZag Conversion][Algorithms-6]|[WindomZ][Algorithms-6-Go]|Medium|[51.5 ns/op/1 test cases][Algorithms-6-Test]|9 ms|
|5|[Longest Palindromic Substring][Algorithms-5]|[WindomZ][Algorithms-5-Go]|Medium|[28.7 ns/op/1 test cases][Algorithms-5-Test]|9 ms|
|4|[Median of Two Sorted Arrays][Algorithms-4]|[WindomZ][Algorithms-4-Go]|Hard|[5.74 ns/op/1 test cases][Algorithms-4-Test]|32 ms|
|3|[Longest Substring Without Repeating Characters][Algorithms-3]|[WindomZ][Algorithms-3-Go]|Medium|[36.3 ns/op/1 test cases][Algorithms-3-Test]|6 ms|
|2|[Add Two Numbers][Algorithms-2]|[WindomZ][Algorithms-2-Go]|Medium|[79.4 ns/op/1 test cases][Algorithms-2-Test]|29 ms|
|1|[Two Sum][Algorithms-1]|[WindomZ][Algorithms-1-Go]|Easy|[113 ns/op/1 test cases][Algorithms-1-Test]|6 ms|

> NOTE: All tests should be run on the same machine. (include your code)

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

[Algorithms-8-Test]:algorithms/string_to_integer_atoi/myAtoi_test.go
[Algorithms-8-Go]:algorithms/string_to_integer_atoi/myAtoi.go
[Algorithms-8]:https://leetcode.com/problems/string-to-integer-atoi/
[Algorithms-7-Test]:algorithms/reverse_integer/reverse_test.go
[Algorithms-7-Go]:algorithms/reverse_integer/reverse.go
[Algorithms-7]:https://leetcode.com/problems/reverse-integer/
[Algorithms-6-Test]:algorithms/zigzag_conversion/convert_test.go
[Algorithms-6-Go]:algorithms/zigzag_conversion/convert.go
[Algorithms-6]:https://leetcode.com/problems/zigzag-conversion/
[Algorithms-5-Test]:algorithms/longest_palindromic_substring/longestPalindrome_test.go
[Algorithms-5-Go]:algorithms/longest_palindromic_substring/longestPalindrome.go
[Algorithms-5]:https://leetcode.com/problems/longest-palindromic-substring/
[Algorithms-4-Test]:algorithms/median_of_two_sorted_arrays/findMedianSortedArrays_test.go
[Algorithms-4-Go]:algorithms/median_of_two_sorted_arrays/findMedianSortedArrays.go
[Algorithms-4]:https://leetcode.com/problems/median-of-two-sorted-arrays/
[Algorithms-3-Test]:algorithms/longest_substring_without_repeating_characters/lengthOfLongestSubstring_test.go
[Algorithms-3-Go]:algorithms/longest_substring_without_repeating_characters/lengthOfLongestSubstring.go
[Algorithms-3]:https://leetcode.com/problems/longest-substring-without-repeating-characters/
[Algorithms-2-Test]:algorithms/add_two_numbers/addTwoNumbers_test.go
[Algorithms-2-Go]:algorithms/add_two_numbers/addTwoNumbers.go
[Algorithms-2]:https://oj.leetcode.com/problems/add-two-numbers/
[Algorithms-1-Test]:algorithms/two_sum/twoSum_test.go
[Algorithms-1-Go]:algorithms/two_sum/twoSum.go
[Algorithms-1]:https://oj.leetcode.com/problems/two-sum/
