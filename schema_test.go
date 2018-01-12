package jsonschema

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestDraft3(t *testing.T) {
	runJSONTests(t, []string{
	// "testdata/draft3/additionalItems.json",
	// "testdata/draft3/disallow.json",
	// "testdata/draft3/items.json",
	// "testdata/draft3/minItems.json",
	// "testdata/draft3/pattern.json",
	// "testdata/draft3/refRemote.json",
	// "testdata/draft3/additionalProperties.json",
	// "testdata/draft3/divisibleBy.json",
	// "testdata/draft3/maxItems.json",
	// "testdata/draft3/minLength.json",
	// "testdata/draft3/patternProperties.json",
	// "testdata/draft3/required.json",
	// "testdata/draft3/default.json",
	// "testdata/draft3/enum.json",
	// "testdata/draft3/maxLength.json",
	// "testdata/draft3/minimum.json",
	// "testdata/draft3/properties.json",
	// "testdata/draft3/type.json",
	// "testdata/draft3/dependencies.json",
	// "testdata/draft3/extends.json",
	// "testdata/draft3/maximum.json",
	// "testdata/draft3/ref.json",
	// "testdata/draft3/uniqueItems.json",
	// "testdata/draft3/optional/bignum.json",
	// "testdata/draft3/optional/format.json",
	// "testdata/draft3/optional/jsregex.json",
	// "testdata/draft3/optional/zeroTerminatedFloats.json",
	})
}

func TestDraft4(t *testing.T) {
	runJSONTests(t, []string{
		// "testdata/draft4/additionalItems.json",
		// "testdata/draft4/definitions.json",
		// "testdata/draft4/maxLength.json",
		// "testdata/draft4/minProperties.json",
		// "testdata/draft4/refRemote.json",
		// "testdata/draft4/additionalProperties.json",
		// "testdata/draft4/dependencies.json",
		// "testdata/draft4/maxProperties.json",
		// "testdata/draft4/minimum.json",
		// "testdata/draft4/pattern.json",
		// "testdata/draft4/required.json",
		// "testdata/draft4/allOf.json",
		// "testdata/draft4/enum.json",
		// "testdata/draft4/maximum.json",
		// "testdata/draft4/multipleOf.json",
		// "testdata/draft4/patternProperties.json",
		"testdata/draft4/type.json",
		// "testdata/draft4/anyOf.json",
		// "testdata/draft4/items.json",
		// "testdata/draft4/minItems.json",
		// "testdata/draft4/not.json",
		// "testdata/draft4/properties.json",
		// "testdata/draft4/uniqueItems.json",
		// "testdata/draft4/default.json",
		// "testdata/draft4/maxItems.json",
		// "testdata/draft4/minLength.json",
		// "testdata/draft4/oneOf.json",
		// "testdata/draft4/ref.json",
		// "testdata/draft4/optional/bignum.json",
		// "testdata/draft4/optional/ecmascript-regex.json",
		// "testdata/draft4/optional/format.json",
		// "testdata/draft4/optional/zeroTerminatedFloats.json",
	})
}

func TestDraft6(t *testing.T) {
	runJSONTests(t, []string{
		// "testdata/draft6/additionalItems.json",
		"testdata/draft6/const.json",
		"testdata/draft6/enum.json",
		// "testdata/draft6/maxLength.json",
		// "testdata/draft6/minProperties.json",
		// "testdata/draft6/ref.json",
		// "testdata/draft6/additionalProperties.json",
		// "testdata/draft6/contains.json",
		// "testdata/draft6/exclusiveMaximum.json",
		// "testdata/draft6/maxProperties.json",
		// "testdata/draft6/minimum.json",
		// "testdata/draft6/pattern.json",
		// "testdata/draft6/refRemote.json",
		// "testdata/draft6/allOf.json",
		// "testdata/draft6/default.json",
		// "testdata/draft6/exclusiveMinimum.json",
		// "testdata/draft6/maximum.json",
		// "testdata/draft6/multipleOf.json",
		// "testdata/draft6/patternProperties.json",
		// "testdata/draft6/required.json",
		// "testdata/draft6/anyOf.json",
		// "testdata/draft6/definitions.json",
		// "testdata/draft6/items.json",
		// "testdata/draft6/minItems.json",
		// "testdata/draft6/not.json",
		// "testdata/draft6/properties.json",
		"testdata/draft6/type.json",
		// "testdata/draft6/boolean_schema.json",
		// "testdata/draft6/dependencies.json",
		// "testdata/draft6/maxItems.json",
		// "testdata/draft6/minLength.json",
		// "testdata/draft6/oneOf.json",
		// "testdata/draft6/propertyNames.json",
		// "testdata/draft6/uniqueItems.json",
		// "testdata/draft6/optional/bignum.json",
		// "testdata/draft6/optional/ecmascript-regex.json",
		// "testdata/draft6/optional/format.json",
		// "testdata/draft6/optional/zeroTerminatedFloats.json",
	})
}

