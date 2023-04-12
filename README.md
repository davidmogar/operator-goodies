# Operator goodies

This repository contains code that can be used in other operators to make some operations easier.

## Goodies

### Controller-runtime predicates

The following predicates are available:

| Name                                 | Description                                                                                                                                                     |
|--------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GenerationUnchangedPredicate         | Skip update events that have a change in the object's metadata.generation field.                                                                                |
| GenerationUnchangedOnUpdatePredicate | Skip any event except updates. In the case of update events that have a change in the object's metadata.generation field, those events will be skipped as well. |
| IgnoreAllPredicate                   | Skip any kind of event.                                                                                                                                         |                                                                                                                                        |
| NewObjectsPredicate                  | Skip all events but creations.                                                                                                                                  |                                                                                                                                  

This is an example of using the `GenerationUnchangedPredicate`:
```go
return ctrl.NewControllerManagedBy(manager).
    For(&v1.Pod{}, builder.WithPredicates(predicates.GenerationUnchangedPredicate{})).
    Complete(reconciler)
```

### Reconciling using a handler

By using the `reconciler.ReconcileHandler` it's possible to define an adapter with reconcile operations that would
reconcile a resource. The adapter would look like the following:
```go
type Adapter struct {
    resources *v1alpha1.Resource
    logger  logr.Logger
    client  client.Client
    context context.Context
}

func (a *Adapter) EnsureAnActionIsTaken() (results.OperationResult, error) {
    return results.ContinueProcessing()
}
```

Now, to use it in the reconcile loop:
```go
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("Release", req.NamespacedName)

	release := &v1alpha1.Release{}
	err := r.Get(ctx, req.NamespacedName, release)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	adapter := NewAdapter(release, logger, r.Client, ctx)

	operations := []operations.ReconcileOperation{
		adapter.EnsureAnActionIsTaken,
	}

	return r.ReconcileHandler(operations)
}
```

These operations will be executed in order and the result will be validated by the `reconciler.ReconcileHandler`.

A good example of an operator defining a reconcile adapter (and handler) is the [release-service](https://github.com/redhat-appstudio/release-service/tree/main/controllers/release).

### Referencing external CRDs when testing operators with EnvTest

When testing an operator with Ginkgo and EnvTest, CRDs YAMLs have to be specified in the Environment declaration.
For external CRDs, those YAMLs can be extracted from the go directory once all `go.mod` dependencies have been fetched.
To make this operation easier, the following functions are defined:

| Function                           | Description                                                                           |
|------------------------------------|---------------------------------------------------------------------------------------|
| GetRelativeDependencyPath          | Returns the path within the GOPATH of a given dependency.                             |
| GetRelativeDependencyPathWithError | Returns the path within the GOPATH of a given dependency and the error raised if any. |

The following example shows how to reference Tekton CRDs though the tektoncd/pipeline dependency defined in the `go.mod`
file.
```go
require (
	...
    github.com/tektoncd/pipeline v0.41.0
	...
)
```
```go
testEnv = &envtest.Environment{
    CRDDirectoryPaths: []string{
        filepath.Join(
            build.Default.GOPATH,
            "pkg", "mod", test.GetRelativeDependencyPath("tektoncd/pipeline"), "config",
        ),
    },
    ErrorIfCRDPathMissing: true,
}
```

### Declaring and using conditions

To declare conditions in the status of CRDs, two types are declared in the `conditions` package:

* **ConditionType:** Represents a Kubernetes condition type.
* **ConditionReason:** Represents the reason of a Kubernetes condition.

On top of that, two functions are supplied to set a new condition:

* `SetCondition`, which expects a reference to the `metav1.Condition` slice and the type, status and reason for the condition.
* `SetConditionWithMessage`, which on top of the above, will expect a message for the new condition.

An example use case would be the following:
```go
const (
    releasedType    conditions.ConditionType   = "Released"
    succeededReason conditions.ConditionReason = "Succeeded"
)

func (m *MyCRD) MarkReleased() {
    conditions.SetCondition(&m.Status.Conditions, releasedType, metav1.ConditionTrue, succeededReason)	
}
```
