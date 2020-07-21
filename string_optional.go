package optional

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"html"
	"log"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stringOptional struct {
	err error
	v   string
}

func (o *stringOptional) is(fn func(r rune) bool) bool {
	for _, v := range o.v {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (o *stringOptional) has(fn func(r rune) bool) bool {
	for _, v := range o.v {
		if fn(v) {
			return true
		}
	}
	return false
}
func (o *stringOptional) converter(fn func(a, v string) string, s string) *stringOptional {
	o.v = fn(o.v, s)
	return o
}

// result
func (o *stringOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *stringOptional) GetError() error {
	return o.err
}

// translate
func (o *stringOptional) Trim(s string) *stringOptional {
	return o.converter(strings.Trim, s)
}
func (o *stringOptional) TrimSpace() *stringOptional {
	o.v = strings.TrimSpace(o.v)
	return o
}
func (o *stringOptional) TrimLeft(s string) *stringOptional {
	return o.converter(strings.TrimLeft, s)
}
func (o *stringOptional) TrimRight(s string) *stringOptional {
	return o.converter(strings.TrimRight, s)
}
func (o *stringOptional) TrimPrefix(s string) *stringOptional {
	return o.converter(strings.TrimPrefix, s)
}
func (o *stringOptional) TrimSuffix(s string) *stringOptional {
	return o.converter(strings.TrimSuffix, s)
}
func (o *stringOptional) RemoveSpace() *stringOptional {
	o.v = strings.ReplaceAll(o.v, " ", "")
	return o
}
func (o *stringOptional) ToUpper() *stringOptional {
	o.v = strings.ToUpper(o.v)
	return o
}
func (o *stringOptional) ToLower() *stringOptional {
	o.v = strings.ToLower(o.v)
	return o
}

func (o *stringOptional) ToPascalCase() *stringOptional {
	o.v = ToPascalCase(o.v)
	return o
}
func (o *stringOptional) ToCamelCase() *stringOptional {
	o.v = ToCamelCase(o.v)
	return o
}
func (o *stringOptional) ToSnakeCase() *stringOptional {
	o.v = ToSnakeCase(o.v)
	return o
}

func (o *stringOptional) Base64StdEncode() *stringOptional {
	o.v = base64.StdEncoding.EncodeToString(strToBytes(o.v))
	return o
}
func (o *stringOptional) Base64StdDecode() *bytesOptional {
	b := &bytesOptional{err: o.err}
	b.v, b.err = base64.StdEncoding.DecodeString(o.v)
	return b
}
func (o *stringOptional) Base64RawStdEncode() *stringOptional {
	o.v = base64.RawStdEncoding.EncodeToString(strToBytes(o.v))
	return o
}
func (o *stringOptional) Base64RawStdDecode() *stringOptional {
	b, err := base64.RawStdEncoding.DecodeString(o.v)
	o.err = err
	o.v = bytesToStr(b)
	return o
}
func (o *stringOptional) Base64URLEncode() *stringOptional {
	o.v = base64.URLEncoding.EncodeToString(strToBytes(o.v))
	return o
}
func (o *stringOptional) Base64URLDecode() *stringOptional {
	bytes, err := base64.URLEncoding.DecodeString(o.v)
	o.v = bytesToStr(bytes)
	o.err = err
	return o
}

func (o *stringOptional) Base64RawURLEncode() *stringOptional {
	o.v = base64.RawURLEncoding.EncodeToString(strToBytes(o.v))
	return o
}
func (o *stringOptional) Base64RawURLDecode() *stringOptional {
	bytes, err := base64.RawURLEncoding.DecodeString(o.v)
	o.v = bytesToStr(bytes)
	o.err = err
	return o
}

func (o *stringOptional) HTMLUnescape() *stringOptional {
	o.v = html.UnescapeString(o.v)
	return o
}
func (o *stringOptional) HTMLEscape() *stringOptional {
	o.v = html.EscapeString(o.v)
	return o
}
func (o *stringOptional) URLPathUnescape() *stringOptional {
	value, err := url.PathUnescape(o.v)
	o.v = value
	o.err = err
	return o
}
func (o *stringOptional) URLPathEscape() *stringOptional {
	o.v = url.PathEscape(o.v)
	return o
}
func (o *stringOptional) URLQueryEscape() *stringOptional {
	o.v = url.QueryEscape(o.v)
	return o
}

func (o *stringOptional) JoinStr(values ...string) *stringOptional {
	var value strings.Builder
	value.WriteString(o.v)
	for k := range values {
		value.WriteString(values[k])
	}
	o.v = value.String()
	return o
}
func (o *stringOptional) JoinBytes(values ...[]byte) *stringOptional {
	var value strings.Builder
	value.WriteString(o.v)
	for k := range values {
		value.Write(values[k])
	}
	o.v = value.String()
	return o
}
func (o *stringOptional) JoinByte(values ...byte) *stringOptional {
	var value strings.Builder
	value.WriteString(o.v)
	for k := range values {
		value.WriteByte(values[k])
	}
	o.v = value.String()
	return o
}
func (o *stringOptional) JoinRune(values ...rune) *stringOptional {
	var value strings.Builder
	value.WriteString(o.v)
	for k := range values {
		value.WriteRune(values[k])
	}
	o.v = value.String()
	return o
}

//
//
// validate
func (o *stringOptional) IsNil() *boolOptional {
	return &boolOptional{err: o.err, v: o.v == ""}
}
func (o *stringOptional) IsBool() *boolOptional {
	b := &boolOptional{
		err: o.err,
	}
	b.v, b.err = strconv.ParseBool(o.v)
	return b
}
func (o *stringOptional) Equals(s string) *boolOptional {
	if o.v == s {
		return &boolOptional{err: o.err, v: true}
	}
	return &boolOptional{err: o.err, v: false}
}
func (o *stringOptional) IsLower() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(unicode.IsLower)}
}
func (o *stringOptional) IsUpper() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(unicode.IsUpper)}

}
func (o *stringOptional) IsLetter() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(unicode.IsLetter)}
}
func (o *stringOptional) IsDigit() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(unicode.IsDigit)}
}
func (o *stringOptional) IsLowerOrDigit() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(func(r rune) bool {
		if unicode.IsLower(r) || unicode.IsDigit(r) {
			return true
		}
		return false
	})}
}
func (o *stringOptional) IsUpperOrDigit() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(func(r rune) bool {
		if unicode.IsUpper(r) || unicode.IsDigit(r) {

			return true
		}
		return false
	})}
}
func (o *stringOptional) IsLetterOrDigit() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
		return false
	})}
}
func (o *stringOptional) IsChinese() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(func(r rune) bool {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
		return false
	})}
}
func (o *stringOptional) IsChinaTel() *boolOptional {
	//todo 添加错误处理
	b := &boolOptional{err: o.err, v: false}
	telSlice := strings.Split(o.v, "-")
	if len(telSlice) != 2 {
		b.err = errors.New("")
		return b
	}
	regionCode, err := strconv.Atoi(telSlice[0])
	if err != nil {
		b.err = err
		return b
	}
	if regionCode < 10 || regionCode > 999 {
		b.err = errors.New("")
		return b
	}
	code, err := strconv.Atoi(telSlice[1])
	if err != nil {
		b.err = err
		return b
	}
	if code < 1000000 || code > 99999999 {
		b.err = errors.New("")
		return b
	}
	return b
}
func (o *stringOptional) IsMail() *boolOptional {
	//todo 错误判断
	b := &boolOptional{err: o.err, v: false}
	emailSlice := strings.Split(o.v, "@")
	if len(emailSlice) != 2 {
		b.err = errors.New("")
		return b
	}
	if emailSlice[0] == "" || emailSlice[1] == "" {
		b.err = errors.New("")
		return b
	}

	for k, v := range emailSlice[0] {
		if k == 0 && !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			b.err = errors.New("")
			return b
		} else if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '@' && v != '.' && v != '_' && v != '-' {
			b.err = errors.New("")
			return b
		}
	}

	domainSlice := strings.Split(emailSlice[1], ".")
	if len(domainSlice) < 2 {
		b.err = errors.New("")
		return b
	}
	domainSliceLen := len(domainSlice)
	for i := 0; i < domainSliceLen; i++ {
		for k, v := range domainSlice[i] {
			// nolint
			if i != domainSliceLen-1 && k == 0 && !unicode.IsLetter(v) && !unicode.IsDigit(v) {
				b.err = errors.New("")
				return b
			} else if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '.' && v != '_' && v != '-' {
				b.err = errors.New("")
				return b
			} else if i == domainSliceLen-1 && !unicode.IsLetter(v) {
				b.err = errors.New("")
				return b
			}
		}
	}
	return b
}
func (o *stringOptional) IsIP() *boolOptional {
	b := &boolOptional{err: o.err, v: false}
	if v := net.ParseIP(o.v); v != nil {
		//todo ip的具体判断
		return b
	}
	return b
}
func (o *stringOptional) IsJSON() *boolOptional {
	b := &boolOptional{err: o.err, v: true}
	var js json.RawMessage
	if json.Unmarshal([]byte(o.v), &js) != nil {
		return b
	}
	return b
}
func (o *stringOptional) IsChinaIDNumber() *boolOptional {
	//todo 完善中国电话号码验证
	b := &boolOptional{err: o.err, v: true}
	var idV int
	if o.v[17:] == "X" {
		idV = 88
	} else {
		var err error
		if idV, err = strconv.Atoi(o.v[17:]); err != nil {
			b.err = err
			return b
		}
	}

	var verify int
	id := o.v[:17]
	arr := make([]int, 17)
	for i := 0; i < 17; i++ {
		var err error
		arr[i], err = strconv.Atoi(string(id[i]))
		if err != nil {
			b.err = err
			return b
		}
	}
	wi := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += arr[i] * wi[i]
	}
	verify = res % 11

	var temp int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp != idV {
		b.err = errors.New("自己定义errror")
		return b
	}
	return b
}
func (o *stringOptional) IsChinaMobile() *boolOptional {
	b := &boolOptional{err: o.err, v: o.is(unicode.IsDigit)}
	if len(o.v) != 11 {
		o.err = errors.New("自定义错误")
		return b
	}
	var (
		prefix      uint8
		prefixValid bool
	)
	if prefix64, err := strconv.ParseUint(o.v[0:3], 10, 8); err != nil {
		o.err = errors.New("自定义错误")
		return b
	} else {
		prefix = uint8(prefix64)
	}
	log.Print(prefix)
	//todo
	//for k := range chinaMobilePrefix {
	//	if chinaMobilePrefix[k] == prefix {
	//		prefixValid = true
	//		break
	//	}
	//}
	if !prefixValid {
		o.err = errors.New("自定义错误")
		return b
	}
	if _, err := strconv.ParseUint(o.v[3:], 10, 32); err != nil {
		o.err = errors.New("自定义错误")
		return b
	}
	return b
}
func (o *stringOptional) IsSQLObject() *boolOptional {
	return &boolOptional{err: o.err, v: o.is(func(r rune) bool {
		// 是否是sql对象
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_' {
			return false
		}
		return true
	})}
}
func (o *stringOptional) IsURL() *boolOptional {
	b := &boolOptional{err: o.err, v: true}
	if _, err := url.ParseRequestURI(o.v); err != nil {
		b.err = err
		b.v = false
	}
	return b
}
func (o *stringOptional) IsUUID() *boolOptional {
	b := &boolOptional{err: o.err, v: false}
	str := o.v
	//todo 验证完善uuid
	var uuid [16]byte
	switch len(str) {
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36:

	// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 36 + 9:
		if strings.EqualFold(strings.ToLower(str[:9]), "urn:uuid:") {
			b.err = errors.New("自定义错误")
			b.v = false
			return b
		}
		str = str[9:]

	// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
	case 36 + 2:
		str = str[1:]

	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	case 32:
		var ok bool
		for i := range uuid {
			uuid[i], ok = xtob(str[i*2], str[i*2+1])
			if !ok {
				b.err = errors.New("自定义错误")
				b.v = false
				return b
			}
		}
		return b
	default:
		b.err = errors.New("自定义错误")
		b.v = false
		return b
	}
	// s is now at least 36 bytes long
	// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		b.err = errors.New("自定义错误")
		b.v = false
		return b
	}
	for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		v, ok := xtob(str[x], str[x+1])
		if !ok {
			b.err = errors.New("自定义错误")
			b.v = false
			return b
		}
		uuid[i] = v
	}
	return b
}
func (o *stringOptional) In(s []string) *boolOptional {
	b := &boolOptional{err: o.err, v: false}
	for i := range s {
		if s[i] == o.v {
			b.v = true
			return b
		}
	}
	return b
}

