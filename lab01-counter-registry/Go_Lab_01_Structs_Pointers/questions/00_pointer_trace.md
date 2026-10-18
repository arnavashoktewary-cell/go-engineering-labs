# Warm-up 00 — Predict the output (no running yet)

The exact code from our conversation:

```go
package main

import "fmt"

type Counter struct {
    Name  string
    Value int
}

func (c Counter) IncrementA() {
    c.Value++
}

func (c *Counter) IncrementB() {
    c.Value++
}

func main() {
    a := Counter{
        Name:  "Requests",
        Value: 10,
    }

    b := a

    a.IncrementA()
    b.IncrementB()

    p := &a
    p.IncrementB()

    fmt.Println(a.Value, b.Value)
}
```

## Your task

**Without compiling or running it, predict the exact two integers printed by `fmt.Println`.**

Trace the state after each statement:

| Statement | `a.Value` | `b.Value` | Why? |
|---|---:|---:|---|
| `a := Counter{..., Value: 10}` | ? | (not created) | |
| `b := a` | ? | ? | |
| `a.IncrementA()` | ? | ? | |
| `b.IncrementB()` | ? | ? | |
| `p := &a; p.IncrementB()` | ? | ? | |

Think about these three questions:

1. Does `b := a` alias the original struct or copy its fields?
2. Does a value receiver modify the original value or its copy?
3. Does `p := &a` point to `a` or `b`?

*The answer is intentionally not included.*
