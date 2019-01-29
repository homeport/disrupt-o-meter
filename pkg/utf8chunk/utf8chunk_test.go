// Copyright © 2019 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package utf8chunk_test

import (
	"github.com/homeport/disrupt-o-meter/pkg/utf8chunk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UTF8 Chunk util test", func() {
	Context("Tests if the util methods work correctly", func() {
		It("should remove the first colour code", func() {
			code := "\033[31m"
			message := "test"

			m, c := utf8chunk.RemoveStartingColour(code + message)

			Expect(m).To(BeEquivalentTo(message))
			Expect(c).To(BeEquivalentTo(code))
		})

		It("should just returns the default", func() {
			message := "test"

			m, c := utf8chunk.RemoveStartingColour(message)

			Expect(m).To(BeEquivalentTo(message))
			Expect(c).To(BeEquivalentTo(""))
		})
	})
})