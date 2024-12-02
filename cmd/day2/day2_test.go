package day2

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed test2.txt
	testData string
	//go:embed input2.txt
	inputData string
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{{
		name: "test",
		args: testData,
		want: "2",
	}, {
		name: "input",
		args: inputData,
		want: "299",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte(tt.args))
			got := Part1(buf)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{{
		name: "test",
		args: testData,
		want: "4",
	}, {
		name: "input",
		args: inputData,
		want: "364",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte(tt.args))
			got := Part2(buf)
			require.Equal(t, tt.want, got)
		})
	}
}