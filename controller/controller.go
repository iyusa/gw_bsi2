package controller

import (
	"fmt"

	"github.com/ussidev/bjb/common"
	"github.com/ussidev/bjb/model"
)

var ErrorCodes = map[string]string{
	"0000": "Sukses",
	"0005": "Error Other",
	"0012": "Unmatch Billing Type (open/fix)",
	"0013": "Invalid Amount",
	"0014": "Invalid VA Number",
	"0016": "Client Not Registered",
	"0017": "VA Already Exist for Open Payment",
	"0018": "Invalid Product Code",
	"0019": "Kota/Kabupaten Not Registered",
	"0020": "Invalid Tax Type",
	"0021": "Unregistered VA Number",
	"0022": "Client Already Exist",
	"0023": "Product Already Exist",
	"0030": "Invalid Request Message",
}

type CallbackRequest struct {
	VaNumber          string `json:"va_number"`
	ClientRefnum      string `json:"client_refnum"`
	TransactionDate   string `json:"transaction_date"`
	TransactionAmount int    `json:"transaction_amount"`
	UnpaidAmount      int    `json:"unpaid_amount"`
	Status            int    `json:"status"`
}

type CallbackResponse struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

func Execute(in *CallbackRequest) (*CallbackResponse, error) {
	out := &CallbackResponse{"0000", ErrorCodes["0000"]}

	// TODO cari va, parse, lalu cari di db
	fmt.Println("Starting controller:Execute")

	var (
		stan                 = common.RandomString(12)
		mti                  = "2200"
		processingCode       = "200700"
		productCode          = "EDUPCR"
		rc                   = "0000" // TODO tergantung in.Status
		vano                 = in.VaNumber
		uuid                 = ""
		reference            = in.ClientRefnum
		isoMessage           = ""
		advCount       int64 = 0
		amountBefore         = float64(in.TransactionAmount)
		admin                = 0.0
		va                   = vano[4:]
		iid                  = vano[0:4]
	)

	// reconnect db
	err := model.Reconnect()
	if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		fmt.Println("End controller:Execute with error")
		return nil, err
	}

	billerCode, apexCode, giroType, err := model.FindBillerAndApex(iid)
	if err != nil {
		return nil, err
	}
	fmt.Printf("BillerCode %s, apexCode %s, giroType %d", billerCode, apexCode, giroType)

	subscriberID := apexCode + va

	model.InsertLog("C", in.TransactionDate, in.TransactionDate, isoMessage, mti, processingCode, productCode, rc,
		stan, subscriberID, uuid, reference, advCount, amountBefore, admin)

	fmt.Println("End controller:Execute")

	return out, nil
}
