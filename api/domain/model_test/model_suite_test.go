package model_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModelTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ModelTest Suite")
}
