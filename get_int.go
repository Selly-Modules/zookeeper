package zookeeper

import "strconv"

// GetIntValue ...
func GetIntValue(path string) int {
	s := GetStringValue(path)
	v, _ := strconv.Atoi(s)
	return v
}