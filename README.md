# A golang monorepo using [bazel](https://bazel.build/)

[bazel-gazelle](https://github.com/bazelbuild/bazel-gazelle) is used to generate BUILD.bazel files.

Check [WORKSPACE](./WORKSPACE) and [BUILD](./BUILD)

## Creating docker images

Run `bazel run //cmd/service1:image`

See the `container_image` rule in [/cmd/service1/BUILD.bazel](./cmd/service1/BUILD.bazel)
