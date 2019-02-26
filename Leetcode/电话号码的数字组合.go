package main

var dd map[string]string

func letterCombinations(digits string) []string {
	return nil
}

func main() {
	dd := make(map[string]string)
	dd["1"] = ""
	dd["2"] = "abc"
	dd["3"] = "def"
	dd["4"] = "ghi"
	dd["5"] = "jkl"
	dd["6"] = "mno"
	dd["7"] = "pqs"
	dd["8"] = "tuv"
	dd["9"] = "wxyz"
	letterCombinations("23")
}
