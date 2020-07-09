/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package convert

import (
	"reflect"
	"testing"
)

func Test_base62Decimal_Decode(t *testing.T) {
	type args struct {
		str62 string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Decode",
			args: args{
				"a",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base62Decimal{}
			if got := b.Decode(tt.args.str62); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base62Decimal_Encode(t *testing.T) {
	type args struct {
		num interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Decode",
			args: args{
				10,
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &base62Decimal{}
			if got := b.Encode(tt.args.num); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	bd := &base62Decimal{}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bd.Encode(i)
	}
	b.StopTimer()
}

func BenchmarkDecode(b *testing.B) {
	bd := &base62Decimal{}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bd.Decode("aA1vL")
	}
	b.StopTimer()
}
