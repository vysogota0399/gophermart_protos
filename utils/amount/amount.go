package amount

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/type/money"
)

type Amount struct {
	Money *money.Money `json:"money"`
}

func New(m *money.Money) *Amount {
	return &Amount{Money: m}
}

func (a *Amount) Float64() float64 {
	units := a.Money.GetUnits()
	nanos := a.Money.GetNanos()

	floatUnits := float64(units)
	floatNanos := float64(nanos) / 1e9

	return floatUnits + floatNanos
}

func (a *Amount) NanoBonuses() int64 {
	return a.Money.Units*1e9 + int64(a.Money.Nanos)
}

func FromInt64(value int64) *Amount {
	units := value / 1e9
	nanos := value % 1e9
	return &Amount{Money: &money.Money{Units: units, Nanos: int32(nanos)}}
}

func FromFloat64(value float64) (*Amount, error) {
	valueStr := fmt.Sprintf("%.2f", value)
	parts := strings.Split(valueStr, ".")
	var units int64
	var nanos int

	if len(parts) == 2 {
		intPart := parts[0]
		fracPart := parts[1]

		var err error
		units, err = strconv.ParseInt(intPart, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("convert to int failed error %w", err)
		}

		nanos, err = strconv.Atoi(fracPart + strings.Repeat("0", 9-len(fracPart)))
		if err != nil {
			return nil, fmt.Errorf("convert to int failed error: %w", err)
		}
	} else if len(parts) == 1 {
		intPart := parts[0]

		var err error
		units, err = strconv.ParseInt(intPart, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("convert to int failed error: %w", err)
		}

		nanos = 0
	}
	m := &money.Money{
		CurrencyCode: "RUB",
		Units:        units,
		Nanos:        int32(nanos),
	}

	return &Amount{Money: m}, nil
}
