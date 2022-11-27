/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtilTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Utilities Suite")
}

var _ = Describe("Test utilities", func() {
	Context("When GetRelativeDependencyPath is called", func() {
		It("returns the module path of a given dependency", func() {
			moduleRelativePath, err := GetRelativeDependencyPath("ginkgo")
			Expect(err).NotTo(HaveOccurred())
			Expect(moduleRelativePath).To(ContainSubstring("ginkgo@v"))
		})

		It("returns the module path of a given indirect dependency", func() {
			moduleRelativePath, err := GetRelativeDependencyPath("yaml.v3")
			Expect(err).NotTo(HaveOccurred())
			Expect(moduleRelativePath).To(ContainSubstring("yaml.v3@v"))
		})

		It("returns an empty string when the given dependency is not found", func() {
			moduleRelativePath, err := GetRelativeDependencyPath("nonexistent")
			Expect(err).To(HaveOccurred())
			Expect(moduleRelativePath).To(BeEmpty())
		})
	})

	Context("When FindGoModPath is called", func() {
		It("returns the filepath of the go.mod file", func() {
			currentWorkDir, _ := os.Getwd()
			goModFilepath, err := FindGoModPath(currentWorkDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(goModFilepath).To(ContainSubstring("go.mod"))
		})

		It("returns an empty string and an error when go.mod is not found", func() {
			goModFilepath, err := FindGoModPath("/var/tmp")
			Expect(goModFilepath).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("go.mod file not found"))
		})

		It("return an empty string and an error when the given directory does not exist", func() {
			goModFilepath, err := FindGoModPath("/nonexistent/dir")
			Expect(goModFilepath).To(BeEmpty())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("go.mod file not found"))
		})
	})
})
