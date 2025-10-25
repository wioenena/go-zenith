package http

import "strings"

func GetFirstPathSegment(path string) string {
	if len(path) <= 1 {
		return path
	}

	idx := strings.IndexByte(path[1:], '/')
	if idx == -1 {
		return path
	}

	return path[:idx+1]
}
