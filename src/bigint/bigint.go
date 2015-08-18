package bigint

import (
	"strconv"
	"strings"
)

type BigInt struct {
	Digits []int
}

func (b BigInt) String() string {
	str_digits := make([]string, 0)
	for i := len(b.Digits) - 1; i >= 0; i-- {
		str_digits = append(str_digits, strconv.Itoa(b.Digits[i]))
	}
	return strings.Join(str_digits, "")
}

func (b BigInt) Copy() BigInt {
    cp_digits := make([]int, len(b.Digits))
    copy(cp_digits, b.Digits)
    return BigInt{cp_digits}
}

func (self *BigInt) Add(oth BigInt) {
    carry := 0
    for ind, my_val := range self.Digits {
        var oth_val int
        if ind < len(oth.Digits) {
            oth_val = oth.Digits[ind]
        } else {
            oth_val = 0
        }
        new_val := my_val + oth_val + carry
        self.Digits[ind] = new_val % 10
        carry = new_val / 10
    }
    if len(self.Digits) < len(oth.Digits) {
        for i := len(self.Digits); i < len(oth.Digits); i++ {
            new_val := oth.Digits[i] + carry
            self.Digits = append(self.Digits, new_val % 10)
            carry = new_val / 10
        }
    }
    for carry > 0 {
        self.Digits = append(self.Digits, carry%10)
        carry /= 10
    }
}

func (b *BigInt) Mult(n int) {
	carry := 0
	for ind, val := range b.Digits {
		new_val := n*val + carry
		b.Digits[ind] = new_val % 10
		carry = new_val / 10
	}
	for carry > 0 {
		b.Digits = append(b.Digits, carry%10)
		carry /= 10
	}
}

func (b BigInt) DigitSum() int {
	sum := 0
	for _, val := range b.Digits {
		sum += val
	}
	return sum
}
