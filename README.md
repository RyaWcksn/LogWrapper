# Package lists

## Logger  

The log package, it's wrap the [Zap](https://github.com/uber/zap) package and provides interface for logging in our microservices  

### How to integrate?  

1. Include the module in the project  
2. Create a new instance of the logger with following code below  

```go
package main

import (
	"github.com/RyaWcksn/logging/logger"
)

func main(){
	serviceName := "dev"
	env := "dev"
	logLevel := "debug"
	
	log := logger.New(serviceName, env, logLevel)
	
	log.Debug("Hello from log!")
}
```

Supported log level hierarchy are :  
- debug
- info
- warn
- error
- panic
- fatal

### Methods  

This logger package contains with some following methods and the implementation :

For `Debug` method, if it's using `development` or `stage` env, the log will be debug, if else, the log will be info level
_Debug_
```go
log.Debug("This is debug level")
```

_Info_
```go
log.Info("This is info level")
```

_Warn_
```go
log.Warn("This is warn level")
```

_Error_
```go
log.Error("This is error level")
```

_Panic_
```go
log.Panic("This is panic level")
```

_Fatal_
```go
log.Fatal("This is fatal level")
```

For `Debugf` method, if it's using `development` or `stage` env, the log will be debug, if else, the log will be info level
_Debugf_
```go
log.Debugf("This is %s level", "debugf")
```

_Infof_
```go
log.Infof("This is %s level", "infof")
```

_Warnf_
```go
log.Warnf("This is %s level", "warnf")
```

_Errorf_
```go
log.Errorf("This is %s level", "errorf")
```

_Panicf_
```go
log.Panic("This is %s level", "panicf")
```

_Fatalf_
```go
log.Fatalf("This is %s level", "fatalf")
```

