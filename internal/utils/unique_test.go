package utils

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestFindUniqueLines(t *testing.T) {
	type args struct {
		input *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"Sorted list of string",
			args{
				bufio.NewScanner(strings.NewReader("Chaucer\nChaucer\nLarkin\nLarkin\nOrwell")),
			},
			map[string]int{"Chaucer": 2, "Larkin": 2, "Orwell": 1},
		},
		{
			"Un Sorted list of string",
			args{
				bufio.NewScanner(strings.NewReader("Chaucer\nLarkin\nOrwell\nChaucer\nLarkin")),
			},
			map[string]int{"Chaucer": 2, "Larkin": 2, "Orwell": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := FindUniqueLines(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) && len(got) != 3 {
				t.Errorf("FindUniqueLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
