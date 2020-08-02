package optional

import (
	"fmt"
	"strconv"
	"testing"
)

func TestType_IsPrimitiveType(t *testing.T) {
	tests := []struct {
		Type Type
		Want bool
	}{
		{String, true},
		{Int, true},
		{Bool, true},
		{StringMap(), false},

		// Make sure our primitive constants are correctly constructed
		{True.Type(), true},
		{False.Type(), true},
		{Zero.Type(), true},
		{PositiveInfinity.Type(), true},
		{NegativeInfinity.Type(), true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d %#value", i, test.Type), func(t *testing.T) {
			got := test.Type.IsPrimitiveType()
			if got != test.Want {
				t.Errorf(
					"wrong result\ntype: %#value\ngot:  %#value\nwant: %#value",
					test.Type,
					test.Want, got,
				)
			}
		})
	}
}

func TestConverter(t *testing.T) {
	var i10 = 10
	if i, err := IntVal(i10).Converter().Int(); err != nil || i != i10 {
		t.Errorf("wrong result\n Func: %s\ngot %d\n want:%d", "Inval()", i, i10)
	}

	var i8 int8 = 127
	if i, err := Int8Val(i8).Converter().Int8(); err != nil || i != i8 {
		t.Errorf("wrong result\n Func: %s\ngot %d\n want:%d", "Int8Val()", i, i8)
	}

	var i16 int16 = 127
	if i, err := Int16Val(i16).Converter().Int16(); err != nil || i != i16 {
		t.Errorf("wrong result\n Func: %s\ngot %d\n want:%d", "Int16Val()", i, i16)
	}
	var i32 int32 = 1<<31 - 1
	if i, err := Int32Val(i32).Converter().Int32(); err != nil || i != i32 {
		t.Errorf("wrong result\n Func: %s\ngot %d\n want:%d", "Int32Val()", i, i32)
	}

	var i64 int64 = 1<<63 - 1
	if i, err := Int64Val(i64).Converter().Int64(); err != nil || i != i64 {
		t.Errorf("wrong result\n Func: %s\ngot %d\n want:%d", "Int64Val()", i, i64)
	}

	f64, err := strconv.ParseFloat("9.223372036854776E+18", 64)
	if err != nil {
		t.Error(err)
	}
	if i, err := Float64Val(f64).Converter().Float64(); err != nil || i != f64 {
		t.Errorf("wrong result\n Func: %s\ngot %f\n want:%f", "Float64Val()", i, f64)
	}
	var f32 float32 = 123.3
	if err != nil {
		t.Error(err)
	}
	if i, err := Float32Val(f32).Converter().Float32(); err != nil || i != f32 {
		t.Errorf("wrong result\n Func: %s\ngot %f\n want:%f", "Float32Val()", i, f32)
	}

	var s = "hello world"
	if i, err := StringVal(s).Converter().String(); err != nil || i != s {
		t.Errorf("wrong result\n Func: %s\ngot %s\n want:%s", "StringVal()", i, s)
	}
	var want = "{\"hello\":\"hello world\"}"
	if i, err := MapStringVal(map[string]Value{"hello": StringVal(s)}).Converter().String(); err != nil || i != want {
		t.Errorf("wrong result\n Func: %s\ngot %s\n want:%s", "StringVal()", i, want)
	}

}

func TestValidator(t *testing.T) {
	tests := []struct {
		Value Value
		Want  bool
	}{
		{StringVal("name"), true},
		{IntVal(12), true},
		{BoolVal(true), true},
		{MapStringValEmpty(), true},
		{MapStringVal(map[string]Value{
			"status": BoolVal(true),
			"name":   StringVal("ggg"),
		}), false},
	}

	for i := range tests {
		if !tests[i].Value.IsNull() && tests[i].Value.IsPrimitiveValue() {
			if err := tests[i].Value.Validate("字段名", MustNotNil()).GetError(); err != nil && !tests[i].Want {
				t.Errorf("wrong result\ntype: %#value \n err:%value", tests[i].Value, err)
			}
		}
		if !tests[i].Value.IsNull() && tests[i].Value.IsMapValue() {
			if err := tests[i].Value.Validates(
				Validate("status", MustTrue()),
				Validate("name", MustString(), MustNotNil(), MustHasLetter()),
			).GetError(); err != nil && !tests[i].Want {
				t.Error(err)
			}
		}
	}
}

func TestProcessor(t *testing.T) {
	//todo
}

func TestAlign(t *testing.T) {
	// todo
}
