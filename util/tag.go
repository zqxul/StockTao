package util

import "strings"

// GetTagPart => get a part of Tag
func GetTagPart(tag, key string) (string, bool) {
	if tag == "" {
		return nil, false
	}
	parts := strings.Split(tag, ";")
	for _, part := range parts {
		if part == "" {
			continue
		}
		elements := strings.Split(part, ":")
		if len(elements) == 2 && elements[0] == key {
			return elements[1], true
		}
	}
	return nil, false
}
