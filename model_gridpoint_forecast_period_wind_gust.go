/*
weather.gov API

weather.gov API

API version: 1.8.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package weatherApi

import (
	"encoding/json"
	"fmt"
)

// GridpointForecastPeriodWindGust - Peak wind gust for the period. This property as an string value is deprecated. Future versions will express this value as a quantitative value object. To make use of the future standard format now, set the \"forecast_wind_speed_qv\" feature flag on the request. 
type GridpointForecastPeriodWindGust struct {
	QuantitativeValue *QuantitativeValue
	String *string
}

// QuantitativeValueAsGridpointForecastPeriodWindGust is a convenience function that returns QuantitativeValue wrapped in GridpointForecastPeriodWindGust
func QuantitativeValueAsGridpointForecastPeriodWindGust(v *QuantitativeValue) GridpointForecastPeriodWindGust {
	return GridpointForecastPeriodWindGust{
		QuantitativeValue: v,
	}
}

// stringAsGridpointForecastPeriodWindGust is a convenience function that returns string wrapped in GridpointForecastPeriodWindGust
func StringAsGridpointForecastPeriodWindGust(v *string) GridpointForecastPeriodWindGust {
	return GridpointForecastPeriodWindGust{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GridpointForecastPeriodWindGust) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into QuantitativeValue
	err = newStrictDecoder(data).Decode(&dst.QuantitativeValue)
	if err == nil {
		jsonQuantitativeValue, _ := json.Marshal(dst.QuantitativeValue)
		if string(jsonQuantitativeValue) == "{}" { // empty struct
			dst.QuantitativeValue = nil
		} else {
			match++
		}
	} else {
		dst.QuantitativeValue = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QuantitativeValue = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GridpointForecastPeriodWindGust)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GridpointForecastPeriodWindGust)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GridpointForecastPeriodWindGust) MarshalJSON() ([]byte, error) {
	if src.QuantitativeValue != nil {
		return json.Marshal(&src.QuantitativeValue)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GridpointForecastPeriodWindGust) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QuantitativeValue != nil {
		return obj.QuantitativeValue
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGridpointForecastPeriodWindGust struct {
	value *GridpointForecastPeriodWindGust
	isSet bool
}

func (v NullableGridpointForecastPeriodWindGust) Get() *GridpointForecastPeriodWindGust {
	return v.value
}

func (v *NullableGridpointForecastPeriodWindGust) Set(val *GridpointForecastPeriodWindGust) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointForecastPeriodWindGust) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointForecastPeriodWindGust) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointForecastPeriodWindGust(val *GridpointForecastPeriodWindGust) *NullableGridpointForecastPeriodWindGust {
	return &NullableGridpointForecastPeriodWindGust{value: val, isSet: true}
}

func (v NullableGridpointForecastPeriodWindGust) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointForecastPeriodWindGust) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

