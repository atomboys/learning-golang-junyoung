package study

import (
	"fmt"
)

// Map 맵 실습
func Map() {
	// var 맵이름 map[키_자료형]값_자료형
	var a map[string]int

	//맵은 make 함수를 사용하여 공간을 할당해야 값을 넣을 수 있음
	a = make(map[string]int)

	// 맵을 생성하면서 값을 초기화
	a = map[string]int{
		"Hello": 10,
		"World": 20,
	}

	fmt.Println(a)

	// 맵에 데이터 저장하고 조회하기
	solarSystem := make(map[string]float32)

	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641

	fmt.Println(solarSystem["Earth"])

	// 맵에 존재하지 않는 키를 조회했을 때는 빈 값이 출력
	// 키 자료형이 int면 0, string이면 ""

	// 맵에 데이터 있는지 검사
	value, hasValue := solarSystem["Pluto"]
	fmt.Println(value, hasValue)

	// 맵 순회
	for key, value := range solarSystem {
		fmt.Println(key, value)
	}

	// 맵에서 데이터 삭제
	delete(solarSystem, "Mercury")
	fmt.Println(solarSystem["Mercury"])

	// 맵 안에 맵 만들기
	// map[키_자료형]map[키_자료형]값_자료형

}
