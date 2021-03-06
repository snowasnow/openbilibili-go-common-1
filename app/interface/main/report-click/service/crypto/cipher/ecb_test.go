package cipher

import (
	iaes "crypto/aes"
	"testing"

	"go-common/app/interface/main/report-click/service/crypto/padding"

	. "github.com/smartystreets/goconvey/convey"
)

func TestECB(t *testing.T) {
	var (
		aesKey = "903ef9925be1300f843b41df2ca55410"
		bs     = []byte("this is  test massage ")
	)
	Convey("cipher test ", t, func() {
		var p = padding.PKCS5
		bs = p.Padding(bs, iaes.BlockSize)
		b, _ := iaes.NewCipher([]byte(aesKey))
		ecbe := NewECBEncrypter(b)
		encryptText := make([]byte, len(bs))
		ecbe.CryptBlocks(encryptText, bs)
		ecbd := NewECBDecrypter(b)
		decryptText := make([]byte, len(bs))
		ecbd.CryptBlocks(decryptText, bs)
		p.Unpadding(bs, iaes.BlockSize)
	})
}
