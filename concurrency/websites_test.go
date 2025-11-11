package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://notmywebsite.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://mywebsite.com",
		"https://notmywebsite.com",
		"https://anotherwebsite.com",
	}

	want := map[string]bool{
		"https://mywebsite.com":      true,
		"https://notmywebsite.com":   false,
		"https://anotherwebsite.com": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "https://example.com"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
