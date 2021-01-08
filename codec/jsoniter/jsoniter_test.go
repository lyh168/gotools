package jsoniter

import (
	"fmt"
	"testing"
	"time"
)

type t struct {
	BeginTime time.Time
}

func TestUnmarshal(testt *testing.T)  {
	ttt := t{BeginTime: time.Now()}
	tttBytes, err := Marshal(ttt)
	if err != nil {
	  fmt.Println(err)
	  return
	}
  
	s := string(tttBytes)

	testt.Log(s)
	var tttt t
	if err := Unmarshal(tttBytes, &tttt); err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(tttt)
	
	info := `{"begin_time": "1970-01-01T08:00:01+08:00"}`
  
	var tt t
	if err := Unmarshal([]byte(info), &tt); err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(tt)
}