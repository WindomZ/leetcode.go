package simplifypath

import "strings"

func simplifyPath(path string) string {
	ps := strings.Split(path, "/")
	paths := make([]string, 0, len(ps)) // as a stack
	for _, p := range ps {
		if p == "" || p == "." {
			continue
		}
		if p == ".." {
			if len(paths) != 0 {
				paths = paths[:len(paths)-1] // pop
			}
		} else {
			paths = append(paths, p) // push
		}
	}
	if len(paths) == 0 {
		return "/"
	}
	return "/" + strings.Join(paths, "/")
}
