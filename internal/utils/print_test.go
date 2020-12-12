package utils

import (
	"bytes"
	"testing"
)

func TestPrintStringCount(t *testing.T) {
	type args struct {
		keys  []string
		input map[string]int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"Positive Scenario Test",
			args{
				[]string{"Chaucer", "Larkin", "Orwell"},
				map[string]int{"Chaucer": 2, "Larkin": 2, "Orwell": 1},
			},
			"2\tChaucer\n2\tLarkin\n1\tOrwell\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			PrintStringCount(tt.args.keys, tt.args.input, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("PrintStringCount() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
