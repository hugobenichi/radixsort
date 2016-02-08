#!/bin/bash

go test -bench=. radix \
  | grep "^Benchmark" \
  | tr '_' ' ' \
  | awk '{print $2, $4, $3, $6}' \
  | sort -nk 4 \  # sort by run time
  | sort -snk 2   # sort by input size
