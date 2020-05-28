package basename2

import "strings"

func basename(s string) string {
	if slash := strings.LastIndex(s, "/"); slash >= 0{
		s = s[slash+1:]
	}

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
