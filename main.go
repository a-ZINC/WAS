package main

import (
	"github.com/a-ZINC/WAS/cmd/was"
	"github.com/a-ZINC/WAS/internal/agent"
	"github.com/a-ZINC/WAS/internal/broker"
)

func main() {
	was.Execute()

	if was.IsAgent {
		agent.StartAgent()
	}
	if was.IsBroker {
		broker.StartBroker(was.Port)
	}
}
