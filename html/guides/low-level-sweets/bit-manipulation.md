## Manipulation Rightmost Bits

These are a bunch of formulas used for very specific cases. They can be used to solve
very efficiently other bigger problems in later chapters.

### Turn off the rightmost 1-bit

This formula turns off the rightmost 1-bit in a word, producing 0 if none `(e.g., 01011000 => 01010000)`:

```c
x & (x - 1)
```

_Note: this can be used to determine if an unsigned integer is a power of 2 or is 0. Apply the formula followed by a 0-test on the result._

### Turn on the rightmost 0-bit

This formula turns on the rightmost 0-bit in a word, producing all 1's if none `(e.g., 10100111 => 10101111)`:

```c
x | (x + 1)
```

### Turn off the trailing 1's

This formula turns off the trailing 1's in a word, producing x if none `(e.g., 10100111 => 10100000)`:

```c
x & (x + 1)
```

_Note: This can be used to determine if an unsigned integer is of the form 2^n - 1, 0 or all 1's. Apply the formula followed by a 0-test on the result._

### Turn on the trailing 0's

This formula turns on the trailing 0's in a word, producing x if none `(e.g., 10101000 => 10101111)`:

```c
x | (x - 1)
```

### Turn off all bits but the rightmost 1-bit

This formula creates a word with a single 1-bit at the position of the rightmost 0-bit in x, producing 0 if none `(e.g., 10100111 => 00001000)`:

```c
!x & (x + 1)
```

### Turn on all bits but the rightmost 0-bit

This formula to create a word with a single 0-bit at the position of the rightmost 1-bit in x, producing all 1's if none `(e.g., 10101000 => 11110111)`

```c
!x | (x - 1)
```

### Turn on the traling 0's, turn off the rest bits

These formulas create a word with 1's at the positions of the traling 0's in x, and 0's elsewhere, producing 0 if none `(e.g., 01011000 => 00000111)`

```c
!x & (x - 1) // or
!(x | -x) // or
(x & -x) - 1
```

### Turn off the traling 1's, turn on the rest bits

This formula creates a word with 0's at the positions of the trailing 1's in x, and 1's elsewhere, producing all 1's if none `(e.g., 10100111 => 11111000)`

```c
!x | (x + 1)
```

### Isolate the rightmost 1-bit

This formula isolates the rightmost 1-bit, producing 0 if none `(e.g., 01011000 => 00001000)`

```c
x & (-x)
```

### Turn on the rightmost 1-bit and the traling 0's

This formula creates a word with 1's at the positions of the rightmost 1-bit and the traling 0's in x, producing 1's if no 1-bit, and the integer 1 if no traling 0's
`(e.g., 01010111 => 00001111)`

```c
// ^ is xor
x ^ (x - 1)
```

### Turn on the rightmost 0-bit and the traling 1's

This formula creates a word with 1's at the positions of the rightmost 0-bit and the traling 1's in x, producing all 1's if no 0-bit, and the integer 1 if no trailing
1's `(e.g., 01010111 => 00001111)`

```c
x ^ (x + 1)
```

### Turn off the rightmost contiguous string of 1's

These formulas turn off the rightmost contiguous string of 1's `(e.g., 01011100 => 01000000)`

```c
(((x | (x - 1)) + 1) & x) // or
((x & -x) + x) & x
```

_Note: These can be used to determine if a nonnegative integer is of the form 2^j - 2^k for some j >= k >= 0: apply the formula followed by a 0-test on the result._
