package racer

import "testing"

func TestRace(t *testing.T) {
	slowUrl := "https://www.baidu.com"
	fastUrl := "http://www.sina.com"
	want := fastUrl
	got := websiteRacer(fastUrl, slowUrl)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}
