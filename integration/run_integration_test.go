package integration

import (
	"github.com/calyptia/fluent-bit-ci/integration/tests"
	"github.com/calyptia/fluent-bit-ci/integration/tests/elasticsearch"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestFluentBitSuites(t *testing.T) {
	suite.Run(t, &elasticsearch.Suite{BaseTestSuite: tests.BaseTestSuite{Name: "elasticsearch"}})
}
