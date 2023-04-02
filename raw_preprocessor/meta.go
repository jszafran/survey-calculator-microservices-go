package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// QuestionMetadata represents metadata of a survey question:
// text, column name in raw data (header), is nullable, minimal value, maximal value.
type QuestionMetadata struct {
	Text       string `yaml:"text"`
	ColumnName string `yaml:"column_name"`
	Nullable   bool   `yaml:"nullable"`
	MinValue   int    `yaml:"min_value"`
	MaxValue   int    `yaml:"max_value"`
}

// DemographicMetadata represents metadata of a demographic:
// text, column name in raw data (header), is nullable, minimal value, maximal value.
type DemographicMetadata struct {
	Text       string `yaml:"text"`
	ColumnName string `yaml:"column_name"`
	Nullable   bool   `yaml:"nullable"`
	MinValue   int    `yaml:"min_value"`
	MaxValue   int    `yaml:"max_value"`
}

// Schema represents a schema of survey file: name of the organizational node column,
// along with info about survey questions and demographics.
type Schema struct {
	OrgNodeColumnName string                `yaml:"org_node_column_name"`
	Questions         []QuestionMetadata    `yaml:"questions"`
	Demographics      []DemographicMetadata `yaml:"demographics"`
}

// SchemaFromYAML parses a yaml file into a Schema object.
func SchemaFromYAML(path string) (Schema, error) {
	var schema Schema
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return schema, err
	}

	err = yaml.Unmarshal(f, &schema)
	if err != nil {
		return schema, err
	}

	return schema, nil
}
