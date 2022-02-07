package split

import "strings"

func Split(str, sep string) (result []string) {
	result = make([]string, 0, strings.Count(str, sep)+1)
	i := strings.Index(str, sep)
	for i > -1 {
		result = append(result, str[:i])
		str = str[i+len(sep):]
		i = strings.Index(str, sep)
	}
	result = append(result, str)
	return
}
