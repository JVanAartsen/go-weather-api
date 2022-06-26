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

// NWSForecastOfficeId Three-letter identifier for a NWS office.
type NWSForecastOfficeId string

// List of NWSForecastOfficeId
const (
	AKQ NWSForecastOfficeId = "AKQ"
	ALY NWSForecastOfficeId = "ALY"
	BGM NWSForecastOfficeId = "BGM"
	BOX NWSForecastOfficeId = "BOX"
	BTV NWSForecastOfficeId = "BTV"
	BUF NWSForecastOfficeId = "BUF"
	CAE NWSForecastOfficeId = "CAE"
	CAR NWSForecastOfficeId = "CAR"
	CHS NWSForecastOfficeId = "CHS"
	CLE NWSForecastOfficeId = "CLE"
	CTP NWSForecastOfficeId = "CTP"
	GSP NWSForecastOfficeId = "GSP"
	GYX NWSForecastOfficeId = "GYX"
	ILM NWSForecastOfficeId = "ILM"
	ILN NWSForecastOfficeId = "ILN"
	LWX NWSForecastOfficeId = "LWX"
	MHX NWSForecastOfficeId = "MHX"
	OKX NWSForecastOfficeId = "OKX"
	PBZ NWSForecastOfficeId = "PBZ"
	PHI NWSForecastOfficeId = "PHI"
	RAH NWSForecastOfficeId = "RAH"
	RLX NWSForecastOfficeId = "RLX"
	RNK NWSForecastOfficeId = "RNK"
	ABQ NWSForecastOfficeId = "ABQ"
	AMA NWSForecastOfficeId = "AMA"
	BMX NWSForecastOfficeId = "BMX"
	BRO NWSForecastOfficeId = "BRO"
	CRP NWSForecastOfficeId = "CRP"
	EPZ NWSForecastOfficeId = "EPZ"
	EWX NWSForecastOfficeId = "EWX"
	FFC NWSForecastOfficeId = "FFC"
	FWD NWSForecastOfficeId = "FWD"
	HGX NWSForecastOfficeId = "HGX"
	HUN NWSForecastOfficeId = "HUN"
	JAN NWSForecastOfficeId = "JAN"
	JAX NWSForecastOfficeId = "JAX"
	KEY NWSForecastOfficeId = "KEY"
	LCH NWSForecastOfficeId = "LCH"
	LIX NWSForecastOfficeId = "LIX"
	LUB NWSForecastOfficeId = "LUB"
	LZK NWSForecastOfficeId = "LZK"
	MAF NWSForecastOfficeId = "MAF"
	MEG NWSForecastOfficeId = "MEG"
	MFL NWSForecastOfficeId = "MFL"
	MLB NWSForecastOfficeId = "MLB"
	MOB NWSForecastOfficeId = "MOB"
	MRX NWSForecastOfficeId = "MRX"
	OHX NWSForecastOfficeId = "OHX"
	OUN NWSForecastOfficeId = "OUN"
	SHV NWSForecastOfficeId = "SHV"
	SJT NWSForecastOfficeId = "SJT"
	SJU NWSForecastOfficeId = "SJU"
	TAE NWSForecastOfficeId = "TAE"
	TBW NWSForecastOfficeId = "TBW"
	TSA NWSForecastOfficeId = "TSA"
	ABR NWSForecastOfficeId = "ABR"
	APX NWSForecastOfficeId = "APX"
	ARX NWSForecastOfficeId = "ARX"
	BIS NWSForecastOfficeId = "BIS"
	BOU NWSForecastOfficeId = "BOU"
	CYS NWSForecastOfficeId = "CYS"
	DDC NWSForecastOfficeId = "DDC"
	DLH NWSForecastOfficeId = "DLH"
	DMX NWSForecastOfficeId = "DMX"
	DTX NWSForecastOfficeId = "DTX"
	DVN NWSForecastOfficeId = "DVN"
	EAX NWSForecastOfficeId = "EAX"
	FGF NWSForecastOfficeId = "FGF"
	FSD NWSForecastOfficeId = "FSD"
	GID NWSForecastOfficeId = "GID"
	GJT NWSForecastOfficeId = "GJT"
	GLD NWSForecastOfficeId = "GLD"
	GRB NWSForecastOfficeId = "GRB"
	GRR NWSForecastOfficeId = "GRR"
	ICT NWSForecastOfficeId = "ICT"
	ILX NWSForecastOfficeId = "ILX"
	IND NWSForecastOfficeId = "IND"
	IWX NWSForecastOfficeId = "IWX"
	JKL NWSForecastOfficeId = "JKL"
	LBF NWSForecastOfficeId = "LBF"
	LMK NWSForecastOfficeId = "LMK"
	LOT NWSForecastOfficeId = "LOT"
	LSX NWSForecastOfficeId = "LSX"
	MKX NWSForecastOfficeId = "MKX"
	MPX NWSForecastOfficeId = "MPX"
	MQT NWSForecastOfficeId = "MQT"
	OAX NWSForecastOfficeId = "OAX"
	PAH NWSForecastOfficeId = "PAH"
	PUB NWSForecastOfficeId = "PUB"
	RIW NWSForecastOfficeId = "RIW"
	SGF NWSForecastOfficeId = "SGF"
	TOP NWSForecastOfficeId = "TOP"
	UNR NWSForecastOfficeId = "UNR"
	BOI NWSForecastOfficeId = "BOI"
	BYZ NWSForecastOfficeId = "BYZ"
	EKA NWSForecastOfficeId = "EKA"
	FGZ NWSForecastOfficeId = "FGZ"
	GGW NWSForecastOfficeId = "GGW"
	HNX NWSForecastOfficeId = "HNX"
	LKN NWSForecastOfficeId = "LKN"
	LOX NWSForecastOfficeId = "LOX"
	MFR NWSForecastOfficeId = "MFR"
	MSO NWSForecastOfficeId = "MSO"
	MTR NWSForecastOfficeId = "MTR"
	OTX NWSForecastOfficeId = "OTX"
	PDT NWSForecastOfficeId = "PDT"
	PIH NWSForecastOfficeId = "PIH"
	PQR NWSForecastOfficeId = "PQR"
	PSR NWSForecastOfficeId = "PSR"
	REV NWSForecastOfficeId = "REV"
	SEW NWSForecastOfficeId = "SEW"
	SGX NWSForecastOfficeId = "SGX"
	SLC NWSForecastOfficeId = "SLC"
	STO NWSForecastOfficeId = "STO"
	TFX NWSForecastOfficeId = "TFX"
	TWC NWSForecastOfficeId = "TWC"
	VEF NWSForecastOfficeId = "VEF"
	AER NWSForecastOfficeId = "AER"
	AFC NWSForecastOfficeId = "AFC"
	AFG NWSForecastOfficeId = "AFG"
	AJK NWSForecastOfficeId = "AJK"
	ALU NWSForecastOfficeId = "ALU"
	GUM NWSForecastOfficeId = "GUM"
	HPA NWSForecastOfficeId = "HPA"
	HFO NWSForecastOfficeId = "HFO"
	PPG NWSForecastOfficeId = "PPG"
	STU NWSForecastOfficeId = "STU"
	NH1 NWSForecastOfficeId = "NH1"
	NH2 NWSForecastOfficeId = "NH2"
	ONA NWSForecastOfficeId = "ONA"
	ONP NWSForecastOfficeId = "ONP"
)

