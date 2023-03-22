package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	netClient = &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 15 * time.Second,
		},
	}
)

func ping() error {
	u, err := url.JoinPath(destURL, "/healthz")
	if err != nil {
		return err
	}
	r, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	requestWithHeaders(r)
	requestWithAuth(r)
	resp, err := netClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := wantsStatusCode(resp, http.StatusOK); err != nil {
		return err
	}
	return nil
}

func deleteProvisionedRule(uid string) error {
	u, err := url.JoinPath(destURL, "/api/v1/provisioning/alert-rules/", uid)
	if err != nil {
		return err
	}
	r, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	requestWithHeaders(r)
	requestWithAuth(r)
	resp, err := netClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := wantsStatusCode(resp, http.StatusNoContent); err != nil {
		return err
	}
	log.Printf("Deleted %s", uid)
	return nil
}

func deleteProvisionedRules() {
	rules, err := listProvisionedRules()
	if err != nil {
		log.Fatalf("Failed to list rules: %s", err)
	}
	for _, rule := range rules {
		if err = deleteProvisionedRule(rule.UID); err != nil {
			log.Printf("Failed to delete rule %s: %s", rule.UID, err)
		}
	}
	log.Println("Done")
}

func listProvisionedRules() ([]ProvisionedAlertRule, error) {
	u, err := url.JoinPath(destURL, "/api/v1/provisioning/alert-rules")
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	requestWithHeaders(r)
	requestWithAuth(r)
	resp, err := netClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := wantsStatusCode(resp, http.StatusOK); err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []ProvisionedAlertRule
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func provisionRule(v ProvisionedAlertRule) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	u, err := url.JoinPath(destURL, "/api/v1/provisioning/alert-rules")
	if err != nil {
		return err
	}
	r, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return err
	}
	requestWithHeaders(r)
	requestWithAuth(r)
	resp, err := netClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := wantsStatusCode(resp, http.StatusCreated); err != nil {
		return err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := ProvisionedAlertRule{}
	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}
	log.Printf("Provisioned %s", result.UID)
	return nil
}

func provisionRules(v ProvisionedAlertRule) {
	for i := offset; i < offset+rules; i++ {
		v.Title = "test-" + strconv.Itoa(i)
		v.RuleGroup = "test-" + strconv.Itoa(i/rulesPerGroup+1)
		if err := provisionRule(v); err != nil {
			log.Fatalf("Failed to provision alert rule: %s", err)
		}
	}
	log.Println("Done")
}

func requestWithHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "alertbench/0.1")
}

func requestWithAuth(r *http.Request) {
	if token != "" {
		r.Header.Set("Authorization", "Bearer "+token)
	} else {
		r.SetBasicAuth(username, password)
	}
}

func wantsStatusCode(r *http.Response, wants int) error {
	if r.StatusCode != wants {
		return fmt.Errorf("wants HTTP %d, got %d", wants, r.StatusCode)
	}
	return nil
}
