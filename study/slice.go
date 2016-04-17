package study

import "fmt"

// Slice 슬라이스 실습
func Slice() {
	// 슬라이스 생성 및 할당
	var a []int = make([]int, 5) // make 함수로 int 형에 길이가 5인 슬라이스에 공간할당
	var b = make([]int, 5)       // 슬라이스를 선언할 때 자료형과 [] 생략
	c := make([]int, 5)          // 슬라이스를 선언할 때 var 키워드, 자료형과 [] 생략

	a = []int{1, 2, 4, 8, 16, 32}
	b = []int{
		1,
		3,
		5,
		7,
		9, // 여러줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	// make([]자료형, 길이, 용량)
	c = make([]int, 5, 10)

	// 슬라이스의 길이
	_ = len(a)
	_ = len(b)

	// 슬라이스의 용량 (Capacity)
	_ = cap(c)

	// 용량이 길이보다 크더라도 길이를 벗어난 인덱스에는 접근할 수 없음

	// 슬라이스에 값 추가하기
	a = append(a, 1, 5, 9)

	// 슬라이스는 배열과 달리 레퍼런스 타입
	// 배열은 대입하면 값이 복사되고, 슬라이스는 포인트가 복사됨

	// 슬라이스 복사하기 copy(target, source)
	// target의 길이만큼만 복사됨
	copy(b, a)

	// 부분 슬라이스 만들기
	b = a[0:5]

	fmt.Println(a[0:3])
	fmt.Println(b[1:3])
	fmt.Println(a[2:5])

	// 부분 슬라이스 만들면서 용량 설정
	b = a[0:4:10] // 0~3까지 슬라이스하고 용량을 10으로 설정
}
