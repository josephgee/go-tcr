package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("does nothing", func() {
	It("really nothing", func() {
		Expect(true).To(Equal(true))
	})
})
