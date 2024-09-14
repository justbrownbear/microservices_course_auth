package type_converters

import (
	"database/sql"
	"strconv"
	"time"
)



func TimeToInt64String( value time.Time ) string {
	result := strconv.FormatInt( value.Unix(), 10 )

	return result
}


func SqlNullTimeToInt64String( value sql.NullTime ) string {
	resultInt64 := int64(0)

	if value.Valid {
		resultInt64 = value.Time.Unix()
	}

	result := strconv.FormatInt( resultInt64, 10 )

	return result
}


func Int64ToSqlNullTime( value int64 ) sql.NullTime {
	result := sql.NullTime{
		Time: time.Unix( value, 0 ),
		Valid: value != 0,
	}

	return result
}


func StringToSqlNullString( value string ) sql.NullString {
	isValid := len(value) != 0
	result := sql.NullString{String: value, Valid: isValid}

	return result;
}
