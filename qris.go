package qris

import (
	"errors"
	"strconv"
)

type Qris struct {
	PayloadFormatIndicator                              string                       // 00
	PointOfInitiationMethod                             string                       // 01
	MerchantAccountInformationVisa                      interface{}                  // 02-03
	MerchantAccountInformationMasterCard                interface{}                  //04-05
	MerchantAccountInformationEMVCo                     interface{}                  // 06-08, 17-25
	MerchantAccountInformationDiscover                  interface{}                  // 09-10
	MerchantAccountInformationAmex                      interface{}                  // 11-12
	MerchantAccountInformationJCB                       interface{}                  // 13-14
	MerchantAccountInformationUnionPay                  interface{}                  // 15-16
	MerchantAccountInformationDomestic                  MerchantAccountInformationID // 26-45
	MerchantAccountInformationReservedDomesticId        interface{}                  // 46-50
	MerchantAccountInformationDomesticCentralRepository interface{}                  // 51
	MerchantCategoryCode                                string                       // 52
	TransactionCurrency                                 string                       // 53
	TransactionAmount                                   string                       // 54
	TipOrConvenienceIndicator                           string                       // 55
	ValueOfConvenienceFeeFixed                          string                       // 56
	ValueOfConvenienceFeePercentage                     string                       // 57
	CountryCode                                         string                       // 58
	MerchantName                                        string                       // 59
	MerchantCity                                        string                       // 60
	PostalCode                                          string                       // 61
	AdditionalData                                      AdditionalData               // 62
	CRC                                                 string                       // 63
	MerchantInformationLanguage                         MerchantInformationLanguage  // 64
	RFU                                                 interface{}                  // 65-79
	Unreserved                                          interface{}                  //80-99
}

type MerchantAccountInformationID struct {
	ReverseDomain string // 00
	GlobalID      string // 01
	ID            string // 02
	Type          string // 03
}

type AdditionalData struct {
	BillNumber                    string      // 01
	MobileNumber                  string      // 02
	StoreLabel                    string      // 03
	LoyaltyNumber                 string      // 04
	ReferenceLabel                string      // 05
	CustomerLabel                 string      // 06
	TerminalLabel                 string      // 07
	PurposeOfTransaction          string      // 08
	AdditionalConsumerDataRequest string      // 09
	MerchantTaxID                 string      // 10
	MerchantChannel               string      // 11
	RFU                           interface{} // 12-49
	PaymentSystemSpecific         interface{} // 50-99
}

type MerchantInformationLanguage struct {
	LanguagePreference            string      // 00
	MerchantNameAlternateLanguage string      // 01
	MerchantCityAlternate         string      // 02
	RFU                           interface{} // 03-99
}

func (q *Qris) Parse(data string) error {
	for len(data) >= 4 {
		tag := data[:2]
		lengthString := data[2:4]

		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return err
		}

		data = data[4:]
		if len(data) < length {
			return errors.New("invalid QRIS")
		}

		value := data[:length]
		data = data[length:]

		if tag == "00" {
			q.PayloadFormatIndicator = value
		} else if tag == "01" {
			if value == "11" {
				q.PointOfInitiationMethod = "static"
			} else if value == "12" {
				q.PointOfInitiationMethod = "dynamic"
			}
		} else if inRange(tag, "26", "45") {
			q.MerchantAccountInformationDomestic = q.parseMerchantAccountInformationID(value)
		} else if tag == "52" {
			q.MerchantCategoryCode = value
		} else if tag == "53" {
			q.TransactionCurrency = value
		} else if tag == "54" {
			q.TransactionAmount = value
		} else if tag == "55" {
			q.TipOrConvenienceIndicator = value
		} else if tag == "56" {
			q.ValueOfConvenienceFeeFixed = value
		} else if tag == "57" {
			q.ValueOfConvenienceFeePercentage = value
		} else if tag == "58" {
			q.CountryCode = value
		} else if tag == "59" {
			q.MerchantName = value
		} else if tag == "60" {
			q.MerchantCity = value
		} else if tag == "61" {
			q.PostalCode = value
		} else if tag == "62" {
			q.AdditionalData = q.parseAdditionalData(value)
		} else if tag == "63" {
			q.CRC = value
		} else if tag == "64" {
			q.MerchantInformationLanguage = q.parseMerchantInformationLanguage(value)
		} else if inRange(tag, "65", "79") {
			q.RFU = value
		} else if inRange(tag, "80", "99") {
			q.Unreserved = value
		}
	}

	return nil
}

func (q *Qris) parseMerchantAccountInformationID(data string) MerchantAccountInformationID {
	result := MerchantAccountInformationID{}
	for len(data) >= 4 {
		tag := data[:2]
		lengthString := data[2:4]

		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return result
		}

		data = data[4:]
		if len(data) < length {
			return result
		}

		value := data[:length]
		data = data[length:]

		if tag == "00" {
			result.ReverseDomain = value
		} else if tag == "01" {
			result.GlobalID = value
		} else if tag == "02" {
			result.ID = value
		} else if tag == "03" {
			result.Type = value
		}
	}

	return result
}

func (q *Qris) parseAdditionalData(data string) AdditionalData {
	result := AdditionalData{}
	for len(data) >= 4 {
		tag := data[:2]
		lengthString := data[2:4]

		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return result
		}

		data = data[4:]
		if len(data) < length {
			return result
		}

		value := data[:length]
		data = data[length:]

		if tag == "01" {
			result.BillNumber = value
		} else if tag == "02" {
			result.MobileNumber = value
		} else if tag == "03" {
			result.StoreLabel = value
		} else if tag == "04" {
			result.LoyaltyNumber = value
		} else if tag == "05" {
			result.ReferenceLabel = value
		} else if tag == "06" {
			result.CustomerLabel = value
		} else if tag == "07" {
			result.TerminalLabel = value
		} else if tag == "08" {
			result.PurposeOfTransaction = value
		} else if tag == "09" {
			result.AdditionalConsumerDataRequest = value
		} else if tag == "10" {
			result.MerchantTaxID = value
		} else if tag == "11" {
			result.MerchantChannel = value
		} else if inRange(tag, "12", "49") {
			result.RFU = value
		} else if inRange(tag, "50", "99") {
			result.PaymentSystemSpecific = value
		}
	}

	return result
}

func (q *Qris) parseMerchantInformationLanguage(data string) MerchantInformationLanguage {
	result := MerchantInformationLanguage{}
	for len(data) >= 4 {
		tag := data[:2]
		lengthString := data[2:4]

		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return result
		}

		data = data[4:]
		if len(data) < length {
			return result
		}

		value := data[:length]
		data = data[length:]

		if tag == "00" {
			result.LanguagePreference = value
		} else if tag == "01" {
			result.MerchantNameAlternateLanguage = value
		} else if tag == "02" {
			result.MerchantCityAlternate = value
		} else if inRange(tag, "03", "99") {
			result.RFU = value
		}
	}

	return result
}
