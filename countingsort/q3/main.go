// 한반에 학생들 중 키가 특정 범위의 학생들만 출력 하시오
// 범위 값은 여러번 입력 받을 수 있습니다.
// 키는 소수점 1자리까지 주어집니다.

// 입력 예
// 140 170
// 180 190
// 200 210
// 160 180

package main

import "fmt"

func main() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "ddd1", Height: 173.4},
		{Name: "ddd2", Height: 164.5},
		{Name: "ddd3", Height: 178.8},
		{Name: "ddd4", Height: 154.2},
		{Name: "ddd5", Height: 188.8},
		{Name: "ddd6", Height: 209.8},
		{Name: "ddd7", Height: 197.7},
		{Name: "ddd8", Height: 164.8},
		{Name: "ddd9", Height: 164.8},
	}

	// O(N) 이나 입력이 한번일 경우에만 해당
	for i := 0; i < len(students); i++ {
		if students[i].Height >= 170.0 && students[i].Height < 180.0 {
			fmt.Println(students[i].Name, students[i].Height)
		}
	}

	for _, v := range students {
		if v.Height >= 170.0 && v.Height < 180.0 {
			fmt.Println(v.Name, v.Height)
		}
	}

	// counting sort를 이용하여 풀이
	// 사람의 키 이기 때문에 범위 지정이 가능 0 ~ 300.0
	var heightArr [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightArr[idx] = append(heightArr[idx], students[i].Name)
	}

	fmt.Println("140 ~ 170")
	//140 ~170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightArr[i] {
			fmt.Println("name", name, "height", float64(i)/10)
		}
	}
	fmt.Println("180 ~ 210")
	//180 ~210
	for i := 1800; i < 2100; i++ {
		for _, name := range heightArr[i] {
			fmt.Println("name", name, "height", float64(i)/10)
		}
	}
}
