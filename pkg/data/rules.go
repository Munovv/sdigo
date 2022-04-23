package data

import (
	"encoding/json"
	"io/ioutil"
)

type Rule struct {
	Value string `json:"rule_value" validate:"required"`
}

func ParseRules(filepath string) (*[]Rule, error) {
	var rules *[]Rule
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(file, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}
