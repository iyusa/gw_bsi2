package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sengaja
	"github.com/ussidev/permata_trx/common"
)

var authDB *gorm.DB
var authMaps []Permata = make([]Permata, 0)
var ApiKeys map[string]string
var OrgNames map[string]string

// Initialize database
func InitAuth() (err error) {
	if authDB != nil {
		return
	}

	authDB, err = gorm.Open("mysql", common.Config.AuthDB)
	if err != nil {
		return
	}
	err = LoadAuthMap()
	if err != nil {
		return
	}
	return
}

// Close connection
func CloseAuth() {
	if authDB != nil {
		authDB.Close()
	}
}

type Permata struct {
	InstitutionCode  string `gorm:"column:institution_code"`
	OrganizationName string `gorm:"column:organization_name"`
	ApiKey           string `gorm:"column:api_key"`
	ClientID         string `gorm:"column:client_id"`
	ClientSecret     string `gorm:"column:client_secret"`
	StaticKey        string `gorm:"column:permata_static_key"`
	Note             string `gorm:"column:note"`
}

func (b *Permata) TableName() string {
	return "credential_permata"
}

func LoadAuthMap() (err error) {
	err = authDB.Find(&authMaps).Error
	if err != nil {
		return
	}

	OrgNames = make(map[string]string)
	for _, x := range authMaps {
		OrgNames[x.InstitutionCode] = x.OrganizationName
	}

	// load ussi cred
	var ussiCreds []UssiCred
	err = db.Find(&ussiCreds).Error
	if err == nil {
		ApiKeys = make(map[string]string)
		for _, v := range ussiCreds {
			ApiKeys[v.InstitutionCode] = v.ApiKey
		}
	}

	return
}

// GetApiKey get auth key based on organization name, return empty string when not found
func GetApiKeyByName(orgName string) (apiKey, clientID, clientSecret, staticKey string) {
	for _, v := range authMaps {
		if v.OrganizationName == orgName {
			apiKey = v.ApiKey
			clientID = v.ClientID
			clientSecret = v.ClientSecret
			staticKey = v.StaticKey
			return
		}
	}
	return
}

func GetOrgName(instCode string) string {
	for _, v := range authMaps {
		if v.InstitutionCode == instCode {
			return v.OrganizationName
		}
	}
	return ""
}

/// Credential USSI

type UssiCred struct {
	InstitutionCode string `gorm:"column:institution_code"`
	ApiKey          string `gorm:"column:api_key"`
}

func (b *UssiCred) TableName() string {
	return "credential_ussi"
}
