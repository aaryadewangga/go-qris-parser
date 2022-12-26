package main

import (
	qrisParse "github.com/aaryadewangga/go-qris-parse/parse"
)

func QrisParser(qrisCode string) *qrisParse.QrisParseResponse {
	res := qrisParse.Parse(qrisCode)
	return res
}
