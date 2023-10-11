package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("사람들에게 인사하기", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("빈 문자열에 대한 기본값 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("스페인어 사용", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("프랑스어 사용", func(t *testing.T) {
		got := Hello("Louis", french)
		want := "Bonjour, Louis"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// 테스트 실패시, 실패한 라인에 대한 정보를 해당 라인이 아닌 해당 함수를 호출한 라인으로 표시하도록 설정하는 함수
	// 즉, t.Helper() 함수를 명시하지 않은 상태에서 테스트 실패시 26번째 라인에서 발생했다고 출력
	// t.Helper() 함수를 명시한 상태에서 테스트 실패시 9번째 라인 또는 15번째 라인에서 발생했다고 출력
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
