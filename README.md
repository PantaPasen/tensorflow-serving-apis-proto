# tensorflow-serving-apis-proto

A repository to store tensorflow serving APIs protobuf files.

If you are using go this module can now be installed using `go get github.com/pantapasen/tensorflow-serving-apis-proto`.

This is the recommended approach as it's compatible with both Bazel (via
Gazelle) and standard Go tooling.


## Bazel usage

You can also use Bazel to autogenerate the golang (or other languages) libraries on demand.
### WORKSPACE

```starlark
http_archive(
    name = "com_github_pantapasen_tfs_proto",
    strip_prefix = "tensorflow-serving-apis-proto-2.16.1",
    url = "https://github.com/pantapasen/tensorflow-serving-apis-proto/archive/2.16.1.tar.gz",
)
```

### BUILD file

```starlark
# golang example

# tensorflow serving apis
go_proto_library(
    name = "tensorflow_serving_apis_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "github.com/tensorflow/serving/tensorflow_serving/apis",
    proto = "@com_github_pantapasen_tfs_proto//:tensorflow_serving_apis_proto",
    deps = [
        ":tensorflow_core_example_go_proto",
        ":tensorflow_core_framework_go_proto",
        ":tensorflow_core_protobuf_go_proto",
        ":tensorflow_serving_config_go_proto",
    ],
)

go_proto_library(
    name = "tensorflow_serving_config_go_proto",
    importpath = "github.com/tensorflow/serving/tensorflow_serving/config",
    proto = "@com_github_pantapasen_tfs_proto//:tensorflow_serving_config_proto",
)

# tensorflow protos
go_proto_library(
    name = "tensorflow_core_framework_go_proto",
    importpath = "github.com/tensorflow/tensorflow/tensorflow/go/core/framework",
    proto = "@com_github_pantapasen_tfs_proto//:tensorflow_core_framework_proto",
)

go_proto_library(
    name = "tensorflow_core_example_go_proto",
    importpath = "github.com/tensorflow/tensorflow/tensorflow/go/core/example",
    proto = "@com_github_pantapasen_tfs_proto//:tensorflow_core_example_proto",
)

go_proto_library(
    name = "tensorflow_core_protobuf_go_proto",
    importpath = "github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf",
    proto = "@com_github_pantapasen_tfs_proto//:tensorflow_core_protobuf_proto",
    deps = [
        ":tensorflow_core_framework_go_proto",
    ],
)
```
