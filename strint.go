package gormext

import (
	"encoding/json"
	"strconv"
)

// StrInt create a type alias for type int
type StrInt int

// UnmarshalJSON create a custom unmarshal for the StrInt
/// this helps us check the type of our value before unmarshalling it

func (st *StrInt) UnmarshalJSON(b []byte) error {
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
		*st = StrInt(v)
	case float64:
		*st = StrInt(int(v))
	case string:
		///here convert the string into
		///an integer
		if v == "" {
			*st = StrInt(0)
			return nil
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StrInt(i)

	}
	return nil
}
