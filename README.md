# GoTools
> Golang编写的工具工具库，包括Base64等使用小工具。
## Base64
> 使用方法
- 编码<br>
```go
base64 := utils.Base64{}
result := base64.EncodeStr("hello")
```
- 解码<br>
```go
base64 := utils.Base64{}
result, err = base64.DecodeStr("aGVsbG8=")
```
