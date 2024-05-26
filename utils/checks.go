package utils

func Check(err error) int {
	if err != nil {
        panic(err)
    } else {
		return 0
	}
}