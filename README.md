# Go 类型安全容器库

## 使用示例

### 1. 类型安全的 `sync.Map`

```go
package main

import (
    "fmt"
    ts "github.com/button-chen/typesafecontainer-go/typesafecontainer"
)

func main() {
    // 创建一个类型安全的TMap(封装sync.Map)
    ssm := ts.TMap[string, string]{}

    ssm.Store("foo", "bar")
    fmt.Println(ssm.Load("foo"))
}
