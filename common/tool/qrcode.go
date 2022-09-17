package tool

import qrcode "github.com/skip2/go-qrcode"

func MakeQrcode(con string) (png []byte, err error) {
	png, err = qrcode.Encode(con, qrcode.Medium, 256)
	return
}
