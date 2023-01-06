package qris

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testData         = "00020101021226530012COM.XXXX.WWW0118120000123400001234020412340303UMI51380014ID.CO.QRIS.WWW02091234567890303UMI5204723053033605405100005802ID5915TOKO MIXUE BARU6015Kabupaten Tokyo6106123131622001MIXUE BARU70703K19630433CE"
	testDataNoAmount = "00020101021226640013COM.MYWEB.WWW01181234567890123456780214123456789012340303UKE5909QRIS ARYA6015Jakarta Selatan6304XXXX"
)

func TestQrisParseWithAmount(t *testing.T) {
	q := Qris{}
	err := q.Parse(testData)
	assert.Equal(t, q.MerchantAccountInformationDomestic.ReverseDomain, "COM.XXXX.WWW")
	assert.Equal(t, q.MerchantAccountInformationDomestic.GlobalID, "120000123400001234")
	assert.Equal(t, q.MerchantAccountInformationDomestic.ID, "1234")
	assert.Equal(t, q.MerchantAccountInformationDomestic.Type, "UMI")
	assert.Equal(t, q.MerchantName, "TOKO MIXUE BARU")
	assert.Equal(t, q.MerchantCity, "Kabupaten Tokyo")
	assert.Equal(t, q.TransactionAmount, "10000")
	assert.Nil(t, err)

	log.Printf("%+v\n", q)
}

func TestQrisParseNoAmount(t *testing.T) {
	q := Qris{}
	err := q.Parse(testDataNoAmount)
	assert.Equal(t, q.MerchantAccountInformationDomestic.ReverseDomain, "COM.MYWEB.WWW")
	assert.Equal(t, q.MerchantAccountInformationDomestic.GlobalID, "123456789012345678")
	assert.Equal(t, q.MerchantAccountInformationDomestic.ID, "12345678901234")
	assert.Equal(t, q.MerchantAccountInformationDomestic.Type, "UKE")
	assert.Equal(t, q.MerchantName, "QRIS ARYA")
	assert.Equal(t, q.MerchantCity, "Jakarta Selatan")
	assert.Equal(t, q.TransactionAmount, "")
	assert.Nil(t, err)

	log.Printf("%+v\n", q)
}
