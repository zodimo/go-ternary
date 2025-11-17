# go-ternary

A Go package providing ternary operator functionality with support for lazy evaluation.

## Features

- **Ternary**: A simple ternary operator function for conditional value selection
- **TernaryLazy**: A lazy-evaluated ternary operator that only evaluates the selected branch

## Installation

```bash
go get github.com/zodimo/go-ternary
```

## Usage

### Basic Ternary

```go
import "github.com/zodimo/go-ternary"

result := ternary.Ternary(x > 0, "positive", "negative")
```

### Lazy Ternary

```go
import (
    "github.com/zodimo/go-ternary"
    "github.com/zodimo/go-lazy"
)

result := ternary.TernaryLazy(
    lazy.NewLazy(func() bool { return x > 0 }),
    lazy.NewLazy(func() string { return expensiveOperation1() }),
    lazy.NewLazy(func() string { return expensiveOperation2() }),
)
// Only the selected branch will be evaluated
value := result.Get()
```

## Future Plans

- **Optional Memoization**: Add support for optional memoization to cache results of ternary operations, allowing for efficient repeated evaluations with the same inputs.

## License

MIT License

Copyright (c) 2024

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

