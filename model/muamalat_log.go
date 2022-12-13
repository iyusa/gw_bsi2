package model

import (
	"database/sql"
	"fmt"

	"github.com/ussidev/bjb/common"
)

// MuamalatLog table
type MuamalatLog struct {
	Aggregator        string          `gorm:"column:aggregator" json:"aggregator"`
	BCN               string          `gorm:"column:bcn" json:"bcn"`
	Amount            sql.NullFloat64 `gorm:"column:amount" json:"amount"`
	Admin             sql.NullFloat64 `gorm:"column:admin" json:"admin"`
	AmountBeforeAdmin sql.NullFloat64 `gorm:"column:amount_before_admin" json:"amount_before_admin"`
	CS                sql.NullString  `gorm:"column:cs" json:"cs"`
	DateTime          sql.NullString  `gorm:"column:date_time" json:"date_time"`
	DateLocalTrans    sql.NullString  `gorm:"column:dt_local_trans" json:"dt_local_trans"`
	Message           sql.NullString  `gorm:"column:msg" json:"msg"`
	MTI               sql.NullString  `gorm:"column:mti" json:"mti"`
	ProcessingCode    sql.NullString  `gorm:"column:process_code" json:"process_code"`
	ProductCode       sql.NullString  `gorm:"column:product_code" json:"product_code"`
	ResponseCode      sql.NullString  `gorm:"column:response_code" json:"response_code"`
	RevAdvCount       sql.NullInt64   `gorm:"column:rev_adv_count" json:"rev_adv_count"`
	Stan              sql.NullString  `gorm:"column:stan" json:"stan"`
	TrxID             int             `gorm:"column:trxid;primary_key" json:"trxid"`
	VaNO              sql.NullString  `gorm:"column:vano" json:"vano"`
	UUID              string          `gorm:"column:uuid" json:"uuid"`
	Description       string          `gorm:"column:description" json:"description"`
	Reference         string          `gorm:"column:ref_number" json:"ref_number"` // confirm to
}

// TableName sets the insert table name for this struct type
func (m *MuamalatLog) TableName() string {
	return "muamalat_log"
}

// Save table
func (m *MuamalatLog) Save() error {
	return db.Save(m).Error
}

// InsertLog to MuamalatLog
// TODO: Cari duplicate
func InsertLog(cs, dateTime, localTime, message, mti, processingCode, productCode, rc, stan, vano, uuid, reference string, advCount int64, amountBefore, admin float64) error {
	fmt.Printf("Prepare to insert log with db: %v\n", db)
	var m MuamalatLog

	m.CS.String = cs
	m.CS.Valid = true
	m.DateTime.String = dateTime
	m.DateTime.Valid = true
	m.DateLocalTrans.String = localTime
	m.DateLocalTrans.Valid = true
	m.Message.String = message
	m.Message.Valid = true
	m.MTI.String = mti
	m.MTI.Valid = true
	m.ProcessingCode.String = processingCode
	m.ProcessingCode.Valid = true
	m.ProductCode.String = productCode
	m.ProductCode.Valid = true
	m.ResponseCode.String = rc
	m.ResponseCode.Valid = true
	m.RevAdvCount.Int64 = advCount
	m.RevAdvCount.Valid = true
	m.Stan.String = stan
	m.Stan.Valid = true
	m.VaNO.String = vano
	m.VaNO.Valid = true
	m.Amount.Float64 = amountBefore - admin
	m.Amount.Valid = true
	m.AmountBeforeAdmin.Float64 = amountBefore
	m.AmountBeforeAdmin.Valid = true
	m.Admin.Float64 = admin
	m.Admin.Valid = true
	m.UUID = uuid
	m.Description = ""
	m.Reference = reference
	m.Aggregator = "bjb"
	m.BCN = "013"

	err := m.Save()
	common.WarnIfError(err)
	return err
}
