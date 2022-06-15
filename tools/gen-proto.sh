#!/bin/sh

bazel build $(bazel query 'kind(go_proto_library, //...)') && \
find . -name '*.pb.go' -delete && \
find bazel-bin/pkg/proto -name '*.pb.go' | while read f; do
	cat $f > $(echo $f | sed -e 's|.*/pkg/proto/|pkg/proto/|')
done