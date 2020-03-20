package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type Usernote struct {
	Timestamp int    `json:t`
	Note      string `json:n`
	Moderator int    `json:m`
	Warning   int    `json:w`
	Link      string `json:l`
}

type UsernoteList struct {
	Notes []Usernote `json:ns`
}

type User string
type Warning string
type UsernoteBlob map[User]UsernoteList

func (ub *UsernoteBlob) UnmarshalJSON(b []byte) error {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	_, err := base64.StdEncoding.Decode(decoded, b[1:len(b)-1])
	fmt.Println(string(b[:50]))
	if err != nil {
		return nil
	}
	buffer := bytes.NewBuffer(decoded)
	buf := new(bytes.Buffer)
	r, err := zlib.NewReader(buffer)
	if err != nil {
		log.Fatal(err)
	}
	buf.ReadFrom(r)
	// buf, err := ioutil.ReadAll(r)
	if err := json.Unmarshal(buf.Bytes(), &ub); err != nil {
		return err
	}
	return nil
}

type UsernoteConstants struct {
	Users    []User
	Warnings []Warning
}

type UsernoteManager struct {
	Ver       int `json:"ver"`
	Constants UsernoteConstants
	Blob      UsernoteBlob
}

func NewUsernoteManager(rawUsernotes []byte) *UsernoteManager {
	var usernoteManager UsernoteManager
	err := json.Unmarshal(rawUsernotes, &usernoteManager)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &usernoteManager
}
