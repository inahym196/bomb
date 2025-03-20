package shared

import "testing"

func TestNumToExcelColumn(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0->A", args{0}, "A"},
		{"26->AA", args{26}, "AA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumToExcelColumn(tt.args.n); got != tt.want {
				t.Errorf("NumToExcelColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
