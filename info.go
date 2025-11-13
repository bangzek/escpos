package escpos

const (
	INFO_HEADER = '_'
)

type Info []byte

func MakeInfo(b []byte) Info {
	n := make([]byte, len(b))
	copy(n, b)
	return Info(n)
}

func (i Info) IsValid() bool {
	return len(i) >= 2 && i[0] == INFO_HEADER && i[len(i)-1] == NUL
}

func (i Info) String() string {
	if i.IsValid() {
		return string(i[1 : len(i)-1])
	} else {
		return ""
	}
}
