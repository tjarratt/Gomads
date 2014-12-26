package gomads_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/gomads"
)

var _ = Describe("Monads", func() {
	Context("Maybe", func() {
		It("can sometimes contain a value", func() {
			m := Maybe(func() interface{} {
				return nil
			})

			Expect(m.Value()).To(BeNil())
		})
	})
})
