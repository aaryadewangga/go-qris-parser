package parse

import (
	"strconv"
	"strings"
)

func Parse(qris string) *QrisParseResponse {

	data := qris
	tags := []string{"00", "01", "26", "51", "52", "53", "54", "58", "59", "60", "61", "62", "63"}
	value := []string{}

	for _, v := range tags {

		// check if not include amount
		if data[:2] != v {
			continue
		}

		// get all of index
		tag := strings.Index(data, v)
		lenght := data[tag+2 : tag+4]
		infoLenght, _ := strconv.ParseInt(lenght, 10, 10)
		maxInfoIndex := tag + 4 + int(infoLenght)

		// get raw value and append to new array
		info := data[tag:maxInfoIndex]
		value = append(value, info)

		// update the raw data
		data = data[maxInfoIndex:]
	}

	res := getDataFromValueTag(value)

	return res
}

func getValueForTag26(val string) []string {
	data := val[4:]
	tags := []string{"00", "01", "02", "03"}
	value := []string{}
	info := []string{}

	for _, v := range tags {
		// get all of index
		tag := strings.Index(data, v)
		lenght := data[tag+2 : tag+4]
		infoLenght, _ := strconv.ParseInt(lenght, 10, 10)
		maxInfoIndex := tag + 4 + int(infoLenght)

		// get raw value and append to new array
		info := data[tag:maxInfoIndex]
		value = append(value, info)

		// update the raw data
		data = data[maxInfoIndex:]
	}

	for _, v := range value {
		if v[:2] == "00" {
			info = append(info, v[4:])
		}
		if v[:2] == "01" {
			info = append(info, v[4:])
		}
		if v[:2] == "02" {
			info = append(info, v[4:])
		}
		if v[:2] == "03" {
			info = append(info, v[4:])
		}
	}

	return info
}

func getDataFromValueTag(val []string) *QrisParseResponse {
	res := &QrisParseResponse{}
	for _, v := range val {
		if v[:2] == "00" {
			res.PayloadFormatIndicator = v[4:]
		}

		if v[:2] == "01" {
			res.PointOfInitiationMethod = v[4:]
		}

		if v[:2] == "26" {
			res26 := getValueForTag26(val[2])
			res.QRISIssuer = res26[0]
			res.QRISGlobalMerchantID = res26[1]
			res.QRISAcquirerMerchantID = res26[2]
			res.MerchantBusinessType = res26[3]
		}

		if v[:2] == "52" {
			res.MerchantCategoryCode = v[4:]
		}

		if v[:2] == "53" {
			res.TransactionCurrency = v[4:]
		}

		if v[:2] == "54" {
			res.TransactionAmount = v[4:]
		}

		if v[:2] == "58" {
			res.CountryCode = v[4:]
		}

		if v[:2] == "59" {
			res.MerchantName = v[4:]
		}

		if v[:2] == "60" {
			res.MerchantCity = v[4:]
		}

		if v[:2] == "61" {
			res.PostalCode = v[4:]
		}

		if v[:2] == "63" {
			res.CRC = v[4:]
		}
	}
	return res
}
