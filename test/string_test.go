/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2024-01-04 14:47:18
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2024-01-04 14:50:45
 * @FilePath: /test/test/string_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func StringToByte(key *string) []byte {
	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
	strPtr.Cap = strPtr.Len
	b := *(*[]byte)(unsafe.Pointer(strPtr))
	return b
}

func TestStringToBytes(t *testing.T) {
	decryptContent := "/AvYEjm4g6xJ3LVrk2/Adk"
	iv := decryptContent[0:16]
	key := decryptContent[2:18]
	fmt.Println(&iv)
	fmt.Println(&key)
	ivBytes := StringToByte(&iv)
	keyBytes := StringToByte(&key)
	fmt.Println(string(ivBytes))
	fmt.Println(string(keyBytes))
}
