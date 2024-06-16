# errorg

## Introduction

errog helps to error handling for golang apps. 


## Usage

```shell
go get -u github.com/cemayan/errorg
```

In order to initialize errog you can do like below:

```go
func init() {
	errorg.Init(errorg.JsonLog())
}
```

After that when do you want return an error you can do like below:


```go
func test() error {
	return errorg.New("http", "401")
}
```



> [!NOTE]  
> "http" => group | "401"  => code of group (For this ex it means http **401** error code.)
