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


## Performances

In general I have found that for 32bits ints LSD radix sort is faster on my
machine (MBP 2.5 GHz Intel Core i7 with 16GB memory).
The code structure is simpler and has only 5 array iterations.

For 64bits ints LSD radix sort becomes approximately twice as slow due to the 4
additional digits, while MSD radix sort have almost identical performances.

Both LSD and MSD have runtime roughly linear to the size of the array,
as expected.
They perform quite a lot better than standard sort for a wide range of sizes:

| type   | array size | algorithm    | ns/op    |
| ---    | ---        | ---          | ---      |
| Int    | 100000     | Radix        |  1984536 |
| Int    | 100000     | StandardSort | 21836864 |
| Int    | 10000      | Radix        |   250743 |
| Int    | 10000      | StandardSort |  1701220 |
| Int    | 1000       | Radix        |    14731 |
| Int    | 1000       | StandardSort |   129036 |
| Int    | 100        | Radix        |     1845 |
| Int    | 100        | StandardSort |     8540 |
| Uint   | 100000     | Radix        |  1869603 |
| Uint   | 100000     | StandardSort | 21537914 |
| Uint   | 10000      | Radix        |   243012 |
| Uint   | 10000      | StandardSort |  1702709 |
| Uint   | 1000       | Radix        |    14701 |
| Uint   | 1000       | StandardSort |   131914 |
| Uint   | 100        | Radix        |     1799 |
| Uint   | 100        | StandardSort |     8543 |
| Int64  | 100000     | RadixMSD     |  1881905 |
| Int64  | 100000     | RadixLSD     |  3304843 |
| Int64  | 100000     | StandardSort | 22708883 |
| Int64  | 10000      | RadixMSD     |   253543 |
| Int64  | 10000      | RadixLSD     |   305071 |
| Int64  | 10000      | StandardSort |  1818001 |
| Int64  | 1000       | RadixMSD     |    14674 |
| Int64  | 1000       | RadixLSD     |    33201 |
| Int64  | 1000       | StandardSort |   131369 |
| Int64  | 100        | RadixMSD     |     2084 |
| Int64  | 100        | Insertion    |     3727 |
| Int64  | 100        | RadixLSD     |     5379 |
| Int64  | 100        | StandardSort |     8551 |
| Uint64 | 100000     | RadixMSD     |  1940448 |
| Uint64 | 100000     | RadixLSD     |  3210152 |
| Uint64 | 100000     | StandardSort | 23593543 |
| Uint64 | 10000      | RadixMSD     |   252472 |
| Uint64 | 10000      | RadixLSD     |   305155 |
| Uint64 | 10000      | StandardSort |  1893974 |
| Uint64 | 1000       | RadixMSD     |    14994 |
| Uint64 | 1000       | RadixLSD     |    33648 |
| Uint64 | 1000       | StandardSort |   141852 |
| Uint64 | 100        | RadixMSD     |     1795 |
| Uint64 | 100        | Insertion    |     3703 |
| Uint64 | 100        | RadixLSD     |     5580 |
| Uint64 | 100        | StandardSort |     9260 |
| Int32  | 100000     | RadixLSD     |  1632580 |
| Int32  | 100000     | RadixMSD     |  1848354 |
| Int32  | 100000     | StandardSort | 21300107 |
| Int32  | 10000      | RadixLSD     |   161597 |
| Int32  | 10000      | RadixMSD     |   279245 |
| Int32  | 10000      | StandardSort |  1731624 |
| Int32  | 1000       | RadixMSD     |    14019 |
| Int32  | 1000       | RadixLSD     |    17815 |
| Int32  | 1000       | StandardSort |   128745 |
| Int32  | 100        | RadixMSD     |     1905 |
| Int32  | 100        | RadixLSD     |     2924 |
| Int32  | 100        | Insertion    |     4219 |
| Int32  | 100        | StandardSort |     8579 |
| Uint32 | 100000     | RadixLSD     |  1624866 |
| Uint32 | 100000     | RadixMSD     |  1869044 |
| Uint32 | 100000     | StandardSort | 21341060 |
| Uint32 | 10000      | RadixLSD     |   156249 |
| Uint32 | 10000      | RadixMSD     |   267186 |
| Uint32 | 10000      | StandardSort |  1724809 |
| Uint32 | 1000       | RadixMSD     |    14411 |
| Uint32 | 1000       | RadixLSD     |    16775 |
| Uint32 | 1000       | StandardSort |   128274 |
| Uint32 | 100        | RadixMSD     |     1756 |
| Uint32 | 100        | RadixLSD     |     2817 |
| Uint32 | 100        | Insertion    |     4102 |
| Uint32 | 100        | StandardSort |     8346 |


## How to

| Command            | Description                             |
| ---                | ---                                     |
| go build radixsort | build the package                       |
| go test radixsort  | run test                                |
| ./bench.sh         | run benchmarks and pretty print results |


## How it works


 * __LSD:__
   1. all existing radix digits are counted for all bytes of all elements in
      the array, grouped by byte position, over a single array iteration.
   2. The radix count table is turned into an indice offsets table by summation.
   3. Finally for all digit positions from LSD to MSD, the elements in the input
      array are swapped according to the indice offsets table.

 * __MSD:__ the algorithm works recursively one digit position at a time,
   starting from most significant digit.
   It essentially follows the same 3 steps as LSD for that position.
   After swapping the radix count table is scanned for buckets with more than 1
   element and the algorithm recursively call itself for these bucket, moving
   the digit position by one. When the bucket is sufficiently small, insertion
   sort is used instead.

 * For small arrays (currently size 64 or less), both LSD and MSD sorts
   delegates to insertion sort directly for better performance.

 * Both LSD and MSD uses swap space equal to the size of the input array.

 * Both LSD and MSD sorts have a int32 and int64 versions.
   Other sort methods for uint32, uint64, int and uint delegates to these two
   base versions by casting slice pointers.
