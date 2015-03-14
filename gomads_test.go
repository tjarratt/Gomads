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

		It("can be tricked into returning something if nothing was provided", func() {
			m := Maybe(func() interface{} {
				return nil
			}).OrSome("fooled you!")

			Expect(m.Value()).To(Equal("fooled you!"))
		})

		It("can chain maybe-values together", func() {
			m := Maybe(func() interface{} {
				return nil
			}).OrSome(Maybe(func() interface{} {
				return "chained"
			}))

			Expect(m.Value()).To(Equal("chained"))
		})
	})
})
