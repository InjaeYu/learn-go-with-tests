package concurrency

import (
	"reflect"
	"testing"
)

func mockWebSiteChecker(url string) bool {
	if url == "watt://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.amoeba.com",
		"watt://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":       true,
		"http://blog.amoeba.com":  true,
		"watt://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebSiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}
