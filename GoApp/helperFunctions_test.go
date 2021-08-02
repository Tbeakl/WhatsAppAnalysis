package main

import (
	"testing"
)

func Test_mostCommonString(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "MostCommonEmptyString",
			args: args{
				values: []string{"", "", "Hello"},
			},
			want: "",
		},
		{
			name: "MostCommonNotEmpty",
			args: args{
				values: []string{"Hello", "Hello", "World", "World", "Hello", "World", "World"},
			},
			want: "World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonString(tt.args.values...); got != tt.want {
				t.Errorf("mostCommonString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageStringLength(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "AverageLength0",
			args: args{
				values: []string{"", "", ""},
			},
			want: 0.0,
		},
		{
			name: "AllSameLength",
			args: args{
				values: []string{"Hello", "Hello", "World", "World", "Hello", "World", "World"},
			},
			want: 5.0,
		},
		{
			name: "DifferingLengthsWholeNumberAverage",
			args: args{
				values: []string{"A", "AA", "AAA", "AAAA", "AAAAA"},
			},
			want: 3.0,
		},
		{
			name: "DifferingLengthsRationalNumberAverage",
			args: args{
				values: []string{"A", "AA", "AAA", "AAAA"},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageStringLength(tt.args.values...); got != tt.want {
				t.Errorf("averageStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageNumberOfWords(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "EmptyStrings",
			args: args{
				values: []string{"", "", ""},
			},
			want: 0.0,
		},
		{
			name: "AllSpaceStrings",
			args: args{
				values: []string{" ", "  ", "   ", "    "},
			},
			want: 0.0,
		},
		{
			name: "AllSameNumberOfWords",
			args: args{
				values: []string{"The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat."},
			},
			want: 6.0,
		},
		{
			name: "Differing number of words",
			args: args{
				values: []string{"The", "The cat ", "The cat sat ", "The cat sat on ", "The cat sat on the"},
			},
			want: 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageNumberOfWords(tt.args.values...); got != tt.want {
				t.Errorf("averageNumberOfWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxStringLengthWords(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "EmptyStrings",
			args: args{
				values: []string{"", "", ""},
			},
			want: 0,
		},
		{
			name: "AllSpaceStrings",
			args: args{
				values: []string{" ", "  ", "   ", "    "},
			},
			want: 0,
		},
		{
			name: "AllSameNumberOfWords",
			args: args{
				values: []string{"The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat."},
			},
			want: 6,
		},
		{
			name: "Differing lengths",
			args: args{
				values: []string{"The", "The cat ", "The cat sat on the", "The cat sat ", "The cat sat on "},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxStringLengthWords(tt.args.values...); got != tt.want {
				t.Errorf("findMaxStringLengthWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxStringLengthCharacters(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "EmptyStrings",
			args: args{
				values: []string{"", "", ""},
			},
			want: 0,
		},
		{
			name: "AllSpaceStrings",
			args: args{
				values: []string{" ", "  ", "   ", "    "},
			},
			want: 4,
		},
		{
			name: "AllSameNumberOfWords",
			args: args{
				values: []string{"The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat.", "The cat sat on the mat."},
			},
			want: 23,
		},
		{
			name: "Differing lengths",
			args: args{
				values: []string{"The", "The cat ", "The cat sat on the", "The cat sat ", "The cat sat on "},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxStringLengthCharacters(tt.args.values...); got != tt.want {
				t.Errorf("findMaxStringLengthCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
