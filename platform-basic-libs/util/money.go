package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
)

// Money 金额，单位为分
type Money int64

// ParseMoney 解析成金额
// decimal 必须为0或2
func ParseMoney(decimal int, v interface{}) (Money, error) {
	var m Money
	err := m.Parse(decimal, v)
	return m, err
}

// Format 金额转成字符串
// decimal 必须为0或2
func (m Money) Format(decimal int) string {
	if decimal == 0 {
		return strconv.Itoa(int(m))
	} else if decimal == 2 {
		return strconv.FormatFloat(float64(m)/100, 'f', 2, 64)
	} else {
		panic("!!! now Money Format is only equal 0 or 2 !!!")
	}
}

// Parse 解析成金额
// decimal 必须为0或2
func (m *Money) Parse(decimal int, v interface{}) error {
	if decimal == 0 {
		return m.parseFen(v)
	} else if decimal == 2 {
		return m.parseYuan(v)
	} else {
		panic("!!! now Money Format is only equal 0 or 2 !!!")
	}
}

// parseFen 解析成金额
func (m *Money) parseFen(toParse interface{}) error {
	if m == nil {
		return errors.New("caller is invalid")
	}
	switch v := toParse.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err == nil {
			*m = Money(i)
			return nil
		}
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			*m = Money(f)
			return nil
		}
		return errors.New("can't parse money")
	case []byte:
		i, err := strconv.Atoi(string(v))
		if err == nil {
			*m = Money(i)
			return nil
		}
		f, err := strconv.ParseFloat(string(v), 64)
		if err == nil {
			*m = Money(f)
			return nil
		}
		return errors.New("can't parse money")
	case int:
		*m = Money(v)
		return nil
	case int64:
		*m = Money(v)
		return nil
	case uint:
		*m = Money(v)
		return nil
	case uint8:
		*m = Money(v)
		return nil
	case uint16:
		*m = Money(v)
		return nil
	case uint32:
		*m = Money(v)
		return nil
	case uint64:
		*m = Money(v)
		return nil
	case float32:
		*m = Money(v)
		return nil
	case float64:
		*m = Money(v)
		return nil
	default:
		return fmt.Errorf("invalid type(%T)", v)
	}
}

// parseYuan 解析成金额
func (m *Money) parseYuan(toParse interface{}) error {
	switch v := toParse.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err == nil {
			*m = Money(i) * 100
			return nil
		}
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			*m = Money(f * 100)
			return nil
		}
		return errors.New("can't parse money")
	case []byte:
		i, err := strconv.Atoi(string(v))
		if err == nil {
			*m = Money(i) * 100
			return nil
		}
		f, err := strconv.ParseFloat(string(v), 64)
		if err == nil {
			*m = Money(f * 100)
			return nil
		}
		return errors.New("can't parse money")
	case int:
		*m = Money(v) * 100
		return nil
	case int64:
		*m = Money(v) * 100
		return nil
	case uint:
		*m = Money(v) * 100
		return nil
	case uint8:
		*m = Money(v) * 100
		return nil
	case uint16:
		*m = Money(v) * 100
		return nil
	case uint32:
		*m = Money(v) * 100
		return nil
	case uint64:
		*m = Money(v) * 100
		return nil
	case float32:
		*m = Money(v * 100)
		return nil
	case float64:
		*m = Money(v * 100)
		return nil
	default:
		return fmt.Errorf("invalid type(%T)", v)
	}
}

// Scan implements the Scanner interface.
func (m *Money) Scan(value interface{}) error {
	return m.parseYuan(value)
}

// Value implements the driver Valuer interface.
func (m Money) Value() (driver.Value, error) {
	return m.Format(2), nil
}
