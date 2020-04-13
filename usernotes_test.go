package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestReadNotes(t *testing.T) {
	want := UsernoteBlob{
		"SpyTec13": UsernoteList{
			Notes: []Usernote{
				{
					Time:      1586733280,
					Note:      "Note 2",
					Moderator: 0,
					Warning:   0,
					Link:      "l,3aevs8",
				},
				{
					Time:      1586733273,
					Note:      "Note 1",
					Moderator: 0,
					Warning:   1,
					Link:      "l,3aevs8",
				},
			},
		},
	}
	blob, err := ioutil.ReadFile("blob.json")

	if err != nil {
		log.Fatal(err)
	}

	got := NewUsernoteManager(blob).Blob
	eq := reflect.DeepEqual(got, want)
	// t.Errorf("Usernotes %v: got %+v; want %+v", eq, got, want)
	if !eq {
		t.Errorf("Usernotes: got\n %+v; want\n %+v", got, want)
	}
}
