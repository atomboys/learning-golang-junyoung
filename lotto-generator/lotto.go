package main

import "fmt"
import "math/rand"
import "time"
import "sort"

func main() {
	generatedNumbers := getLottoNumbers()

	printNumbers(generatedNumbers)
}

func printNumbers(generatedNumbers []int) {
	for i := 0; i < len(generatedNumbers); i++ {
		fmt.Println(generatedNumbers[i])
	}
}

// 로또 번호 생성
func getLottoNumbers() []int {
	// 랜덤 시드 설정
	rand.Seed(time.Now().UnixNano())

	generatedNumbers := make([]int, 6)

	for i := 0; i < len(generatedNumbers); i++ {
		generatedNumbers[i] = generateNumberExceptFor(generatedNumbers)
	}

	// 생성된 로또 번호를 내림차순 정렬
	sort.Ints(generatedNumbers)

	return generatedNumbers
}

// 랜덤 번호 생성
func generateNumberExceptFor(generatedNumbers []int) (generatedNumber int) {
	for true {
		generatedNumber = rand.Intn(45) + 1

		if !hasNumber(generatedNumbers, generatedNumber) {
			break
		}
	}

	return
}

// 배열에 특정 숫자가 있는지 확인
func hasNumber(generatedNumbers []int, targetNumber int) bool {
	for _, value := range generatedNumbers {
		if value == targetNumber {
			return true
		}
	}

	return false
}
