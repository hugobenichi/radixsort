## What is it

The radixsort package offers sort implementations for various Go built-in
numeric types based on radix sort.
A most significant digit (MSD) and least significant digit (LSD) radix sort are
offered for every numeric type, using names such as xxxMSD or xxxLSD, where xxx
is the type name.
For every type a wrapper to the fastest implementation also exists. At the
moment for 32bits types the wrapper delegates to LSD radix sort, and for 64bits
types the wrapper delegates to MSD radix sort.

In practice simply import the package and use

```
radixsort.Int32(array) // replace Int32 with your favorite numeric type
```

Supported types are int, uint, int32, uint32, int64, uint64.
The package has no external dependency.

## How to

| Command | Description |
| --- | --- |
| go build radixsort | build the package |
| go test radixsort  | run test |
| ./bench.sh         | run benchmarks and pretty print results |

## How it works

TODO: explain code structure
TODO: explain code reuse via pointer casting
TODO: talk about performance
