package bitmap

// 位图
type Bitmap struct {
	data    []byte // 保存bit数据
	bitsize uint64 // bit容量
	maxpos  uint64 // 被设置为1的最大位数
}

// SetBit 将 offset 位置的 bit 置为 value (0/1)
func (this *Bitmap) SetBit(offset uint64, value uint8) bool {
	if this.bitsize < offset {
		return false
	}

	index, pos := offset/8, offset%8
	if value == 0 {
		// &^ 清位
		this.data[index] &^= 0x01 << pos
	} else {
		this.data[index] |= 0x01 << pos

		// 记录曾经设置为 1 的最大位置
		if this.maxpos < offset {
			this.maxpos = offset
		}
	}
	return true
}

func (this *Bitmap) Range() []uint64 {
	res := []uint64{}
	var offset uint64
	for offset < this.maxpos {
		index, pos := offset/8, offset%8
		if this.data[index]&0x01<<pos != 0 {
			res = append(res, offset)
		}
	}
	return res
}