func (o *stringOptional) HasLetter() *boolOptional {
	return &boolOptional{err: o.err, v: o.has(unicode.IsLetter)}
}
func (o *stringOptional) HasLower() *boolOptional {
	return &boolOptional{err: o.err, v: o.has(unicode.IsLower)}
}
func (o *stringOptional) HasUpper() *boolOptional {
	return &boolOptional{err: o.err, v: o.has(unicode.IsUpper)}
}
func (o *stringOptional) HasDigit() *boolOptional {
	return &boolOptional{err: o.err, v: o.has(unicode.IsDigit)}
}
func (o *stringOptional) HasSymbol() *boolOptional {
	return &boolOptional{err: o.err, v: o.has(unicode.IsSymbol)}
}

func (o *stringOptional) Contains(s string) *boolOptional {
	return &boolOptional{err: o.err, v: strings.Contains(o.v, s)}
}
func (o *stringOptional) HasString(s string) *boolOptional {
	return &boolOptional{err: o.err, v: strings.Contains(o.v, s)}
}
func (o *stringOptional) HasPrefix(s string) *boolOptional {
	return &boolOptional{err: o.err, v: strings.HasPrefix(o.v, s)}
}
func (o *stringOptional) HasSuffix(s string) *boolOptional {
	return &boolOptional{err: o.err, v: strings.HasSuffix(o.v, s)}
}

