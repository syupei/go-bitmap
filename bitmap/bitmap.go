/*
* Copyright (c) 2013 Thordur Bjornsson <thorduri@secnorth.net>
*
* Permission to use, copy, modify, and distribute this software for any
* purpose with or without fee is hereby granted, provided that the above
* copyright notice and this permission notice appear in all copies.
*
* THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
* WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
* MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
* ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
* WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
* ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
* OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
*/

// Package bitmap provides basic bit twiddling functionality for bitmaps.
// The package has no specific endianess.
package bitmap

const byteWidth = 8
const byteShift = 3

// Set sets bit in bitmap.
func Set(bit uint, bitmap []byte) {
	bitmap[bit >> byteShift] |= (1 << (bit & (byteWidth - 1)))
}

// SetAll sets all bits in the bitmap.
func SetAll(bitmap []byte) {
	for i, _ := range bitmap {
		bitmap[i] |= 0xFF
	}
}

// Clr clears bit in bitmap.
func Clr(bit uint, bitmap[]byte) {
	bitmap[bit >> byteShift] &= ^(1 << (bit & (byteWidth - 1)))
}

// CLrAll clears all bits in the bitmap.
func ClrAll(bitmap []byte) {
	for i, _ := range bitmap {
		bitmap[i] = 0
	}
}

// IsSet tests if bit is set in bitmap.
func IsSet(bit uint, bitmap []byte) (bool) {
	return bitmap[bit >> byteShift] & (1 << (bit & (byteWidth - 1))) != 0
}

// Areset tests if bits are set in bitmap.
func AreSet(bitmap []byte, bits ...uint) (bool) {
	for _, bit := range bits {
		if IsSet(bit, bitmap) == false {
			return false
		}
	}

	return true
}
