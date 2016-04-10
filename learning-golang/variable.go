package main

import "fmt"

func main() {
  var i int
  var s string

  var age int = 10
  var name string = "Junyoung"

  // 자료형 생략
  var age1 = 10
  var name1 = "Junyoung Lee"
  // var address // 컴파일 에러

  // 짧은 선언
  age2 := 10
  name2 := "Junyoung Lee"

  // 변수 여러 개 선언하기
  var x, y int = 30, 50
  var age3, name3 = 10, "Junyoung Lee"

  a, b, c, d := 1, 3.4, "Hello, world!", false

  // 병렬 할당(Parallel assignement)
  var e, f int
  var age5 int
  e, f, age5 = 10, 20, 5

  var (
    g, h  int = 30, 50
    age6, name6 = 10, "Maria"
    i, j := 3.14, "Hello"
  )

  // 사용하지 않는 변수 및 패키지 처리
  /*
  Go 언어는 선언만 하고 사용하지 않는 변수가 있으면 컴파일 에러 발생
  import 하고 사용하지 않는 패키지가 있을 때도 컴파일 에러 발생
  */

  _ = a // 사용하지 않는 변수로 인한 컴파일 에러 방지

  // import의 경우
  // _import _ "time" // 패키지 명 앞에 _ 를 붙여 컴파일 에러 방지

}
