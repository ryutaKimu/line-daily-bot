package main

import "testing"

func TestFetchNarrationAPI(t *testing.T) {
	n, err := fetchNarrationAPI()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", n)
}

func TestBuildMessage(t *testing.T) {
	got := buildMessage(Narration{
		Episode:    3,
		Title:      "無謀なる旅路",
		Narrations: []string{"1行目", "2行目", "3行目"},
	})
	want := "Episode 第3話\n『無謀なる旅路』\n\n1行目\n2行目\n3行目"
	if got != want {
		t.Errorf("buildMessage() =\n%q\nwant\n%q", got, want)
	}
}

func TestFormatNarration(t *testing.T) {
	n, err := formatNarattions()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", n)
}
