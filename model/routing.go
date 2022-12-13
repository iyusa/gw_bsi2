package model

import (
	"database/sql"
	"errors"
)

// Routing struct
type Routing struct {
	ApexBankCode  sql.NullString `gorm:"column:apex_bank_code" json:"apex_bank_code"`
	BillerCode    sql.NullString `gorm:"column:biller_code" json:"biller_code"`
	Host          sql.NullString `gorm:"column:host" json:"host"`
	InstitutionID string         `gorm:"column:institution_id;primary_key" json:"institution_id"`
	Note          sql.NullString `gorm:"column:note" json:"note"`
	Port          sql.NullInt64  `gorm:"column:port" json:"port"`
	Target        sql.NullInt64  `gorm:"column:target" json:"target"`
	GiroType      int            `gorm:"column:giro_type" json:"giro_type"`
}

// TableName sets the insert table name for this struct type
func (r *Routing) TableName() string {
	return "routing"
}

// FindBillerAndApex cari biller code & apex code
func FindBillerAndApex(id string) (billerCode string, apexCode string, giroType int, err error) {
	var r Routing
	if db.First(&r, "institution_id = ?", id).RecordNotFound() {
		err = errors.New("kode institusi tidak dikenal")
		giroType = 1
		return
	}
	apexCode = r.ApexBankCode.String
	billerCode = r.BillerCode.String

	// versi baru
	// giroType = r.GiroType
	// if giroType == 2 {
	// 	a, e := FindApex(id)
	// 	if e != nil {
	// 		log.Printf("Gagal cari kode institusi: %v\n", e)
	// 	} else {
	// 		apexCode = a
	// 	}
	// }

	return
}
