package main

import (
	"bufio"
	"fmt"
	_ "io/ioutil"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Lotto struct {
	drawNo, bonusNo int
	numbers         [6]int
}

func main() {
	generatedNumbers := getLottoNumbers()

	lottoNumbers := readLottoData()

	fmt.Println(lottoNumbers)

	printNumbers(generatedNumbers)
}

func readLottoData() (lottoNumbers []Lotto) {
	lottoNumbers = make([]Lotto, 0, 700)

	file, err := os.OpenFile(
		"lotto-data.txt",
		os.O_RDONLY,
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for i := 0; ; i++ {
		var lotto Lotto

		_, err := fmt.Fscanf(reader, "%d %d %d %d %d %d %d %d\n",
			&lotto.drawNo,
			&lotto.numbers[0],
			&lotto.numbers[1],
			&lotto.numbers[2],
			&lotto.numbers[3],
			&lotto.numbers[4],
			&lotto.numbers[5],
			&lotto.bonusNo)

		if err != nil && err.Error() == "EOF" {
			fmt.Println(err)
			return
		}

		lottoNumbers = append(lottoNumbers, lotto)
	}
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
