package go4cpp

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lgo4c

#include "go4c.h"
#include <stdlib.h>

void FnCallBackLibGO(char* data, int32_t len);
void FnCallBackCmdGO(int32_t cmdId, char* data, int32_t len);
*/
import "C"
import (
	"log"
	"math"
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"
)

//export FnCallBackLibGO
func FnCallBackLibGO(data *C.char, len C.int32_t) {
	//var s = C.GoStringN(data, len)
	var s []byte = C.GoBytes(unsafe.Pointer(data), len)

	defaultCallback(s)
}

//export FnCallBackCmdGO
func FnCallBackCmdGO(cmdId C.int32_t, data *C.char, len C.int32_t) {
	var s []byte = C.GoBytes(unsafe.Pointer(data), len)

	log.Println(string(s))

	var iCmdId = int32(cmdId)
	if fun, ok := mapCmdFun.Load(iCmdId); ok {
		if cb, success := fun.(FunCallBackNormal); success {
			cb(s)
		}
		mapCmdFun.Delete(iCmdId)
	}
}

// InitLibrary 初始化
func InitLibrary(fun FunCallBackNormal) int {
	defaultCallback = fun

	var rc C.int32_t = C.Go4CInit_C(C.FnCallBackLib_C(C.FnCallBackLibGO))

	return int(rc)
}

// ReleaseLibrary 释放库
func ReleaseLibrary() {
	C.Go4CRelease_C()
}

// Command 调用动态库函数
func Command(data []byte, fun FunCallBackNormal) int {
	var cmdId = getCmdId()
	mapCmdFun.Store(cmdId, fun)

	var rc C.int32_t = C.Go4CInitCommand_C(
		(*C.char)(unsafe.Pointer(
			(*reflect.StringHeader)(unsafe.Pointer(&data)).Data)),
		C.int32_t(len(data)),
		C.FnCallBackCmd_C(C.FnCallBackCmdGO),
		C.int32_t(cmdId))

	return int(rc)
}

func init() {
	defaultCallback = func(data []byte) {
		log.Printf(">> %v\n", string(data))
	}
}

// FunCallBackNormal 回调函数
type FunCallBackNormal func(data []byte)

// 用于命令回调函数的辅助
var (
	defaultCallback FunCallBackNormal     // 默认库回调函数
	mapCmdFun       sync.Map              // 函数回调函数
	mapCmdId        int32             = 1 // 回调函数Id
)

// getCmdId 获取cmdId
func getCmdId() int32 {
	if mapCmdId > math.MaxInt16 {
		atomic.StoreInt32(&mapCmdId, 1)
	}

	for {
		v := atomic.LoadInt32(&mapCmdId)
		if atomic.CompareAndSwapInt32(&mapCmdId, v, v+1) {
			return v + 1
		}
	}
}
