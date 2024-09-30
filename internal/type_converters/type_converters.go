package type_converters

import (
	"database/sql"
	"strconv"
	"time"
)

// TimeToInt64String converts a time.Time to a string representation of Unix time.
// It takes a time.Time value as input and returns a string.
//
// Parameters:
// - value: The time.Time value to be converted.
//
// Returns:
// - A string representation of the Unix time.
func TimeToInt64String(value time.Time) string {
	result := strconv.FormatInt(value.Unix(), 10)

	return result
}

// SQLNullTimeToInt64String converts a sql.NullTime to a string representation of Unix time.
// If the sql.NullTime is valid, it returns the Unix time as a string. If it is not valid, it returns an empty string.
//
// Parameters:
// - nullTime: The sql.NullTime value to be converted.
//
// Returns:
// - A string representation of the Unix time if the sql.NullTime is valid, otherwise an empty string.
func SQLNullTimeToInt64String(value sql.NullTime) string {
	resultInt64 := int64(0)

	if value.Valid {
		resultInt64 = value.Time.Unix()
	}

	result := strconv.FormatInt(resultInt64, 10)

	return result
}

// Int64ToSQLNullTime converts an int64 value to a sql.NullTime.
// If the value is non-zero, it sets the Valid field to true and
// converts the value to a time.Time using time.Unix. If the value
// is zero, it sets the Valid field to false.
func Int64ToSQLNullTime(value int64) sql.NullTime {
	result := sql.NullTime{
		Time:  time.Unix(value, 0),
		Valid: value != 0,
	}

	return result
}

// StringToSQLNullString converts a string to a sql.NullString.
// If the input string is not empty, the resulting sql.NullString will be valid.
// If the input string is empty, the resulting sql.NullString will be invalid.
//
// Parameters:
//   - value: The input string to be converted.
//
// Returns:
//   - sql.NullString: The resulting sql.NullString with the appropriate validity.
func StringToSQLNullString(value string) sql.NullString {
	isValid := len(value) != 0
	result := sql.NullString{String: value, Valid: isValid}

	return result
}
