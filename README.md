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
