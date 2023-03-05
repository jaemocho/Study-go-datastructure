// 알파벳 소문자로 이루어진 문자열 중 가장 많이 나오는 알파벳 캐릭터를 출력하시오

package main

import "fmt"

func main() {
	str := "asdfasdfhjlwejkljsadfkljdsaf"
	// rs := []rune(str)

	// sorted := make([]rune, 26)

	// for i := 0; i < len(rs); i++ {
	// 	sorted[rs[i]-'a']++
	// }

	// fmt.Println(sorted)

	var count [26]int

	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0

	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}

	fmt.Printf("Max char : %c count: %d \n", maxCh, maxCount)

}
