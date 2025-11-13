package escpos_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/bangzek/escpos"
)

var _ = Describe("Parser", func() {
	It("convert esc/pos text to raw string", func() {
		text := `' this is comment
NUL SOH` + "\t" + `STX  ETX
EOT    ENQ` + "\r\n" + `
    ACK BEL BS HT LF VT FF CR SO SI
DLE DC1 DC2 DC3 DC4 NAK SYN ETB CAN EM SUB ESC FS GS RS US DEL
"Testing do" SP "ang" LF
	"Apakah " DQ "ok" DQ " atau " SQ "tidak" SQ LF
			0 1 2 3 4 5 6 7 8 9 10 11 99 123 234 255 LF
			0x00 0xFF 0xBE 0xEF 0x1A
`
		b := []byte{
			0, 1, 2, 3, 4, 5, 6, 7,
			8, 9, 10, 11, 12, 13, 14, 15,
			16, 17, 18, 19, 20, 21, 22, 23,
			24, 25, 26, 27, 28, 29, 30, 31,
			127, 'T', 'e', 's', 't', 'i', 'n', 'g', ' ',
			'd', 'o', ' ', 'a', 'n', 'g', '\n',
			'A', 'p', 'a', 'k', 'a', 'h', ' ', '"', 'o', 'k', '"', ' ',
			'a', 't', 'a', 'u', ' ', '\'', 't', 'i', 'd', 'a', 'k', '\'', '\n',
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 99, 123, 234, 255, '\n',
			0, 0xFF, 0xBE, 0xEF, 0x1A,
		}
		Expect(Parse([]byte(text))).To(Equal(b))
	})
})
