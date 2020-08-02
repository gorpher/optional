package optional

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strings"
	"unicode"
)

const (
	FmtMustNotNil          = "%s filed value must is not nil"
	FmtMustTrue            = "%s filed value must is true"
	FmtMustHasSuffix       = "%s filed value must has %s suffix"
	FmtMustString          = "%s filed value must is a string type"
	FmtMustHasString       = "%s filed value must has %s string value"
	FmtMustHasSymbol       = "%s filed value must has symbol"
	FmtMustHasDigit        = "%s filed value must has digit"
	FmtMustHasLetter       = "%s filed value must has letter"
	FmtMustHasLower        = "%s filed value must has lower"
	FmtMustHasUpper        = "%s filed value must has upper"
	FmtMustIn              = "%s filed value must in %value upper"
	FmtMustEquals          = "%s filed value must equals %s"
	FmtMustIsLower         = "%s filed value must is lower"
	FmtMustIsUpper         = "%s filed value must is upper"
	FmtMustIsLetter        = "%s filed value must is letter"
	FmtMustIsDigit         = "%s filed value must is digit"
	FmtMustIsLowerOrDigit  = "%s filed value must is lower or digit"
	FmtMustIsUpperOrDigit  = "%s filed value must is upper or digit"
	FmtMustIsLetterOrDigit = "%s filed value must is letter or digit"
	FmtMustIsChinese       = "%s filed value must is chinese"
	FmtMustIsUUID          = "%s filed value must is uuid"
	FmtMustIsSQLObject     = "%s filed value must is sql "
	FmtMustIsIp            = "%s filed value must is ip "
	FmtMustIsNumber        = "%s filed value must is number type"
)

