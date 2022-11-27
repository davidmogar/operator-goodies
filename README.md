# Operator goodies

This repository contains code that can be used in other operators to make some operations easier.

## Testing

When testing an operator with Ginkgo and EnvTest, CRDs yamls have to be specified in the Environment declaration.
For external CRDs, those yamls can be extracted from the go directory once all go.mod dependencies have been fetched.
To make this operation easier, the following functions are defined:

| Function                           | Description                                               | Returns                                                  |
|------------------------------------|-----------------------------------------------------------|----------------------------------------------------------|
| GetRelativeDependencyPath          | Returns the path withing the GOPATH of a given dependency | Path of the dependency                                   |
| GetRelativeDependencyPathWithError | Returns the path withing the GOPATH of a given dependency | Path of the dependency and error if the operation failed |
