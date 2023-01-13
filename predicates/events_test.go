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

package predicates

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

var _ = Describe("Predicates", Ordered, func() {

	Context("When using the IgnoreAllPredicate", func() {
		instance := IgnoreAllPredicate{}

		It("should ignore creating events", func() {
			contextEvent := event.CreateEvent{
				Object: &v1.Pod{},
			}
			Expect(instance.Create(contextEvent)).To(BeFalse())
		})

		It("should ignore deleting events", func() {
			contextEvent := event.DeleteEvent{
				Object: &v1.Pod{},
			}
			Expect(instance.Delete(contextEvent)).To(BeFalse())
		})

		It("should ignore generic events", func() {
			contextEvent := event.GenericEvent{
				Object: &v1.Pod{},
			}
			Expect(instance.Generic(contextEvent)).To(BeFalse())
		})

		It("should ignore update events", func() {
			contextEvent := event.UpdateEvent{
				ObjectNew: &v1.Pod{},
				ObjectOld: &v1.Pod{},
			}
			Expect(instance.Update(contextEvent)).To(BeFalse())
		})
	})

	Context("When using the NewObjectsPredicate", func() {
		instance := NewObjectsPredicate{}

		It("should allow creating events", func() {
			contextEvent := event.CreateEvent{
				Object: &v1.Pod{},
			}
			Expect(instance.Create(contextEvent)).To(BeTrue())
		})

		It("should ignore deleting events", func() {
			contextEvent := event.DeleteEvent{
				Object: &v1.Pod{},
			}
			Expect(instance.Delete(contextEvent)).To(BeFalse())
		})

		It("should ignore generic events", func() {
			contextEvent := event.GenericEvent{
				Object: &v1.Pod{},
			}
			Expect(instance.Generic(contextEvent)).To(BeFalse())
		})

		It("should ignore update events", func() {
			contextEvent := event.UpdateEvent{
				ObjectNew: &v1.Pod{},
				ObjectOld: &v1.Pod{},
			}
			Expect(instance.Update(contextEvent)).To(BeFalse())
		})
	})

})
