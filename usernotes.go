package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"log"
)

type Usernote struct {
	Time      int    `json:"t"`
	Note      string `json:"n"`
	Moderator int    `json:"m"`
	Warning   int    `json:"w"`
	Link      string `json:"l"`
}

type UsernoteList struct {
	Notes []Usernote `json:"ns"`
}

type User string
type Warning string
type UsernoteBlob map[User]UsernoteList

func (ub *UsernoteBlob) UnmarshalJSON(b []byte) error {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	_, err := base64.StdEncoding.Decode(decoded, b[1:len(b)-1])
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
	var blob map[User]UsernoteList
	if err := json.Unmarshal(buf.Bytes(), &blob); err != nil {
		return err
	}
	*ub = blob
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
		return nil
	}
	return &usernoteManager
}
