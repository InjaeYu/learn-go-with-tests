package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// map은 runtime.hmap 구조체의 포인터라서 메서드에서 포인터로 넘겨주지 않아도 수정이 가능
// 즉, map의 값은 포인터이기 때문에 nil이 들어갈 수 있으므로 map을 초기화할 때 절대 빈 값으로 초기화를 하면 안된다.
// 빈 값으로 초기화를 하는 경우(var m map[string]string) 읽을 땐 빈 맵처럼 동작, 쓰기를 요청할 때 panic 발생 (nil pointer 접근 에러)
// map을 초기화할 땐 `map[string]string{}`과 같이 비어있는 map으로 초기화를 하던가 `make()`를 사용하여 초기화를 해야한다.
func (d Dictionary) Add(key, value string) {
	d[key] = value
}
