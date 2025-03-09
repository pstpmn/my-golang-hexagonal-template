package utils

import "fmt"

func GenerateLongOrderKey(refId string) string {
	// example output -> order:refId:REFEURKSJOQWIOKJD
	return fmt.Sprintf("order:refId:%s", refId)
}

func GenerateShortOrderKey(accno string, amount float64) string {
	// example output -> order:pending:1923482912:100.01
	return fmt.Sprintf(`order:pending:%s:%0.2f`, accno, amount)
}

func GenerateBackupShortOrderKey(accno string, amount float64) string {
	// example output -> backup:order:pending:1923482912:100.01
	return fmt.Sprintf(`backup:order:pending:%s:%0.2f`, accno, amount)
}

func GenerateCurrentDecimalOrderKey(accno string, amount int) string {
	// example output -> order:current_decimal:1923482912:100.01
	return fmt.Sprintf(`order:current_decimal:%s:%d`, accno, amount)
}

func GenerateAccountKey(token string) string {
	return fmt.Sprintf("token-%s", token)
}

func GenerateWhiteListIPKey(ipAddress string) string {
	return fmt.Sprintf("ip-%s", ipAddress)
}
