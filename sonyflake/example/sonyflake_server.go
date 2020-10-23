package main

import (
	"encoding/json"
	"net/http"

	"github.com/yacen/gotools/sonyflake"
	"github.com/yacen/gotools/sonyflake/awsutil"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	var err error
	st.MachineID = awsutil.AmazonEC2MachineID
	sf, err = sonyflake.NewSonyflake(st)
	if err != nil {
		panic("sonyflake not created")
	}
	if sf == nil {
		panic("sonyflake not created")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := sf.NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(sonyflake.Decompose(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
