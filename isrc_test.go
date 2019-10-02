package isrc

import (
	//"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewISRC(t *testing.T) {
	testStrings := []string{
		"UKABC9954321",
		"UST0T0012345",
		"NONNN5033333",
	}
	for _, testString := range testStrings {
		i, err := NewISRC(testString)
		if err != nil {
			t.Fatal(err)
		}
		response := "expected:%s got:%s"
		if i.String() != testString {
			t.Fatalf(response, testString, i)
		}
		t.Logf(response, testString, i)
	}
}

func TestBadISRC(t *testing.T) {
	testStrings := []string{
		"c",
		"01BBB8898765",
		"DD000XX12345",
	}
	for _, testString := range testStrings {
		_, err := NewISRC(testString)
		if err == nil {
			t.Fatalf("Bad ISRC %s should have errored", testString)
		}
		t.Logf("Bad ISRC %s correctly caused an error", testString)
	}
}
