package word1

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated"){
		t.Errorf(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak"){
		t.Errorf(`IsPalindrome("kayak") = false`)
	}

}

func TestNonPalidrome(t *testing.T){
	if IsPalindrome("kyoka"){
		t.Errorf(`IsPalindrome("kyoka") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}