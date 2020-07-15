package word2

import "unicode"

func IsPalindrome(s string) bool{
	var letters []rune = make([]rune, 0, len(s))

	// Letter case is ignored, as are not-letters
	for _, r := range s{
		if unicode.IsLetter(r){
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters{
		if letters[i] != letters[len(letters) - 1 - i] {
			return false
		}
	}
	return true


}
