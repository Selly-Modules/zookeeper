package zookeeper

import "strconv"

// GetFloat64Value ...
func GetFloat64Value(path string) float64 {
	s := GetStringValue(path)
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
