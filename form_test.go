package optional_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gorpher/optional/v2"
)

func handleRequest(resp http.ResponseWriter, req *http.Request) {
	values := optional.HttpRequestFormVal(req)
	if err := values.GetErrorResponseWriter(resp); err != nil {
		log.Print(err)
		return
	}
	log.Print(values.GetMapValue("author"))
	log.Print(values.GetMapValue("age"))
	log.Print(values.GetMapValue("title"))
	log.Print(values.GetMapValue("content"))
	log.Print(values.GetMapValue("email"))
	vvalues := values.Validates(
		optional.Validate("author", optional.MustIsLetter()),
		optional.Validate("age", optional.MustIsNumberValue()),
	).Value()

	if err := vvalues.GetErrorResponseWriter(resp); err != nil {
		log.Print(err)
		return
	}
	log.Print(vvalues.GetMapValue("author"))
	log.Print(vvalues.GetMapValue("age"))
	log.Print(vvalues.GetMapValue("title"))
	log.Print(vvalues.GetMapValue("content"))
	log.Print(vvalues.GetMapValue("email"))
	pvalue := vvalues.Processors(
		optional.Process("author", optional.ToUpper()),
		optional.Process("age", optional.ToInt()),
	).Value()
	log.Print(pvalue.GetMapValue("author"))
	log.Print(pvalue.GetMapValue("age"))

	var res struct {
		Author string
		Age    int
		Title  string
		Email  string
	}

	if err := values.Aligns(
		optional.Align("author", &res.Author),
		optional.Align("age", &(res.Age)),
		optional.Align("title", &(res.Title)),
		optional.Align("email", &(res.Email)),
	); err != nil {
		log.Print(err)
		return
	}
	log.Printf("%#v", res)

	sv := values.GetMapValue("author").Processor("author", optional.Base64StdEncode()).Value()

	log.Print(sv)
}

var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {

	http.DefaultServeMux.HandleFunc("/", handleRequest)

	w = httptest.NewRecorder()

	os.Exit(m.Run())
}
func TestHttpRequestFormVal(t *testing.T) {
	value := url.Values{}
	value.Add("author", "gorpher")
	value.Add("age", "24")
	value.Add("title", "The Go Standard Library")
	value.Add("content", "It contains many packages.")
	value.Add("email", "gorpher@gmail.com")

	reader := strings.NewReader(value.Encode())

	r, _ := http.NewRequest(http.MethodPost, "/", reader)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	http.DefaultServeMux.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %value", resp.StatusCode)
	}
}
