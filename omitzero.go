package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type OmitZero struct {
	Date         time.Time `json:"date"`
	DateWithOmit time.Time `json:"date_with_omit,omitzero"`
}

func omit_zero() {
	test := OmitZero{
		Date: time.Now(),
	}
	jsonTest, _ := json.MarshalIndent(test, "", "")
	fmt.Println(string(jsonTest))
	//	{
	// "date": "2025-08-24T12:36:20.086295-03:00"
	// }
}
