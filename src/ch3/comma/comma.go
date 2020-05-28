// comma inserts commas in a non-negative decimal integer string.
package comma

func comma(number string) string {
	if len(number) <= 3 {
		return number
	}

	return comma(number[:len(number)-3]) + "," + number[len(number)-3:]
}

