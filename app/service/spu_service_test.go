package service

import "testing"

func Test_getSpuSaleTimeStamp(t *testing.T) {
	type args struct {
		saleAtString string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "2020",
			args: args{
				saleAtString: "2020-01-01 00:00:00",
			},
			want: 1577836800,
		},
		{
			name: "2021",
			args: args{
				saleAtString: "2021-01-01 00:00:00",
			},
			want: 1609459200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := getSpuSaleTimeStamp(tt.args.saleAtString); got != tt.want {
				t.Errorf("getSpuSaleTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
