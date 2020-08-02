package optional

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func HttpRequestQueryVal(req *http.Request) Value {
	if err := req.ParseForm(); err != nil {
		return Value{
			err: err,
		}
	}
	m := map[string]Value{}
	queries := req.URL.Query()
	for k, v := range queries {
		m[k] = StringVal(v[0])
	}
	return Value{}
}

func HttpRequestFormVal(req *http.Request) Value {
	if err := req.ParseForm(); err != nil {
		return Value{
			err: err,
		}
	}
	// form中必须有参数和值
	if len(req.Form) == 0 {
		return Value{
			err: errors.New("没有期望的请求参数和值"),
		}
	}
	m := map[string]Value{}
	// 只获取第一个参数的值
	for k, v := range req.Form {
		m[k] = StringVal(v[0])
	}
	return MapStringVal(m)
}

func HttpRequestBodyVal(req *http.Request) Value {
	var val interface{}
	json.NewDecoder(req.Body).Decode(&val)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Value{err: err}
	}
	err = json.Unmarshal(body, &val)
	if err != nil {
		return Value{err: err}
	}
	//todo 添加对象解析
	return Value{v: val, err: errors.New("TODO")}
}