// All allowed values of NWSForecastOfficeId enum
var AllowedNWSForecastOfficeIdEnumValues = []NWSForecastOfficeId{
	"AKQ",
	"ALY",
	"BGM",
	"BOX",
	"BTV",
	"BUF",
	"CAE",
	"CAR",
	"CHS",
	"CLE",
	"CTP",
	"GSP",
	"GYX",
	"ILM",
	"ILN",
	"LWX",
	"MHX",
	"OKX",
	"PBZ",
	"PHI",
	"RAH",
	"RLX",
	"RNK",
	"ABQ",
	"AMA",
	"BMX",
	"BRO",
	"CRP",
	"EPZ",
	"EWX",
	"FFC",
	"FWD",
	"HGX",
	"HUN",
	"JAN",
	"JAX",
	"KEY",
	"LCH",
	"LIX",
	"LUB",
	"LZK",
	"MAF",
	"MEG",
	"MFL",
	"MLB",
	"MOB",
	"MRX",
	"OHX",
	"OUN",
	"SHV",
	"SJT",
	"SJU",
	"TAE",
	"TBW",
	"TSA",
	"ABR",
	"APX",
	"ARX",
	"BIS",
	"BOU",
	"CYS",
	"DDC",
	"DLH",
	"DMX",
	"DTX",
	"DVN",
	"EAX",
	"FGF",
	"FSD",
	"GID",
	"GJT",
	"GLD",
	"GRB",
	"GRR",
	"ICT",
	"ILX",
	"IND",
	"IWX",
	"JKL",
	"LBF",
	"LMK",
	"LOT",
	"LSX",
	"MKX",
	"MPX",
	"MQT",
	"OAX",
	"PAH",
	"PUB",
	"RIW",
	"SGF",
	"TOP",
	"UNR",
	"BOI",
	"BYZ",
	"EKA",
	"FGZ",
	"GGW",
	"HNX",
	"LKN",
	"LOX",
	"MFR",
	"MSO",
	"MTR",
	"OTX",
	"PDT",
	"PIH",
	"PQR",
	"PSR",
	"REV",
	"SEW",
	"SGX",
	"SLC",
	"STO",
	"TFX",
	"TWC",
	"VEF",
	"AER",
	"AFC",
	"AFG",
	"AJK",
	"ALU",
	"GUM",
	"HPA",
	"HFO",
	"PPG",
	"STU",
	"NH1",
	"NH2",
	"ONA",
	"ONP",
}

func (v *NWSForecastOfficeId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NWSForecastOfficeId(value)
	for _, existing := range AllowedNWSForecastOfficeIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NWSForecastOfficeId", value)
}

// NewNWSForecastOfficeIdFromValue returns a pointer to a valid NWSForecastOfficeId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNWSForecastOfficeIdFromValue(v string) (*NWSForecastOfficeId, error) {
	ev := NWSForecastOfficeId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NWSForecastOfficeId: valid values are %v", v, AllowedNWSForecastOfficeIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NWSForecastOfficeId) IsValid() bool {
	for _, existing := range AllowedNWSForecastOfficeIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NWSForecastOfficeId value
func (v NWSForecastOfficeId) Ptr() *NWSForecastOfficeId {
	return &v
}

type NullableNWSForecastOfficeId struct {
	value *NWSForecastOfficeId
	isSet bool
}

func (v NullableNWSForecastOfficeId) Get() *NWSForecastOfficeId {
	return v.value
}

func (v *NullableNWSForecastOfficeId) Set(val *NWSForecastOfficeId) {
	v.value = val
	v.isSet = true
}

func (v NullableNWSForecastOfficeId) IsSet() bool {
	return v.isSet
}

func (v *NullableNWSForecastOfficeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNWSForecastOfficeId(val *NWSForecastOfficeId) *NullableNWSForecastOfficeId {
	return &NullableNWSForecastOfficeId{value: val, isSet: true}
}

func (v NullableNWSForecastOfficeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNWSForecastOfficeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

