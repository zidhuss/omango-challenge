# OmanGo January 2023 Challenge

## K Numbers

There is a collection of natural numbers (we will call K numbers) which when squared can be can be split (assuming base 10) into two positive integers, such that their sum is equal to the original number.

For example, the number 703 is a K number given that 703² is 494209 which can be split into 494 and 209 and their sum gives again 703. Another example is 9 (9² = 81 and 8 + 1 = 9).

We have to keep in mind that both numbers of the decomposition don't have to contain the same amount of digits. For example, the number 2728, we have 2728² = 7441984 which is a K number because 744 + 1984 = 2728. There can also be a case where the square contains a 0. For example, with 4879 we have 4879² = 23804641 which is a K number because 238 + 04641 = 4879.

The first number of the split can be 0 (e.g. 1 is a K number), however, the second number can't be 0. Due to this, 100 is **not** a K number.

The signature of your function should be:

```go
func Solve(number int)  bool
```

You may implement other functions called by your `Solve` function if you wish.
You can start implementing the function in the `main.go` file.

## Testing

To test your solution you can run the unit tests:

```
go test
```

### Input

The input will be a positive integer between 1 and 65536.

### Output Spec

The output should be a boolean. `true` if it's a K number and false if not.

### Example Input & Output

| Input   | Output  |
| ------- | ------- |
| `22222` | `true`  |
| `75`    | `false` |
| `99`    | `true`  |
| `1`     | `true`  |
| `100`   | `false` |
