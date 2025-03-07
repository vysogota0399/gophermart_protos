package amount

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/money"
)

func TestAmount_Float64(t *testing.T) {
	type fields struct {
		money *money.Money
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "without nanos",
			fields: fields{
				money: &money.Money{
					Units: 100,
				},
			},
			want: 100.0,
		},
		{
			name: "with nanos",
			fields: fields{
				money: &money.Money{
					Units: 100,
					Nanos: 710_000_000,
				},
			},
			want: 100.71,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Amount{
				Money: tt.fields.money,
			}
			got := a.Float64()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFromFloat64(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want *Amount
	}{
		{
			name: "20.0",
			args: args{
				value: 20.0,
			},
			want: &Amount{
				Money: &money.Money{
					Units: 20,
					Nanos: 000_000_000,
				},
			},
		},
		{
			name: "20.25",
			args: args{
				value: 20.25,
			},
			want: &Amount{
				Money: &money.Money{
					Units: 20,
					Nanos: 250_000_000,
				},
			},
		},
		{
			name: "0.25",
			args: args{
				value: 0.25,
			},
			want: &Amount{
				Money: &money.Money{
					Units: 0,
					Nanos: 250_000_000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromFloat64(tt.args.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Money.Nanos, got.Money.Nanos)
			assert.Equal(t, tt.want.Money.Units, got.Money.Units)
		})
	}
}

func TestAmount_NanoBonuses(t *testing.T) {
	type fields struct {
		Money *money.Money
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "with Units",
			fields: fields{
				Money: &money.Money{
					Nanos: 710_000_000,
					Units: 1234,
				},
			},
			want: 1234_710_000_000,
		},
		{
			name: "withour Units",
			fields: fields{
				Money: &money.Money{
					Nanos: 710_000_000,
					Units: 0,
				},
			},
			want: 710_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Amount{
				Money: tt.fields.Money,
			}
			got := a.NanoBonuses()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFromInt64(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want *Amount
	}{
		{
			name: "123000000000",
			args: args{value: 123_000_000_000},
			want: &Amount{Money: &money.Money{Nanos: 000_000_000, Units: 123}},
		},
		{
			name: "700000000",
			args: args{value: 750_000_000},
			want: &Amount{Money: &money.Money{Nanos: 750_000_000, Units: 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromInt64(tt.args.value)
			assert.Equal(t, tt.want.Money.Nanos, got.Money.Nanos)
			assert.Equal(t, tt.want.Money.Units, got.Money.Units)
		})
	}
}
