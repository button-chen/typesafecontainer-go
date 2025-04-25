# Go 类型安全容器库

## 简介
在 Go 中，标准库提供了多种容器类型，如 `list`、`ring`、`heap`、`sync.Pool` 和 `sync.Map`。然而，这些容器默认是类型不安全的，即它们可以接受任何类型的值，这可能导致运行时错误。为了提升代码的类型安全性和可维护性，封装了这些容器，使其支持泛型，从而确保类型安全。

## 类型安全的优点
类型安全的容器具有以下显著优点：
1. **减少运行时错误**：通过在编译时检查类型，避免了因类型不匹配导致的运行时错误。
2. **提高代码可读性**：明确指定容器中存储的类型，使代码更易于理解和维护。
3. **增强代码健壮性**：类型安全的容器可以更好地抵御意外的类型错误，减少潜在的 bug。
4. **可提升开发效率**  ：编译器会帮助开发者发现类型相关的问题，减少调试时间。

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
	ssm := TMap[string, string]{}

	ssm.Store("foo", "bar")
	fmt.Println(ssm.Load("foo"))
}
