package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	pp "github.com/Frontware/promptpay"

	"github.com/google/uuid"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
)

type helper struct{}

// ConvertStrToFloat64 implements port.IHelper.
func (h *helper) ConvertStrToFloat64(value string) (float64, error) {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.New("error converting string to float64")
	}
	return floatValue, nil
}

// GenPromptPayQrCodeString implements port.IHelper.
func (h *helper) GenPromptPayQrCodeString(phoneno string, amount float64) string {
	payment := pp.PromptPay{
		Amount:      amount,      // Positive amount
		PromptPayID: phoneno[1:], // Tax-ID/ID Card/E-Wallet
	}

	qrcode, _ := payment.Gen() // Generate string to be use in QRCode

	return qrcode
}

// ConvertJsonToStruct implements port.IHelper.
func (h *helper) ConvertJsonToStruct(jsonStr string, result interface{}) error {
	return json.Unmarshal([]byte(jsonStr), result)
}

// ConvertStructToStrJson implements port.IHelper.
func (h *helper) ConvertStructToStrJson(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("error in function ConvertStructToStrJson adapter %s\n", err)
		return "", errors.New("something went wrong convert structure data to json")
	}
	return string(jsonBytes), nil
}

// GenUuid implements IHelper.
func (*helper) GenUUID() string {
	return uuid.NewString()
}

func NewHelper() port.IHelper {
	return &helper{}
}
