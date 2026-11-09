package qris

import (
	"testing"
)

const (
	testData         = "00020101021226530012COM.XXXX.WWW0118120000123400001234020412340303UMI51380014ID.CO.QRIS.WWW02091234567890303UMI5204723053033605405100005802ID5915TOKO MIXUE BARU6015Kabupaten Tokyo6106123131622001MIXUE BARU70703K19630433CE"
	testDataNoAmount = "00020101021226640013COM.MYWEB.WWW01181234567890123456780214123456789012340303UKE5909QRIS ARYA6015Jakarta Selatan6304XXXX"
)

func TestQrisParseWithAmount(t *testing.T) {
	q := Qris{}
	if err := q.Parse(testData); err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	assertEqual(t, "MerchantAccountInformationDomestic.ReverseDomain", q.MerchantAccountInformationDomestic.ReverseDomain, "COM.XXXX.WWW")
	assertEqual(t, "MerchantAccountInformationDomestic.GlobalID", q.MerchantAccountInformationDomestic.GlobalID, "120000123400001234")
	assertEqual(t, "MerchantAccountInformationDomestic.ID", q.MerchantAccountInformationDomestic.ID, "1234")
	assertEqual(t, "MerchantAccountInformationDomestic.Type", q.MerchantAccountInformationDomestic.Type, "UMI")
	assertEqual(t, "MerchantName", q.MerchantName, "TOKO MIXUE BARU")
	assertEqual(t, "MerchantCity", q.MerchantCity, "Kabupaten Tokyo")
	assertEqual(t, "TransactionAmount", q.TransactionAmount, "10000")
}

func TestQrisParseNoAmount(t *testing.T) {
	q := Qris{}
	if err := q.Parse(testDataNoAmount); err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	assertEqual(t, "MerchantAccountInformationDomestic.ReverseDomain", q.MerchantAccountInformationDomestic.ReverseDomain, "COM.MYWEB.WWW")
	assertEqual(t, "MerchantAccountInformationDomestic.GlobalID", q.MerchantAccountInformationDomestic.GlobalID, "123456789012345678")
	assertEqual(t, "MerchantAccountInformationDomestic.ID", q.MerchantAccountInformationDomestic.ID, "12345678901234")
	assertEqual(t, "MerchantAccountInformationDomestic.Type", q.MerchantAccountInformationDomestic.Type, "UKE")
	assertEqual(t, "MerchantName", q.MerchantName, "QRIS ARYA")
	assertEqual(t, "MerchantCity", q.MerchantCity, "Jakarta Selatan")
	assertEqual(t, "TransactionAmount", q.TransactionAmount, "")
}

func assertEqual(t *testing.T, field, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %q, want %q", field, got, want)
	}
}
