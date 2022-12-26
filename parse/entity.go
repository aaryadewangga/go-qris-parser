package parse

type QrisParseResponse struct {
	PayloadFormatIndicator  string
	PointOfInitiationMethod string
	QRISIssuer              string
	QRISGlobalMerchantID    string
	QRISAcquirerMerchantID  string
	MerchantBusinessType    string
	MerchantCategoryCode    string
	TransactionCurrency     string
	TransactionAmount       string
	CountryCode             string
	MerchantName            string
	MerchantCity            string
	PostalCode              string
	CRC                     string
}
