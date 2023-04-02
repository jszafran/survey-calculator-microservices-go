package main

import (
	"reflect"
	"testing"
)

func TestMeta_SchemaFromYAML(t *testing.T) {
	type testCase struct {
		path     string
		expected Schema
	}

	questions := []QuestionMetadata{
		{
			Text:       "Question 1 text",
			ColumnName: "q1",
			Nullable:   true,
			MinValue:   1,
			MaxValue:   3,
		},
		{
			Text:       "Question 2 text",
			ColumnName: "q2",
			Nullable:   true,
			MinValue:   1,
			MaxValue:   4,
		},
		{
			Text:       "Question 3 text",
			ColumnName: "q3",
			Nullable:   false,
			MinValue:   1,
			MaxValue:   5,
		},
	}
	demographics := []DemographicMetadata{
		{
			Text:       "Demographic 1 text",
			ColumnName: "d1",
			Nullable:   true,
			MinValue:   1,
			MaxValue:   3,
		},
		{
			Text:       "Demographic 2 text",
			ColumnName: "d2",
			Nullable:   false,
			MinValue:   1,
			MaxValue:   4,
		},
	}

	testCases := []testCase{
		{
			path: "test_resources/test_schema_multiple_questions.yaml",
			expected: Schema{
				OrgNodeColumnName: "org_node",
				Questions:         questions,
				Demographics:      demographics,
			},
		},
		{
			path: "test_resources/test_schema_single_question.yaml",
			expected: Schema{
				OrgNodeColumnName: "org_node",
				Questions:         questions[:1],
				Demographics:      demographics[:1],
			},
		},
	}

	for _, tc := range testCases {
		got, err := SchemaFromYAML(tc.path)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, tc.expected) {
			t.Fatalf("Expected %+v but got %+v", tc.expected, got)
		}
	}
}
