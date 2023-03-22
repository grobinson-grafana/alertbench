package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ProvisionedAlertRule struct {
	Annotations    map[string]string `json:"annotations"`
	Condition      string            `json:"condition"`
	Data           json.RawMessage   `json:"data"`
	ExecErrorState string            `json:"execErrState"`
	FolderUID      string            `json:"folderUID"`
	For            time.Duration     `json:"duration"`
	Labels         map[string]string `json:"labels"`
	NoDataState    string            `json:"noDataState`
	OrgID          int               `json:"orgId"`
	RuleGroup      string            `json:"ruleGroup"`
	Title          string            `json:"title"`
	UID            string            `json:"uid,omitempty"`
}

func ExampleFromFile(file string) (ProvisionedAlertRule, error) {
	var (
		b   []byte
		err error
		v   ProvisionedAlertRule
	)
	b, err = os.ReadFile(file)
	if err != nil {
		return v, fmt.Errorf("Failed to read file %s: %s", file, err)
	}
	if err = json.Unmarshal(b, &v); err != nil {
		return v, fmt.Errorf("Failed to unmarshal JSON: %s", err)
	}
	return v, nil
}
