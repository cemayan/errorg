# errorg

## Introduction

errog helps to error handling for golang apps and  thanks to golang/plugins it can be added the new plugin

The main purpose of this library is to manage all error codes from a common library.


## Example

In order to see on example app make commands can be used like below:

```shell
make init
make example
```

And example app returns like below:

Ex App Group is  "http" and  code is "401" 

```json
{"msg":"Although the HTTP standard specifies \"unauthorized\", semantically this response means \"unauthenticated\". That is, the client must authenticate itself to get the requested response."}
```



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


## Plugins
You can create a plugin for errorg. Errorg uses Golang plugins and some plugins created by default. If you want to further information for plugins check [this](https://pkg.go.dev/plugin) out

You can add the new plugin to under the **plugins** folder.
Example plugin:

```go
package  main   

type httpError struct {
}

func (g httpError) GetErrorMap() map[string]interface{} {
	return map[string]interface{}{
		"400": struct {
			Message string `json:"msg"`
		}{"The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)."},
	}
}

var Group httpError
```

Group represent the interface that will be implemented.

After that you need to add new make command to Makefile

Ex command:

```makefile
http-plugin: # To build a  http errors plugin.
	@echo "  >  Building http-plugin..."
	chmod -R   700  ${CURDIR}/${PLUGINS_FOLDER}
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build  -buildmode=plugin   -o  ${CURDIR}/${PLUGINS_FOLDER}/http/${PLUGIN_NAME} ${PLUGINS_FOLDER}/http/${PLUGIN_GO_FILE}
```

plugin module waits *.so* files Therefore this command will be created the **plugin.so** file. 

Lastly, you need to add new plugin infos to configs/config.yaml 

```yaml 
plugins:
  - name: foo
    location: plugins/foo/plugin.so
```