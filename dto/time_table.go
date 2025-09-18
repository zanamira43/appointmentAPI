package dto

import "gorm.io/datatypes"

type TimeTable struct {
	PatientID   *int           `json:"patient_id"`
	PatientName string         `json:"patient_name"`
	WeekDay     datatypes.JSON `json:"week_day"`
	StartTime   string         `json:"start_time"`
	EndTime     string         `json:"end_time"`
}
