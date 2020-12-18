package structtags

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

type testStruct struct {
	Name              string     `tf5:"name"`
	Age               *big.Float `tf5:"my_age"`
	dontUse           bool
	ExplicitlyDontUse string          `tf5:"-"`
	List              []tftypes.Value `tf5:"list"`
	InferredTag       map[string]tftypes.Value
	Embedded          embeddedStruct
}

func (t *testStruct) FromTerraform5Value(value tftypes.Value) error {
	return StructFromTerraform5Value(t, value)
}

type embeddedStruct struct {
	Foo string     `tf5:"foo"`
	Bar *big.Float `tf5:"bar"`
}

func (e *embeddedStruct) FromTerraform5Value(value tftypes.Value) error {
	return StructFromTerraform5Value(e, value)
}

func TestStructFromTerraform5Value(t *testing.T) {
	var test testStruct
	val := tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"name":   tftypes.String,
			"my_age": tftypes.Number,
			"list": tftypes.List{
				ElementType: tftypes.String,
			},
			"inferred_tag": tftypes.Map{
				AttributeType: tftypes.String,
			},
			"embedded": tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"foo": tftypes.String,
					"bar": tftypes.Number,
				},
			},
		},
	}, map[string]tftypes.Value{
		"name":   tftypes.NewValue(tftypes.String, "Johnny Appleseed"),
		"my_age": tftypes.NewValue(tftypes.Number, big.NewFloat(123)),
		"list": tftypes.NewValue(tftypes.List{
			ElementType: tftypes.String,
		}, []tftypes.Value{
			tftypes.NewValue(tftypes.String, "foo"),
			tftypes.NewValue(tftypes.String, "bar"),
			tftypes.NewValue(tftypes.String, "baz"),
		}),
		"inferred_tag": tftypes.NewValue(tftypes.Map{
			AttributeType: tftypes.String,
		}, map[string]tftypes.Value{
			"red":   tftypes.NewValue(tftypes.String, "a"),
			"blue":  tftypes.NewValue(tftypes.String, "b"),
			"green": tftypes.NewValue(tftypes.String, "c"),
		}),
		"embedded": tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"foo": tftypes.String,
				"bar": tftypes.Number,
			},
		}, map[string]tftypes.Value{
			"foo": tftypes.NewValue(tftypes.String, "hello"),
			"bar": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
		}),
	})
	err := val.As(&test)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(testStruct{
		Name: "Johnny Appleseed",
		Age:  big.NewFloat(123),
		List: []tftypes.Value{
			tftypes.NewValue(tftypes.String, "foo"),
			tftypes.NewValue(tftypes.String, "bar"),
			tftypes.NewValue(tftypes.String, "baz"),
		},
		InferredTag: map[string]tftypes.Value{
			"red":   tftypes.NewValue(tftypes.String, "a"),
			"blue":  tftypes.NewValue(tftypes.String, "b"),
			"green": tftypes.NewValue(tftypes.String, "c"),
		},
		Embedded: embeddedStruct{
			Foo: "hello",
			Bar: big.NewFloat(1234),
		},
	}, test, numberComparer(), tftypes.ValueComparer(), cmpopts.IgnoreFields(testStruct{}, "dontUse")); diff != "" {
		t.Errorf("Unexpected diff (- wanted, + got): %s", diff)
	}
}
