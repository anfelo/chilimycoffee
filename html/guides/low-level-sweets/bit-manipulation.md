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
