#!/bin/bash

set -eu

gopath_is_defined=$GOPATH

go test -bench=. radix \
  | grep "^Benchmark" \
  | tr '_' ' ' \
  | awk '{print $2, $4, $3, $6}' \
  | sort -nk 4 \
  | sort -snrk 2 \
  | sort -sk 1,1 \
  | tr ' ' '\t'
