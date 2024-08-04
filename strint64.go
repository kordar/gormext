package gormext

import (
	"encoding/json"
	"strconv"
)

// StrInt64 create a type alias for type int
type StrInt64 int64

// UnmarshalJSON create a custom unmarshal for the StrInt64
/// this helps us check the type of our value before unmarshalling it

func (st *StrInt64) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into a int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
	case float64:
		*st = StrInt64(int64(v))
		break
	case int64:
		*st = StrInt64(v)
		break
	case string:
		///here convert the string into
		///an integer
		if v == "" {
			*st = StrInt64(0)
			return nil
		}
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StrInt64(i)

	}
	return nil
}
