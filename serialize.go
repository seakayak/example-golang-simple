package main

import (
	"encoding/json"
	"os"
)

func loadIPs(path string) (map[string]struct{}, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return make(map[string]struct{}), nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ips []string
	err = json.NewDecoder(f).Decode(&ips)
	if err != nil {
		return nil, err
	}

	m := make(map[string]struct{})
	for _, ip := range ips {
		m[ip] = struct{}{}
	}
	return m, nil
}

func saveIPs(uniqueIPs map[string]struct{}, path string) error {
	var ips []string
	for ip := range uniqueIPs {
		ips = append(ips, ip)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(ips)
}
