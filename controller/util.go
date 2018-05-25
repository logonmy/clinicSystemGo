package controller

import (
	"database/sql"
	"strconv"
)

// ToNullString 处理空字符串
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// ToNullInt64 处理空数字
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

// NullFloat64 处理 空float
func NullFloat64(s string) sql.NullInt64 {
	f, err := strconv.ParseFloat(s, 64)
	return sql.NullFloat64{Float64: float64(f), Valid: err == nil}
}

// NullBool 处理 bool
func NullFloat64(s string) sql.NullInt64 {
	b, err := strconv.ParseBool(s)
	return sql.NullBool{Bool: float64(b), Valid: err == nil}
}
