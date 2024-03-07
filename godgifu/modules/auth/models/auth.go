package models

import "github.com/google/uuid"

type UserDeviceData struct {
	AccountID        uuid.UUID `db:"account_id"`
	IPAddress        string    `db:"ip_address"`
	UserAgent        string    `db:"user_agent"`
	OperatingSystem  string    `db:"operating_system"`
	ScreenResolution string    `db:"screen_resolution"`
}
