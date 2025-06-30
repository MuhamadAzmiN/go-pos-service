package utils

import (
	"time"
)

func GenerateInvoiceNumber() string {
	return time.Now().Format("20060102")
}