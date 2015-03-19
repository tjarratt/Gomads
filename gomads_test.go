package gomads_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/gomads"
)

var _ = Describe("Monads", func() {
	Describe("Maybe", func() {
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

		It("treats funcs that panic as nil", func() {
			m := Maybe(func() interface{} {
				panic("woah there")
			})

			Expect(m.Value()).To(BeNil())
		})
	})

	Describe("chained errors", func() {
		It("can be used to chain operations that may fail", func() {
			e := Error(func(thunk interface{}) interface{} {
				return 5
			}).Compose(Error(func(thunk interface{}) interface{} {
				return 1
			}))

			Expect(e.Value()).To(Equal(1))
		})

		It("returns nil if any of the operations panics", func() {
			e := Error(func(_ interface{}) interface{} {
				return 5
			}).Compose(Error(func(_ interface{}) interface{} {
				panic("FFFFFFFFFFFFFFFFUUUUUUUUUUUUUUU")
			}))

			Expect(e.Value()).To(BeNil())
		})

		It("chains the result of the computation along through each func", func() {
			e := Error(func(_ interface{}) interface{} {
				return 1
			}).Compose(Error(func(thunk interface{}) interface{} {
				return thunk.(int) + 1
			}))

			Expect(e.Value()).To(Equal(2))
		})
	})
})
