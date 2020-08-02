package optional

import (
	"encoding/base64"
	"strings"
)

func RemoveSpace() Apply {
	return func(val *processor) error {
		return nil
	}
}

func ToUpper() Apply {
	return func(val *processor) error {
		if err := val.value.GetError(); err != nil {
			return err
		}
		val.value.v = strings.ToUpper(val.value.v.(string))
		return nil
	}
}

func ToInt() Apply {
	return func(val *processor) error {
		i, err := val.value.Converter().Int()
		if err != nil {
			return err
		}
		val.value = IntVal(i)
		return nil
	}
}

func Base64StdEncode() Apply {
	return func(val *processor) error {
		if err := val.value.GetError(); err != nil {
			return err
		}
		if val.value.isString() {
			val.value.v = base64.StdEncoding.EncodeToString(strToBytes(val.value.v.(string)))
			return nil
		}
		return errorf(FmtMustString, val.name)
	}
}

func Trim(s string) Apply {
	return func(val *processor) error {
		return nil
	}
}