func MustNotNil() Match {
	return func(val *validator) error {
		if val.value.ty.IsPrimitiveType() {
			switch val.value.v.(type) {
			case string:
				if val.value.v.(string) != "" {
					return nil
				}
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				if val.value.v.(int) != 0 {
					return nil
				}
			case bool:
				if val.value.v.(bool) {
					return nil
				}
			}
		}
		return fmt.Errorf(FmtMustNotNil, val.name)
	}
}
func MustString() Match {
	return func(val *validator) error {
		if val.value.ty == String {
			return nil
		}
		return fmt.Errorf(FmtMustString, val.name)
	}
}
func MustTrue() Match {
	return func(val *validator) error {
		if val.value.Equals(True) {
			return nil
		}
		return errorf(FmtMustTrue, val.name)
	}
}
func MustHasSuffix(s string) Match {
	return func(val *validator) error {
		if val.value.ty == String {
			if strings.HasSuffix(val.value.v.(string), s) {
				return nil
			}
			return errorf(FmtMustHasSuffix, val.name, s)
		}
		if val.value.Equals(True) {
			return nil
		}
		return errorf(FmtMustHasSuffix, val.name, s)
	}
}
func MustHasString(s string) Match {
	return func(val *validator) error {
		if val.value.ty == String {
			if strings.Contains(val.value.v.(string), s) {
				return nil
			}
			return errorf(FmtMustHasString, val.name, s)
		}
		if val.value.Equals(True) {
			return nil
		}
		return errorf(FmtMustHasString, val.name, s)
	}
}
func MustHasSymbol() Match {
	return func(val *validator) error {
		if val.has(unicode.IsSymbol) {
			return nil
		}
		return errorf(FmtMustHasSymbol, val.name)
	}
}
func MustHasDigit() Match {
	return func(val *validator) error {
		if val.has(unicode.IsDigit) {
			return nil
		}
		return errorf(FmtMustHasDigit, val.name)
	}
}
func MustHasLetter() Match {
	return func(val *validator) error {
		if val.has(unicode.IsLetter) {
			return nil
		}
		return errorf(FmtMustHasLetter, val.name)
	}
}
func MustHasLower() Match {
	return func(val *validator) error {
		if val.has(unicode.IsLower) {
			return nil
		}
		return errorf(FmtMustHasLower, val.name)
	}
}
func MustHasUpper() Match {
	return func(val *validator) error {
		if val.has(unicode.IsUpper) {
			return nil
		}
		return errorf(FmtMustHasUpper, val.name)
	}
}
func MustIn(s []string) Match {
	return func(val *validator) error {
		for i := range s {
			if s[i] == val.value.v.(string) {
				return nil
			}
		}
		return errorf(FmtMustIn, val.name, s)
	}
}
func MustEquals(s string) Match {
	return func(val *validator) error {
		if val.value.v.(string) == s {
			return nil
		}
		return errorf(FmtMustEquals, val.name, s)
	}

}
func MustIsLower() Match {
	return func(val *validator) error {
		return isStringFunc(val, unicode.IsLower, FmtMustIsLower, val.name)
	}
}
func MustIsUpper() Match {
	return func(val *validator) error {
		return isStringFunc(val, unicode.IsUpper, FmtMustIsUpper, val.name)
	}
}
func MustIsLetter() Match {
	return func(val *validator) error {
		return isStringFunc(val, unicode.IsLetter, FmtMustIsLetter, val.name)
	}
}
func MustIsDigit() Match {
	return func(val *validator) error {
		return isStringFunc(val, unicode.IsDigit, FmtMustIsDigit, val.name)
	}
}
func MustIsLowerOrDigit() Match {
	return func(val *validator) error {
		return isStringFunc(val, func(r rune) bool {
			if unicode.IsLower(r) || unicode.IsDigit(r) {
				return true
			}
			return false
		}, FmtMustIsLowerOrDigit, val.name)
	}
}
func MustIsUpperOrDigit() Match {
	return func(val *validator) error {
		return isStringFunc(val, func(r rune) bool {
			if unicode.IsUpper(r) || unicode.IsDigit(r) {
				return true
			}
			return false
		}, FmtMustIsUpperOrDigit, val.name)
	}
}
func MustIsLetterOrDigit() Match {
	return func(val *validator) error {
		return isStringFunc(val, func(r rune) bool {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return true
			}
			return false
		}, FmtMustIsLetterOrDigit, val.name)
	}
}
func MustIsChinese() Match {
	return func(val *validator) error {
		return isStringFunc(val, func(r rune) bool {
			if unicode.Is(unicode.Scripts["Han"], r) {
				return true
			}
			return false
		}, FmtMustIsChinese, val.name)
	}
}
func MustIsURL() Match {
	return func(val *validator) error {
		if _, err := url.ParseRequestURI(val.value.v.(string)); err != nil {

			// todo err优化
			return err
		}
		return nil
	}
}
func MustIsUUID() Match {
	return func(val *validator) error {
		str := val.value.v.(string)
		//todo 验证完善uuid
		var uuid [16]byte
		switch len(str) {
		// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		case 36:

		// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		case 36 + 9:
			if strings.EqualFold(strings.ToLower(str[:9]), "urn:uuid:") {
				return errorf(FmtMustIsUUID, val.name)
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
					return errorf(FmtMustIsUUID, val.name)
				}
			}
			return nil
		default:
			return errorf(FmtMustIsUUID, val.name)
		}
		// s is now at least 36 bytes long
		// it must be of the form  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
			return errorf(FmtMustIsUUID, val.name)
		}
		for i, x := range [16]int{
			0, 2, 4, 6,
			9, 11,
			14, 16,
			19, 21,
			24, 26, 28, 30, 32, 34} {
			v, ok := xtob(str[x], str[x+1])
			if !ok {
				return errorf(FmtMustIsUUID, val.name)
			}
			uuid[i] = v
		}
		return nil
	}
}
func MustIsSQLObject() Match {
	return func(val *validator) error {
		return isStringFunc(val, func(r rune) bool {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_' {
				return false
			}
			return true
		}, FmtMustIsSQLObject, val.name)
	}
}
func MustIsChinaMobile() Match {
	return func(val *validator) error {
		//todo
		return nil
	}
}
func MustIsJSON() Match {
	return func(val *validator) error {
		var js json.RawMessage
		if err := json.Unmarshal(val.value.v.([]byte), &js); err != nil {
			//todo bytes and err
			return err
		}
		return nil
	}
}
func MustIsIP() Match {
	return func(val *validator) error {
		if v := net.ParseIP(val.value.v.(string)); v != nil {
			//todo ip的具体判断
			return nil
		}
		return errorf(FmtMustIsIp, val.name)
	}
}
func MustIsEmail() Match {
	return func(val *validator) error {
		//todo email
		return nil
	}
}
func MustIsNumberValue() Match {
	return func(val *validator) error {
		if val.value.isNumber() {
			return nil
		}
		if val.value.isString() {
			if _, err := val.value.Converter().Int64(); err == nil {
				return nil
			}
		}
		return errorf(FmtMustIsNumber, val.name)
	}
}

func isStringFunc(val *validator, fn func(r rune) bool, msg string, a ...interface{}) error {
	if val.is(fn) {
		return nil
	}
	return errorf(msg, a...)
}
func errorf(msg string, a ...interface{}) error {
	return fmt.Errorf(msg, a...)
}
