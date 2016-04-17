package study

// Variable 변수 선언 및 값 할당
func Variable() {
	var _ int
	var _ string

	var _ int = 10
	var _ string = "Junyoung"

	// 자료형 생략
	var _ = 10
	var _ = "Junyoung Lee"
	// var address // 컴파일 에러

	// 짧은 선언
	// _ := 10
	// _ := "Junyoung Lee"

	// 변수 여러 개 선언하기
	var _, _ int = 30, 50
	var _, _ = 10, "Junyoung Lee"

	// _, _, _, _ := 1, 3.4, "Hello, world!", false

	// 병렬 할당(Parallel assignement)
	var _, _ int
	var _ int
	_, _, _ = 10, 20, 5

	var (
		_, _ int = 30, 50
		_, _     = 10, "Maria"
		_, _     = 3.14, "Hello"
	)

	// 사용하지 않는 변수 및 패키지 처리
	/*
	  Go 언어는 선언만 하고 사용하지 않는 변수가 있으면 컴파일 에러 발생
	  import 하고 사용하지 않는 패키지가 있을 때도 컴파일 에러 발생
	*/

	_ = 10 // 사용하지 않는 변수로 인한 컴파일 에러 방지

	// import의 경우
	// _import _ "time" // 패키지 명 앞에 _ 를 붙여 컴파일 에러 방지

}