func TestDraft7(t *testing.T) {
	runJSONTests(t, []string{
		"testdata/draft7/additionalItems.json",
		"testdata/draft7/contains.json",
		"testdata/draft7/exclusiveMinimum.json",
		"testdata/draft7/maximum.json",
		"testdata/draft7/not.json",
		"testdata/draft7/propertyNames.json",
		"testdata/draft7/additionalProperties.json",
		// "testdata/draft7/default.json",
		// "testdata/draft7/if-then-else.json",
		"testdata/draft7/minItems.json",
		"testdata/draft7/oneOf.json",
		// "testdata/draft7/ref.json",
		"testdata/draft7/allOf.json",
		// "testdata/draft7/definitions.json",
		"testdata/draft7/items.json",
		"testdata/draft7/minLength.json",
		// "testdata/draft7/refRemote.json",
		"testdata/draft7/anyOf.json",
		// "testdata/draft7/dependencies.json",
		"testdata/draft7/maxItems.json",
		"testdata/draft7/minProperties.json",
		"testdata/draft7/pattern.json",
		"testdata/draft7/required.json",
		"testdata/draft7/boolean_schema.json",
		"testdata/draft7/enum.json",
		"testdata/draft7/maxLength.json",
		"testdata/draft7/minimum.json",
		"testdata/draft7/patternProperties.json",
		"testdata/draft7/type.json",
		"testdata/draft7/const.json",
		"testdata/draft7/exclusiveMaximum.json",
		"testdata/draft7/maxProperties.json",
		"testdata/draft7/multipleOf.json",
		"testdata/draft7/properties.json",
		"testdata/draft7/uniqueItems.json",

		// "testdata/draft7/optional/bignum.json",
		// "testdata/draft7/optional/content.json",
		// "testdata/draft7/optional/ecmascript-regex.json",
		// "testdata/draft7/optional/zeroTerminatedFloats.json",
		// "testdata/draft7/optional/format/date-time.json",
		// "testdata/draft7/optional/format/hostname.json",
		// "testdata/draft7/optional/format/ipv4.json",
		// "testdata/draft7/optional/format/iri.json",
		// "testdata/draft7/optional/format/relative-json-pointer.json",
		// "testdata/draft7/optional/format/uri-template.json",
		// "testdata/draft7/optional/format/date.json",
		// "testdata/draft7/optional/format/idn-email.json",
		// "testdata/draft7/optional/format/ipv6.json",
		// "testdata/draft7/optional/format/json-pointer.json",
		// "testdata/draft7/optional/format/time.json",
		// "testdata/draft7/optional/format/uri.json",
		// "testdata/draft7/optional/format/email.json",
		// "testdata/draft7/optional/format/idn-hostname.json",
		// "testdata/draft7/optional/format/iri-reference.json",
		// "testdata/draft7/optional/format/regex.json",
		// "testdata/draft7/optional/format/uri-reference.json",
	})
}

// TestSet is a json-based set of tests
// JSON-Schema comes with a lovely JSON-based test suite:
// https://github.com/json-schema-org/JSON-Schema-Test-Suite
type TestSet struct {
	Description string      `json:"description"`
	Schema      *RootSchema `json:"schema"`
	Tests       []TestCase  `json:"tests"`
}

type TestCase struct {
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
	Valid       bool        `json:"valid"`
}

func runJSONTests(t *testing.T, testFilepaths []string) {
	tests := 0
	passed := 0
	for _, path := range testFilepaths {
		base := filepath.Base(path)
		testSets := []*TestSet{}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("error loading test file: %s", err.Error())
			return
		}

		if err := json.Unmarshal(data, &testSets); err != nil {
			t.Errorf("error unmarshaling test set %s from JSON: %s", base, err.Error())
			return
		}

		for _, ts := range testSets {
			sc := ts.Schema
			for i, c := range ts.Tests {
				tests++
				got := sc.Validate(c.Data)
				valid := got == nil
				if valid != c.Valid {
					t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, got)
				} else {
					passed++
				}
			}
		}
	}
	t.Logf("%d/%d tests passed", passed, tests)
}

func TestDataType(t *testing.T) {
	cases := []struct {
		data   interface{}
		expect string
	}{
		{nil, "null"},
		{struct{}{}, "unknown"},
		{float64(4), "integer"},
		{float64(4.5), "number"},
		{"foo", "string"},
		{map[string]interface{}{}, "object"},
		{[]interface{}{}, "array"},
	}

	for i, c := range cases {
		got := DataType(c.data)
		if got != c.expect {
			t.Errorf("case %d result mismatch. expected: '%s', got: '%s'", i, c.expect, got)
		}
	}
}