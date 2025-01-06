package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Groups []Group `hcl:"group,block"`
}

type Group struct {
	Port     int      `hcl:"port"`
	Strategy Strategy `hcl:"strategy"`
	Targets  []Target `hcl:"target,block"`
}

type Strategy string

const (
	RandomStrategy     Strategy = "random"
	RoundRobinStrategy Strategy = "roundRobin"
	LeastConnStrategy  Strategy = "leastConn"
)

type Target struct {
	Addr    string `hcl:"addr"`
	Primary bool   `hcl:"primary,optional"`
}

func Parse(path string) (Config, error) {
	var config Config
	err := hclsimple.DecodeFile(path, nil, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func main() {
	cfg, err := Parse("./cmd/hcl/tcproxy.hcl")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
