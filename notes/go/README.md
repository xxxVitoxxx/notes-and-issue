# interface

用 `type` 宣告介面，介面可以定義要實作介面的方法

```go
// 宣告介面
type Person interface{
    // 要實作該介面的方法
    GetName()string
    SetName(name string)
}
```  

## 隱性實作

go 的 interface 是隱性實作的，不需要另外宣告實作了一個介面，只要實作了某介面的所有方法，就等於你實作了該介面  

```go
type Person interface{
    GetName()string
    SetName(name string)
}

// 實作了 Person 介面
type Man struct {
    Name string
}

func(m *Man)GetName() string {
    return m.Name
}

func(m *Man)SetName(name string) {
    m.Name = name
} 
```