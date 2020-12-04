package godebug

import (
	"testing"
)

func Test_LogPrint(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"test string", args{v: []interface{}{"test", "string"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogPrintln(tt.args.v...)
		})
	}
}
