# count-repeat
## count repeat item in slice. not contain self item.

## Install
```shell
go get github.com/mazezen/count-repeat
```

## Example
```go
fmt.Println(Count([]interface{}{"10", "20", "600", "20", "30"}))
fmt.Println(Count([]interface{}{"20", "20", "600", "20", "30"}))
fmt.Println(Count([]interface{}{"600", "20", "600", "20", "30"}))
fmt.Println(Count([]interface{}{"20", "20", "600", "20", "30"}))
fmt.Println(Count([]interface{}{"80", "20", "600", "20", "30"}))
fmt.Println(Count([]interface{}{"100", "20", "600", "20", "30"}))
```

```shell
1
2
1
2
1
1
--- PASS: TestCount (0.00s)
```
