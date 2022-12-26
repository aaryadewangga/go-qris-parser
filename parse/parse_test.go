package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testData         = "00020101021226530012COM.XXXX.WWW0118936004940000007520020475200303UMI51380014ID.CO.QRIS.WWW02091234567890303UMI5204723053033605405100005802ID5915TOKO MAINAN HIL6015Kabupaten Bangk6106123131622001020102034670703K19630433CE"
	testDataNoAmount = "00020101021226640013COM.MYWEB.WWW01181234567890123456780214123456789012340303UKE5912QRIS WANTUNO6013Jakarta Pusat6304XXXX"
)

func TestQrisParseNoAmount(t *testing.T) {
	res := QrisParser(testDataNoAmount)
	assert.Equal(t, res.QRISAcquirerMerchantID, "12345678901234")
	assert.Equal(t, res.TransactionAmount, "")
}

func TestQrisParseWithAmount(t *testing.T) {
	res := QrisParser(testData)
	assert.Equal(t, res.MerchantCity, "Kabupaten Bangk")
	assert.Equal(t, res.TransactionAmount, "10000")
}
