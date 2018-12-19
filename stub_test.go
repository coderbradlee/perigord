package main

import (
	"testing"

	_ "perigord/tests"
	_ "perigord/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
