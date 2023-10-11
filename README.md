# 테스트와 함께 Go 언어 배우기
TDD 기반 Go 언어 학습을 위한 저장소

## 개요
[해당 사이트](https://quii.gitbook.io/learn-go-with-tests/)에서 제공하는 내용에 따라 TDD 기반 Go 언어 학습을 진행하며, 진행한 내용을 기록 및 공유하기 위한 저장소

## Go 테스트 원칙
1. 테스트 **파일**은 `xxx_test.go`와 같은 형태로 만들어야 한다.
2. 테스트 **함수**는 `Test`로 시작되어야 한다.
3. 테스트 **함수**는 `t *testing.T`라는 파라미터 **1개만 사용**해야 한다.
4. `*testing.T`를 사용하기 위해선 `import testing`을 해주어야 한다.