// set value
func (o *stringOptional) CanAlign(t interface{}) *boolOptional {
	b := &boolOptional{err: o.err, v: false}
	if t == nil {
		b.err = errors.New("target cannot be nil")
		return b
	}
	targetValueOf := reflect.ValueOf(t)
	// 检查对象是否是指针
	if targetValueOf.Kind() != reflect.Ptr {
		b.err = errors.New("target must be a pointer")
		return b
	}
	// 检查对象是否能赋值
	if !targetValueOf.Elem().CanSet() {
		b.err = errors.New("cannot set the value of the target")
		return b
	}
	b.v = true
	return b
}
func (o *stringOptional) SetValue(i interface{}) error {
	b := o.CanAlign(i)
	if b.HashError() {
		return b.err
	}
	v := reflect.ValueOf(i)
	t := v.Elem().Type().Kind()
	switch t {
	case reflect.String:
		v.Elem().SetString(o.v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

	}
	return nil

}

// converter
func (o *stringOptional) String() string {
	return o.v
}
func (o *stringOptional) StringDefault(s string) string {
	if o.err != nil || o.v == "" {
		return s
	}
	return o.v
}
func (o *stringOptional) SliceString(sep string) []string {
	return strings.Split(o.v, sep)
}
func (o *stringOptional) Int() (int, error) {
	return strconv.Atoi(o.v)
}
func (o *stringOptional) Int64() (int64, error) {
	return strconv.ParseInt(o.v, 10, 64)
}
func (o *stringOptional) Float64() (float64, error) {
	return strconv.ParseFloat(o.v, 64)
}
func (o *stringOptional) Length() int {
	return len(o.v)
}
func (o *stringOptional) LengthUTF8() int {
	return utf8.RuneCountInString(o.v)
}
func (o *stringOptional) Uint64() (uint64, error) {
	return strconv.ParseUint(o.v, 10, 0)
}



func (o *stringOptional) SliceStringOptional(sep string) *sliceStringOptional {
	s := &sliceStringOptional{err: o.err}
	s.v = strings.Split(o.v, sep)
	return s
}

func (o *stringOptional) IntOptional() *intOptional {
	i := &intOptional{err: o.err}
	i.v, i.err = strconv.Atoi(o.v)
	return i
}

func (o *stringOptional) Int64Optional() *int64Optional {
	i := &int64Optional{err: o.err}
	i.v, i.err = strconv.ParseInt(o.v, 10, 64)
	return i
}

func (o *stringOptional) Float64Optional() *float64Optional {
	f := &float64Optional{}
	f.v, f.err = strconv.ParseFloat(o.v, 64)
	return f
}

func (o *stringOptional) UintOptional() *uintOptional {
	u := &uintOptional{err: o.err}
	v, err := strconv.ParseUint(o.v, 10, 0)
	u.err = err
	u.v = uint(v)
	return u
}
