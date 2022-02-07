package implode

func Implode(arr []string, str string) (result string) {
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			result = result + str + arr[i]
		} else {
			result = arr[i]
		}
	}
	return
}
