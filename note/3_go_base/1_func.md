### 调用惯例

```go
package main

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	myFunction(66, 77)
}
```

![1_go_func_call.png](img/1_go_func_call.png)