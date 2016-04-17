package study

// NumberType 숫자 타입
func NumberType() {
	/*
	   Go 언어의 숫자 자료형 종류
	   uint8     8bit unsigned int
	   uint16
	   uint32
	   uint64
	   int8
	   int16
	   int32
	   int64
	   uint      32비트 시스템에서는 uint32, 64비트 시스템에서는 uint64
	   int       32비트 시스템에서는 int32, 64비트 시스템에서는 int64
	   uintptr   uint와 크기가 동일하며 포인터 저장 시 사용
	   float32   32비트 단정빌도 부동소수점, 7자리 정밀도 보장
	   float64   64비트 배정밀도 부동소수점, 15자리 정밀도 보장
	   complext64  float32 크기의 실수부와 허수부로 된 복소수
	   complex128  float64 크기의 실수부와 허수부로 된 복소수
	   byte      uint8과 크기가 동일, 바이트 단위로 저장 시 사용
	   rune      int32와 크기가 동일, 유니코드 문자 코드(Code point)를 저장할 때 사용
	*/

	// 정수
	var _ int = 32
	var _ int = -15
	var _ int = 0723
	var _ int = 0x32fa2c75

	// 실수
	// 고정 소수점
	var _ float32 = 0.2
	var _ float32 = .35
	var _ float32 = 132.73287

	// 부동 소수점 방식
	var _ float32 = 1e7
	var _ float32 = .12345E+2
	var _ float32 = 5.32521e-10

	// go 언어 머신 엡실론
	//const epsilon := 1e-14

}
