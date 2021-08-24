package v044

var (
	DenomAddressPrefix = []byte{0x03}
)

func CreateAddressDenomPrefix(denom string) []byte {
	key := DenomAddressPrefix
	key = append(key, []byte(denom)...)
	return append(key, 0)
}
