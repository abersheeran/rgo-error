# (R)Go Error

Like Rust's `Result` type, but for Go.

## Usage

```go
package main

import (
    . "github.com/abersheeran/rgo-error"
)

func normal_error() (string, error) {
    value, err := redis.Get(ctx, "key").Result() 
    if err != nil {
        return "", err
    }
    return value, nil
}

// Use `Result[T]` replace (T, error)

func use_result() Result[string] {
    return AsResult[string](redis.Get(ctx, "key").Result())
}
```

`Result` method:

- `func (r Result[T]) IsOk() bool`
- `func (r Result[T]) IsErr() bool`
- `func (r Result[T]) Error() error`
- `func (o Result[T]) Unwrap() T`: Returns a valid value; panic if err is present

You can use `Ok` and `Err` in your function replace `AsResult`

```python
func ok_or_err() Result[string] {
    return Ok[string]("hello world")
    // return Err[string](error)
}

func ok_or_err_use_res() (res Result[string]) {
    return res.Ok("hello world")
    // return res.Err(error)
}
```
