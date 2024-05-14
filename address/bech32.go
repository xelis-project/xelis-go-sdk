package address

import (
	"errors"
	"fmt"
	"strings"
)

var CHARSET = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
var GENERATOR = [5]uint32{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
var SEPARATOR = ":"

type Bech32Error error

var ErrHrpMixCase Bech32Error = errors.New("mix case is not allowed in human readable part")
var ErrInvalidChecksum Bech32Error = errors.New("invalid checksum")
var ErrNonZeroPadding Bech32Error = errors.New("non zero padding")
var ErrIllegalZeroPadding Bech32Error = errors.New("illegal zero padding")
var ErrHrpEmpty Bech32Error = errors.New("human readable part is empty")

func ErrSeparatorInvalidPosition(pos int) Bech32Error {
	return fmt.Errorf("invalid separator position: %d", pos)
}

func ErrHrpInvalidCharacter(c byte) Bech32Error {
	return fmt.Errorf("invalid character value in human readable part: %d", c)
}

func ErrInvalidIndex(index int) Bech32Error {
	return fmt.Errorf("invalid index %d", index)
}

func ErrInvalidValue(value byte, max int) Bech32Error {
	return fmt.Errorf("invalid value: %d, max is %d", value, max)
}

func polymod(values []byte) uint32 {
	chk := uint32(1)
	for _, value := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ uint32(value)
		for i, item := range GENERATOR {
			if (top>>i)&1 == 1 {
				chk ^= item
			}
		}
	}

	return chk
}

func hrpExpand(hrp string) (result []byte) {
	for _, value := range []byte(hrp) {
		result = append(result, value>>5)
	}
	result = append(result, 0)
	for _, value := range []byte(hrp) {
		result = append(result, value&31)
	}

	return
}

func verifyChecksum(hrp string, data []byte) bool {
	values := hrpExpand(hrp)
	values = append(values, data...)
	return polymod(values) == 1
}

func createChecksum(hrp string, data []byte) (result []byte) {
	var values []byte
	values = append(values, hrpExpand(hrp)...)
	values = append(values, data...)
	result = make([]byte, 6)
	values = append(values, result...)
	pm := polymod(values) ^ 1

	for i := 0; i < 6; i++ {
		result[i] = byte(pm >> (5 * (5 - i)) & 31)
	}

	return
}

func convertBits(data []byte, from uint16, to uint16, pad bool) (result []byte, err Bech32Error) {
	acc := uint16(0)
	bits := uint16(0)
	maxValue := uint16((1 << to) - 1)
	for _, v := range data {
		value := uint16(v)
		if value>>from != 0 {
			return
		}

		acc = (acc << from) | value
		bits += from
		for bits >= to {
			bits -= to
			result = append(result, byte((acc>>bits)&maxValue))
		}
	}

	if pad {
		if bits > 0 {
			result = append(result, byte((acc<<(to-bits))&maxValue))
		}
	} else if bits >= from {
		err = ErrIllegalZeroPadding
		return
	} else if (acc<<(to-bits))&maxValue != 0 {
		err = ErrNonZeroPadding
		return
	}

	return
}

func decode(bech string) (hrp string, decoded []byte, err Bech32Error) {
	if strings.ToUpper(bech) != bech && strings.ToLower(bech) != bech {
		err = ErrHrpMixCase
		return
	}

	pos := strings.Index(bech, ":")
	if pos < 1 || pos+7 > len(bech) {
		err = ErrSeparatorInvalidPosition(pos)
		return
	}

	hrp = bech[0:pos]
	for _, value := range []byte(hrp) {
		if value < 33 || value > 126 {
			err = ErrHrpInvalidCharacter(value)
			return
		}
	}

	for i := pos + 1; i < len(bech); i++ {
		if i >= len(bech) {
			err = ErrInvalidIndex(i)
			return
		}

		c := bech[i]
		found := false
		var index byte
		for i, char := range CHARSET {
			if byte(char) == c {
				index = byte(i)
				found = true
				break
			}
		}

		if !found {
			err = ErrHrpInvalidCharacter(c)
			return
		}

		decoded = append(decoded, index)
	}

	if !verifyChecksum(hrp, decoded) {
		err = ErrInvalidChecksum
		return
	}

	decoded = decoded[:len(decoded)-6]

	return
}

func encode(hrp string, data []byte) (str string, err Bech32Error) {
	if len(hrp) == 0 {
		err = ErrHrpEmpty
		return
	}

	for _, value := range []byte(hrp) {
		if value < 33 || value > 126 {
			err = ErrHrpInvalidCharacter(value)
			return
		}
	}

	if strings.ToUpper(hrp) != hrp && strings.ToLower(hrp) != hrp {
		err = ErrHrpMixCase
		return
	}

	hrp = strings.ToLower(hrp)
	var combined []byte
	combined = append(combined, data...)
	combined = append(combined, createChecksum(hrp, data)...)

	var result []byte
	result = append(result, []byte(hrp)...)
	result = append(result, []byte(SEPARATOR)...)

	for _, index := range combined {
		if index > byte(len(CHARSET)) {
			err = ErrInvalidValue(index, len(CHARSET))
			return
		}

		found := false
		var value byte
		for i, char := range CHARSET {
			if i == int(index) {
				value = byte(char)
				found = true
				break
			}
		}

		if !found {
			err = ErrInvalidIndex(int(index))
			return
		}

		result = append(result, value)
	}

	str = string(result)
	return
}
