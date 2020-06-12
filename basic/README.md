# Basic algorithms

## Closest value in Binary Search Tree

Find the closest value in Binary Search Tree

### Benchmarks

| Name | Quantity | ns/op | B/op | allocs/op |
|--|--|--|--|--|
| BenchmarkClosestValueBstRecursive/n=5-12 | 216392253 | 5.46 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstRecursive/n=3-12 | 120437488 | 9.98 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstRecursive/n=12-12 | 21364524 | 9.86 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstRecursive/n=18-12 | 21631774 | 9.76 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstIterative/n=5-12 | 268988937 | 4.45 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstIterative/n=3-12 | 162285542 | 7.44 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstIterative/n=12-12 | 60098214 | 7.53 ns/op | 0 B/op | 0 allocs/op
| BenchmarkClosestValueBstIterative/n=18-12 | 60784563 | 7.42 ns/op | 0 B/op | 0 allocs/op

## Fibonacci

Returns the nth Fibonacci number

### Benchmarks

| Name | Quantity | ns/op | B/op | allocs/op |
|--|--|--|--|--|
| BenchmarkFibonacciRecursive/n=6-12 | 30080389  | 39.5 ns/op | 0 B/op | 0 allocs/op |
| BenchmarkFibonacciRecursive/n=9-12 | 6914959  | 172 ns/op | 0 B/op | 0 allocs/op |
| BenchmarkFibonacciRecursive/n=14-12 | 633278  | 1887 ns/op | 0 B/op | 0 allocs/op |
| BenchmarkFibonacciRecursive/n=20-12 | 35181  | 33943 ns/op | 0 B/op | 0 allocs/op |
| BenchmarkFibonacciDynamic/n=6-12 | 34238073  | 33.4 ns/op | 64 B/op | 1 allocs/op |
| BenchmarkFibonacciDynamic/n=9-12 | 29981410  | 41.9 ns/op | 96 B/op | 1 allocs/op |
| BenchmarkFibonacciDynamic/n=14-12 | 23144442  | 51.3 ns/op | 128 B/op | 1 allocs/op |
| BenchmarkFibonacciDynamic/n=20-12 | 18792841  | 66.6 ns/op | 176 B/op | 1 allocs/op |
| BenchmarkFibonacciSpace/n=6-12 | 35388938  | 34.3 ns/op | 64 B/op | 1 allocs/op |
| BenchmarkFibonacciSpace/n=9-12 | 29324437  | 42.7 ns/op | 96 B/op | 1 allocs/op |
| BenchmarkFibonacciSpace/n=14-12 | 23439559  | 52.3 ns/op | 128 B/op | 1 allocs/op |
| BenchmarkFibonacciSpace/n=20-12 | 17954664  | 67.2 ns/op | 176 B/op | 1 allocs/op |

## Numbers Sum

Returns two numbers from an array where their sum is equal to the parameter sum

### Benchmarks

| Name | Quantity | ns/op | B/op | allocs/op |
|--|--|--|--|--|
| BenchmarkTwoNumbersSumFor-12 | 54696116  | 46.3 ns/op | 16 B/op | 1 allocs/op |
| BenchmarkTwoNumbersSumMap-12 | 13092234  | 143 ns/op | 16 B/op | 1 allocs/op |
| BenchmarkTwoNumbersSumShrinking-12 | 2445675  | 453 ns/op | 16 B/op | 1 allocs/op |

## Validate Subsequence

Returns a boolean value indicating if one list of numbers is a subsequence of another list of numbers

### Benchmarks

| Name | Quantity | ns/op | B/op | allocs/op |
|--|--|--|--|--|
| BenchmarkValidateSubsequence-12 | 247067107  | 4.74 ns/op | 0 B/op | 0 allocs/op |
