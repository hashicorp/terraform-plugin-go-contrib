[![PkgGoDev](https://pkg.go.dev/badge/github.com/hashicorp/terraform-plugin-go-patterns)](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go-patterns)

# terraform-plugin-go-patterns

terraform-plugin-go-patterns provides a set of common patterns that can be used
with [terraform-plugin-go](https://github.com/hashicorp/terraform-plugin-go).
It is a loose collection of common code that can be used when building
providers, but is not necessary for building providers with
terraform-plugin-go, and not every pattern is useful, applicable, or
appropriate for every provider.

## Status

terraform-plugin-go-patterns is a [Go
module](https://github.com/golang/go/wiki/Modules) versioned using [semantic
versioning](https://semver.org).

The module is currently on a v0 major version, indicating our lack of
confidence in the stability of its exported API. Developers depending on it
should do so with an explicit understanding that the API may change and shift
until we hit v1.0.0, as we learn more about the needs and expectations of
developers working with the module.

We are confident in the correctness of the code and it is safe to build on, so
long as the developer understands that the API may change in backwards
incompatible ways and they are expected to be tracking these changes.

## Documentation

Documentation is a work in progress. The GoDoc for packages, types, functions,
and methods should have complete information, but we're working to add
documentation about the use cases, semantics, and intricacies of each pattern.

Please bear with us as we work to get this information published, and please
[open
issues](https://github.com/hashicorp/terraform-plugin-go/issues/new/choose)
with requests for the kind of documentation you would find useful.

## Scope

This module is attempting to gather a loose collection of patterns. It is not
meant to encompass every pattern people may use with terraform-plugin-go, but
rather a subset of patterns that we feel are useful to a significantly large
enough percentage of the community. Patterns must have enough utility to
justify the cost of their ongoing maintenance and the additional cost to
discovery all patterns will have when a new one is added.

## Contributing

Please see [`.github/CONTRIBUTING.md`](https://github.com/hashicorp/terraform-plugin-go-patterns/blob/main/CONTRIBUTING.md).

## License

This module is licensed under the [Mozilla Public License v2.0](https://github.com/hashicorp/terraform-plugin-go-patterns/blob/main/LICENSE).
