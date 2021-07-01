package dgplus

import "strings"

func containsString(arr []string, str string, ignoreCase bool) bool {
	for _, k := range arr {
		if ignoreCase {
			if e := strings.EqualFold(k, str); !e {
				continue
			} else {
				return true
			}
		} else {
			if k != str {
				continue
			} else {
				return true
			}
		}
	}
	return false
}
