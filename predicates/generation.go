package predicates

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// GenerationUnchangedPredicate implements a default update predicate function on Generation unchanged.
//
// This predicate will skip update events that have a change in the object's metadata.generation field.
// The metadata.generation field of an object is incremented by the API server when writes are made to the spec field of
// an object. This allows a controller to ignore update events where the spec has unchanged, and only the metadata
// and/or status fields are changed.
type GenerationUnchangedPredicate struct {
	predicate.Funcs
}

// Update implements default UpdateEvent filter for validating generation change.
func (GenerationUnchangedPredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld == nil || e.ObjectNew == nil {
		return false
	}

	return e.ObjectNew.GetGeneration() == e.ObjectOld.GetGeneration()
}
