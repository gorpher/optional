package main

import (
	"log"
	"net/http"

	"github.com/gorpher/optional/v2"
)

func main() {
	if err := http.ListenAndServe(":9870", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证单个参数
		var name string
		err := optional.HttpRequestFormVal(r).
			Validates().
			Aligns()
		// 多参数值验证
		var age int
		err = optional.HttpRequestFormVal(r).
			Validates(
				optional.Validate("name", optional.MustString(), optional.MustHasLetter()),
				optional.Validate("age")).
			Aligns(
				optional.Align("name", name),
				optional.Align("age", age))
		// 结构体赋值
		var req struct {
			name string `json:"name,omitempty" validate:"name,MustString,Min(10),Max(20)"`
			age  int    `json:"age,omitempty" `
		}

		err = optional.HttpRequestQueryVal(r).
			Validates(
				optional.Validate("name", optional.MustString(), optional.MustHasLetter()),
				optional.Validate("age")).
			Align(optional.Align("name", &req))

		optional.StringVal(r.PostFormValue("name")).
			Converter().String()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

	})); err != nil {
		log.Fatal(err)
	}
}
