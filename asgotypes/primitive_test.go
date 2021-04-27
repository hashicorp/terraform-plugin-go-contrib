package asgotypes

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestGoPrimitive(t *testing.T) {
	type testCase struct {
		tfval    tftypes.Value
		expected interface{}
	}
	cases := map[string]testCase{
		"string": {
			tfval:    tftypes.NewValue(tftypes.String, "foo"),
			expected: "foo",
		},
		"number": {
			tfval:    tftypes.NewValue(tftypes.Number, big.NewFloat(123)),
			expected: big.NewFloat(123),
		},
		"bool": {
			tfval:    tftypes.NewValue(tftypes.Bool, true),
			expected: true,
		},
		"object-bool-string-number": {
			tfval: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.Bool,
					"b": tftypes.String,
					"c": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.Bool, true),
				"b": tftypes.NewValue(tftypes.String, "bar"),
				"c": tftypes.NewValue(tftypes.Number, big.NewFloat(456)),
			}),
			expected: map[string]interface{}{
				"a": true,
				"b": "bar",
				"c": big.NewFloat(456),
			},
		},
		"list-string": {
			tfval: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "foo"),
				tftypes.NewValue(tftypes.String, "bar"),
				tftypes.NewValue(tftypes.String, "baz"),
				tftypes.NewValue(tftypes.String, "quux"),
			}),
			expected: []string{"foo", "bar", "baz", "quux"},
		},
		"set-string": {
			tfval: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "foo"),
				tftypes.NewValue(tftypes.String, "bar"),
				tftypes.NewValue(tftypes.String, "baz"),
				tftypes.NewValue(tftypes.String, "quux"),
			}),
			expected: []string{"foo", "bar", "baz", "quux"},
		},
		"tuple-bool-string-number": {
			tfval: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{
					tftypes.Bool,
					tftypes.String,
					tftypes.Number,
				},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.String, "test"),
				tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
			}),
			expected: []interface{}{true, "test", big.NewFloat(1234)},
		},
		"map-string": {
			tfval: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.String,
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.String, "foo"),
				"b": tftypes.NewValue(tftypes.String, "bar"),
				"c": tftypes.NewValue(tftypes.String, "baz"),
			}),
			expected: map[string]string{
				"a": "foo",
				"b": "bar",
				"c": "baz",
			},
		},
		"list-map-set-object-string-string-bool": {
			tfval: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Map{
					AttributeType: tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					},
				},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Map{
					AttributeType: tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					},
				}, map[string]tftypes.Value{
					"foo": tftypes.NewValue(tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					}, []tftypes.Value{
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "hello"),
							"b": tftypes.NewValue(tftypes.String, "world"),
							"c": tftypes.NewValue(tftypes.Bool, true),
						}),
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "hey"),
							"b": tftypes.NewValue(tftypes.String, "jude"),
							"c": tftypes.NewValue(tftypes.Bool, false),
						}),
					}),
					"bar": tftypes.NewValue(tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					}, []tftypes.Value{
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "goodnight"),
							"b": tftypes.NewValue(tftypes.String, "moon"),
							"c": tftypes.NewValue(tftypes.Bool, true),
						}),
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "please"),
							"b": tftypes.NewValue(tftypes.String, "clap"),
							"c": tftypes.NewValue(tftypes.Bool, false),
						}),
					}),
				}),
				tftypes.NewValue(tftypes.Map{
					AttributeType: tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					},
				}, map[string]tftypes.Value{
					"baz": tftypes.NewValue(tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					}, []tftypes.Value{
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "red"),
							"b": tftypes.NewValue(tftypes.String, "blue"),
							"c": tftypes.NewValue(tftypes.Bool, false),
						}),
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "yellow"),
							"b": tftypes.NewValue(tftypes.String, "green"),
							"c": tftypes.NewValue(tftypes.Bool, true),
						}),
					}),
					"quux": tftypes.NewValue(tftypes.Set{
						ElementType: tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						},
					}, []tftypes.Value{
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "hungry"),
							"b": tftypes.NewValue(tftypes.String, "caterpillar"),
							"c": tftypes.NewValue(tftypes.Bool, true),
						}),
						tftypes.NewValue(tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"a": tftypes.String,
								"b": tftypes.String,
								"c": tftypes.Bool,
							},
						}, map[string]tftypes.Value{
							"a": tftypes.NewValue(tftypes.String, "giving"),
							"b": tftypes.NewValue(tftypes.String, "tree"),
							"c": tftypes.NewValue(tftypes.Bool, false),
						}),
					}),
				}),
			}),
			expected: []map[string][]map[string]interface{}{
				{
					"foo": {
						{
							"a": "hello",
							"b": "world",
							"c": true,
						},
						{
							"a": "hey",
							"b": "jude",
							"c": false,
						},
					},
					"bar": {
						{
							"a": "goodnight",
							"b": "moon",
							"c": true,
						},
						{
							"a": "please",
							"b": "clap",
							"c": false,
						},
					},
				},
				{
					"baz": {
						{
							"a": "red",
							"b": "blue",
							"c": false,
						},
						{
							"a": "yellow",
							"b": "green",
							"c": true,
						},
					},
					"quux": {
						{
							"a": "hungry",
							"b": "caterpillar",
							"c": true,
						},
						{
							"a": "giving",
							"b": "tree",
							"c": false,
						},
					},
				},
			},
		},
		"list-list-string": {
			tfval: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.List{
					ElementType: tftypes.String,
				},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.List{
					ElementType: tftypes.String,
				}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "a"),
					tftypes.NewValue(tftypes.String, "b"),
					tftypes.NewValue(tftypes.String, "c"),
				}),
				tftypes.NewValue(tftypes.List{
					ElementType: tftypes.String,
				}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "foo"),
					tftypes.NewValue(tftypes.String, "bar"),
					tftypes.NewValue(tftypes.String, "baz"),
				}),
				tftypes.NewValue(tftypes.List{
					ElementType: tftypes.String,
				}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "red"),
					tftypes.NewValue(tftypes.String, "yellow"),
					tftypes.NewValue(tftypes.String, "blue"),
				}),
			}),
			expected: [][]string{
				{"a", "b", "c"},
				{"foo", "bar", "baz"},
				{"red", "yellow", "blue"},
			},
		},
		"map-list-string": {
			tfval: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.List{
					ElementType: tftypes.String,
				},
			}, map[string]tftypes.Value{
				"hello": tftypes.NewValue(tftypes.List{
					ElementType: tftypes.String,
				}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "a"),
					tftypes.NewValue(tftypes.String, "b"),
					tftypes.NewValue(tftypes.String, "c"),
				}),
				"world": tftypes.NewValue(tftypes.List{
					ElementType: tftypes.String,
				}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "foo"),
					tftypes.NewValue(tftypes.String, "bar"),
					tftypes.NewValue(tftypes.String, "baz"),
				}),
			}),
			expected: map[string][]string{
				"hello": {"a", "b", "c"},
				"world": {"foo", "bar", "baz"},
			},
		},
	}

	for name, testCase := range cases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var res GoPrimitive
			err := testCase.tfval.As(&res)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(testCase.expected, res.Value, cmpOpts...); diff != "" {
				t.Errorf("Unexpected value (- wanted, + got): %s", diff)
			}
		})
	}
}
