# 테스트와 함께 Go 언어 배우기
TDD 기반 Go 언어 학습을 위한 저장소

## 개요
[해당 사이트](https://quii.gitbook.io/learn-go-with-tests/)에서 제공하는 내용에 따라 TDD 기반 Go 언어 학습을 진행하며, 진행한 내용을 기록 및 공유하기 위한 저장소

## Go 테스트 원칙
1. 테스트 **파일**은 `xxx_test.go`와 같은 형태로 만들어야 한다.
    - 예 : `adder_test.go`
2. 테스트 **함수**는 `Test`로 시작되어야 한다.
    - 예 : `func TestAdder(t *testing.T) {}`
3. 테스트 **함수**는 `t *testing.T`라는 파라미터 **1개만 사용**해야 한다.
4. `*testing.T`를 사용하기 위해선 `import testing`을 해주어야 한다.

## TDD 사이클
아래의 사이클을 반복하며 개발 진행
1. 테스트 작성
2. 컴파일이 되도록 수정
3. 테스트 실행, 실패 유무 확인 및 에러 메시지가 의미 있는지 확인
4. 테스트가 통과될 수 있도록 충분한 코드 작성
5. 리팩터링 진행

## Go 문서 작성 방법
1. 함수에 대한 설명을 함수 상단에 주석으로 작성하면 Go 문서에 해당 함수 설명으로 등록됨
    - go doc에 표현되는 내용은 [링크](https://pkg.go.dev/github.com/InjaeYu/learn-go-with-tests@v0.0.0-20231011093524-49f4951285ef/integers/v2) 참조
```go
// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}
```
2. **테스트 코드**에서 `Example`로 시작하는 함수를 구현  후 해당 함수의 사용 방법 및 결과를 작성하면 Go 문서에 해당 함수 예제로 등록됨
    - 결과의 경우 다음과 같이 주석으로 작성해야 한다 : `// Output: X`
    - `Example` 뒤에는 실제 함수명 그대로 사용해야 인식 (`Add()`면 `ExampleAdd()`로 정의해야 인식)
```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```
3. [Go 공식 사이트](https://go.dev/)의 `Packages`에 구현한 모듈의 문서를 정상적으로 표출하기 위해선 해당 프로젝트에 License가 추가되어 있어야 함