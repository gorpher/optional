package optional

import (
	"encoding/base64"
	"html"
	"math/big"
	"net/url"
	"strings"
)

type stringProcessor interface {
	Trim(s string) stringProcessor
	TrimSpace() stringProcessor
	TrimLeft(s string) stringProcessor
	TrimRight(s string) stringProcessor
	TrimPrefix(s string) stringProcessor
	TrimSuffix(s string) stringProcessor
	RemoveSpace() stringProcessor
	ToUpper() stringProcessor
}

type numberProcessor interface {
	Add(f0, f1 float64) float64
}

type linkProcessor struct {
	stringProcessor
	numberProcessor
	err error
	v   Value
}

func (o *linkProcessor) Value() Value {
	return o.v
}

func (o *linkProcessor) stringMap(fn func(a, v string) string, s string) stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = fn(o.v.v.(string), s)
	return o
}
func (o *linkProcessor) Trim(s string) stringProcessor {
	if o.err != nil {
		return o
	}
	return o.stringMap(strings.Trim, s)
}
func (o *linkProcessor) TrimSpace() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = strings.TrimSpace(o.v.v.(string))
	return o
}
func (o *linkProcessor) TrimLeft(s string) stringProcessor {
	if o.err != nil {
		return o
	}
	return o.stringMap(strings.TrimLeft, s)
}
func (o *linkProcessor) TrimRight(s string) stringProcessor {
	if o.err != nil {
		return o
	}
	return o.stringMap(strings.TrimRight, s)
}
func (o *linkProcessor) TrimPrefix(s string) stringProcessor {
	if o.err != nil {
		return o
	}
	return o.stringMap(strings.TrimPrefix, s)
}
func (o *linkProcessor) TrimSuffix(s string) stringProcessor {
	if o.err != nil {
		return o
	}
	return o.stringMap(strings.TrimSuffix, s)
}
func (o *linkProcessor) RemoveSpace() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = strings.ReplaceAll(o.v.v.(string), " ", "")
	return o
}
func (o *linkProcessor) ToUpper() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = strings.ToUpper(o.v.v.(string))
	return o
}
func (o *linkProcessor) ToLower() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = strings.ToLower(o.v.v.(string))
	return o
}
func (o *linkProcessor) ToPascalCase() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = ToPascalCase(o.v.v.(string))
	return o
}
func (o *linkProcessor) ToCamelCase() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = ToCamelCase(o.v.v.(string))
	return o
}
func (o *linkProcessor) ToSnakeCase() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = ToSnakeCase(o.v.v.(string))
	return o
}

func (o *linkProcessor) Base64StdEncode() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = base64.StdEncoding.EncodeToString(strToBytes(o.v.v.(string)))
	return o
}
func (o *linkProcessor) Base64StdDecode() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v, o.err = base64.StdEncoding.DecodeString(o.v.v.(string))
	return o
}
func (o *linkProcessor) Base64RawStdEncode() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = base64.RawStdEncoding.EncodeToString(strToBytes(o.v.v.(string)))
	return o
}
func (o *linkProcessor) Base64RawStdDecode() stringProcessor {
	if o.err != nil {
		return o
	}
	b, err := base64.RawStdEncoding.DecodeString(o.v.v.(string))
	o.err = err
	o.v.v = bytesToStr(b)
	return o
}
func (o *linkProcessor) Base64URLEncode() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = base64.URLEncoding.EncodeToString(strToBytes(o.v.v.(string)))
	return o
}
func (o *linkProcessor) Base64URLDecode() stringProcessor {
	if o.err != nil {
		return o
	}
	bytes, err := base64.URLEncoding.DecodeString(o.v.v.(string))
	o.v.v = bytesToStr(bytes)
	o.err = err
	return o
}
func (o *linkProcessor) Base64RawURLEncode() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = base64.RawURLEncoding.EncodeToString(strToBytes(o.v.v.(string)))
	return o
}
func (o *linkProcessor) Base64RawURLDecode() stringProcessor {
	if o.err != nil {
		return o
	}
	bytes, err := base64.RawURLEncoding.DecodeString(o.v.v.(string))
	o.v.v = bytesToStr(bytes)
	o.err = err
	return o
}

func (o *linkProcessor) HTMLUnescape() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = html.UnescapeString(o.v.v.(string))
	return o
}
func (o *linkProcessor) HTMLEscape() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = html.EscapeString(o.v.v.(string))
	return o
}
func (o *linkProcessor) URLPathUnescape() stringProcessor {
	if o.err != nil {
		return o
	}
	value, err := url.PathUnescape(o.v.v.(string))
	o.v.v = value
	o.err = err
	return o
}
func (o *linkProcessor) URLPathEscape() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = url.PathEscape(o.v.v.(string))
	return o
}
func (o *linkProcessor) URLQueryEscape() stringProcessor {
	if o.err != nil {
		return o
	}
	o.v.v = url.QueryEscape(o.v.v.(string))
	return o
}

func (o *linkProcessor) JoinStr(values ...string) stringProcessor {
	if o.err != nil {
		return o
	}
	var value strings.Builder
	value.WriteString(o.v.v.(string))
	for k := range values {
		value.WriteString(values[k])
	}
	o.v.v = value.String()
	return o
}
func (o *linkProcessor) JoinBytes(values ...[]byte) stringProcessor {
	if o.err != nil {
		return o
	}
	var value strings.Builder
	value.WriteString(o.v.v.(string))
	for k := range values {
		value.Write(values[k])
	}
	o.v.v = value.String()
	return o
}
func (o *linkProcessor) JoinByte(values ...byte) stringProcessor {
	if o.err != nil {
		return o
	}
	var value strings.Builder
	value.WriteString(o.v.v.(string))
	for k := range values {
		value.WriteByte(values[k])
	}
	o.v.v = value.String()
	return o
}
func (o *linkProcessor) JoinRune(values ...rune) stringProcessor {
	if o.err != nil {
		return o
	}
	var value strings.Builder
	value.WriteString(o.v.v.(string))
	for k := range values {
		value.WriteRune(values[k])
	}
	o.v.v = value.String()
	return o
}

func (o *linkProcessor) Add(f0, f1 float64) float64 {
	f, _ := big.NewFloat(f0).Add(big.NewFloat(f1), nil).Float64()
	return f
}

func (o *linkProcessor) GetError() error {
	return o.err
}
