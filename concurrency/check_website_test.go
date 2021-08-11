package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteCheck(url string) bool {
	if url == "http://www.google.com" {
		return true
	} else {
		return false
	}
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://blog.gypsydave5.com",
		"http://furhurterwe.geds",
	}

	actualResults := CheckWebsite(mockWebsiteCheck, websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("want '%d' got '%d'", want, got)
	}
	expectResult := map[string]bool{
		"http://www.google.com":      true,
		"http://blog.gypsydave5.com": false,
		"http://furhurterwe.geds":    false,
	}
	if !reflect.DeepEqual(actualResults, expectResult) {
		t.Fatalf("want '%v' got '%v'", expectResult, actualResults)
	}

}
