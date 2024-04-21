package jwtauth

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

type MapClaims jwt.MapClaims

// Exp returns value of exp
func (m MapClaims) Exp() (int64, error) {
	return m.Int64("exp")
}

// Identity returns value of identity
func (m MapClaims) Identity() (int64, error) {
	return m.Int64("identity")
}

// Int64 try to convert to int64 by key
func (m MapClaims) Int64(key string) (int64, error) {
	value := m[key]
	if value == nil {
		return 0, fmt.Errorf("invalid key '%v'", key)
	}

	switch value.(type) {
	case json.Number:
		return value.(json.Number).Int64()
	case float64:
		return int64(value.(float64)), nil
	case string:
		return strconv.ParseInt(value.(string), 10, 0)
	default:
		return 0, fmt.Errorf("invalid value '%v' type '%T'", value, value)
	}
}

// String try to convert to string by key
func (m MapClaims) String(key string) string {
	value := m[key]
	if value == nil {
		return ""
	}

	switch value.(type) {
	case json.Number:
		return value.(json.Number).String()
	case float64:
		return strconv.FormatFloat(value.(float64), 'g', -1, 64)
	case string:
		return value.(string)
	default:
		fmt.Errorf("maptoclaims key to string invalid value '%v' type '%T' key %v", value, value, key)
		return ""
	}
}