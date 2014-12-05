package urlquery

import (
	"fmt"
	"net/url"
	"testing"
)

func TestMarshal(t *testing.T) {
	s1 := struct {
		V1 string `url:"v1"`
		V2 bool   `url:"hello"`
	}{"hi", true}
	fmt.Println(Marshal(s1).Encode())
	fmt.Println(Marshal(&s1).Encode())
	fmt.Println(Marshal(1).Encode())
}

func TestUnmarshal(t *testing.T) {
	s1 := struct {
		V1 string `url:"v1"`
		V2 bool   `url:"hello"`
	}{}
	val := url.Values{}
	val.Set("v1", "hi")
	val.Set("hello", "1")
	err := Unmarshal(val, &s1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s1)
	}
}
