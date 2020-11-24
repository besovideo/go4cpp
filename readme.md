

#### golang调用dll/so(跨平台)


```
package main

import (
	"fmt"
	"github.com/besovideo/go4cpp"
	"log"
	"time"
)

func main() {
	go4cpp.InitLibrary(func(err error, data []byte) {
		log.Printf("=== %v %v", err, string(data))
	})

	var data = fmt.Sprintf("%v", time.Now().Format(time.RFC3339))
	go4cpp.Command([]byte(data), func(err error, data []byte) {
		log.Printf("+++ %v %v", err, string(data))
	})

	go4cpp.ReleaseLibrary()
}

```


说明:
1. dll/so 实现 include/go4c.h接口(参考lib/go4c.cpp)  
2. 将dll/so复制到可执行程序目录(名称更改为go4c.dll/libgo4c.so)  

```
export LD_LIBRARY_PATH=.
```


输出
```
$ go build -o hello.exe
$ ./hello.exe
hello world
2020/10/30 10:25:52 === hell
hello cmd: 2020-10-30T10:25:52+08:00
2020/10/30 10:25:52 2020-10-30T10:25:52+08:00
2020/10/30 10:25:52 +++2020-10-30T10:25:52+08:00
2020/10/30 10:25:52 === 2020-10-30T10:25:52+08:00
```
