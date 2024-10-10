### 4. 自定义错误

有时，内置的 `errors.New` 和 `fmt.Errorf` 不足以表示复杂的错误情况。在这种情况下，你可以定义自己的错误类型。

```go
package main

import (
    "fmt"
)

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{
        Code:    404,
        Message: "Resource not found",
    }
}

func main() {
    err := doSomething()
    if err != nil {
        fmt.Println(err)
    }
}
```

- **自定义错误类型**：`MyError` 类型实现了 `error` 接口，包含了一个 `Code` 和 `Message` 字段，可以包含更多信息。
- **`Error()` 方法**：实现 `Error()` 方法来返回错误的字符串表示。

### 5. 错误包装

Go 1.13 引入了错误包装，通过 `fmt.Errorf` 和 `errors` 包中的新函数来包装和解包错误。

#### 5.1 错误包装

使用 `%w` 占位符来包装错误：

```go
import (
    "fmt"
    "errors"
)

func originalError() error {
    return errors.New("original error")
}

func wrappedError() error {
    err := originalError()
    return fmt.Errorf("an error occurred: %w", err)
}

func main() {
    err := wrappedError()
    fmt.Println(err) // 打印: an error occurred: original error
}
```

#### 5.2 错误解包

`errors.Is` 和 `errors.As` 用于检查和解包错误。

- **`errors.Is`** 用于检查一个错误是否等于另一个错误：

```go
import (
    "fmt"
    "errors"
)

func main() {
    err1 := errors.New("error 1")
    err2 := fmt.Errorf("wrapped: %w", err1)

    if errors.Is(err2, err1) {
        fmt.Println("err2 contains err1")
    }
}
```

- **`errors.As`** 用于将错误解包为特定的错误类型：

```go
import (
    "fmt"
    "errors"
)

type MyError struct {
    Code int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("MyError with code %d", e.Code)
}

func main() {
    myErr := &MyError{Code: 404}
    err := fmt.Errorf("an error occurred: %w", myErr)

    var targetErr *MyError
    if errors.As(err, &targetErr) {
        fmt.Println("Error code:", targetErr.Code)
    }
}
```

### 6. 多个错误的处理

在复杂的系统中，有时需要处理多个错误，Go 的 `multierror` 库可以帮助你收集和处理多个错误。

```go
import (
    "fmt"
    "github.com/hashicorp/go-multierror"
)

func main() {
    var result *multierror.Error

    result = multierror.Append(result, errors.New("Error 1"))
    result = multierror.Append(result, errors.New("Error 2"))

    if result.ErrorOrNil() != nil {
        fmt.Println(result.Error()) // 打印: 2 errors occurred: ...
    }
}
```

### 7. 总结

- **`error` 接口** 是 Go 中表示错误的基本机制。
- **创建错误** 可以使用 `errors.New`、`fmt.Errorf` 或自定义错误类型。
- **错误处理** 通过检查返回的 `error` 来进行，保持代码简单且清晰。
- **错误包装** 和 **解包** 提供了处理和传递上下文信息的机制。
- **自定义错误类型** 允许更丰富的错误表达和处理。
- **多错误处理** 可用于处理多个错误场景。

Go 的错误处理哲学鼓励开发者显式地处理错误，避免复杂的异常机制，使得代码更易于理解和维护。
