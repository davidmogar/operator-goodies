# Operator goodies

This repository contains code that can be used in other operators to make some operations easier.

## Goodies

### Controller-runtime predicates

The following predicates are available:

| Name                         | Description                                                                      |
|------------------------------|----------------------------------------------------------------------------------|
| GenerationUnchangedPredicate | Skip update events that have a change in the object's metadata.generation field. |

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
