/*
 * @Author: LinkyPi trouble.linky@gmail.com
 * @Date: 2023-12-17 10:26:38
 * @LastEditors: LinkyPi trouble.linky@gmail.com
 * @LastEditTime: 2023-12-17 10:48:46
 * @FilePath: /test/interface.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "testing"

func TestWrite(t *testing.T) {
	type args struct {
		w Writer
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Write(tt.args.w)
		})
	}
}
