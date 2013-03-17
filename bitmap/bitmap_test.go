package bitmap

import (
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
}

func BenchmarkSet(b *testing.B) {
	bitmap := make([]byte, 524288)

	for i := 0; i < b.N; i++ {
		bit := uint(b.N % 524288)
		Set(bit, bitmap)
	}
}

func TestSetAll(t *testing.T) {
}

func BenchmarkSetAll(b *testing.B) {
	bitmap := make([]byte, 524288)

	for i := 0; i < b.N; i++ {
		SetAll(bitmap)
	}
}

func TestClr(t *testing.T) {
}

func BenchmarkClr(b *testing.B) {
	bitmap := make([]byte, 524288)

	for i := 0; i < b.N; i++ {
		bit := uint(b.N % 524288)
		Clr(bit, bitmap)
	}
}

func TestClrAll(t *testing.T) {
}

func BenchmarkClrAll(b *testing.B) {
	bitmap := make([]byte, 524288)

	for i := 0; i < b.N; i++ {
		ClrAll(bitmap)
	}
}

func TestIsSet(b *testing.T) {
}

func BenchmarkIsSet(b *testing.B) {
	bitmap := make([]byte, 524288)

	for i := 0; i < b.N; i++ {
		bit := uint(b.N % 524288)
		IsSet(bit, bitmap)
	}
}

func TestAreSet(t *testing.T) {
	var b0, b1, b2, b3, b4, b5, b6, b7 uint
	bitmap := make([]byte, 524288)

	bits := make([]uint, 256)
	for i := 0; i < 256; i++ {
		bits[i] = uint(rand.Intn(524288))
		Set(bits[i], bitmap)
	}

	for i := 0; i < 256; i++ {
		b0 = uint(rand.Intn(256))
		b1 = uint(rand.Intn(256))
		b2 = uint(rand.Intn(256))
		b3 = uint(rand.Intn(256))
		b4 = uint(rand.Intn(256))
		b5 = uint(rand.Intn(256))
		b6 = uint(rand.Intn(256))
		b7 = uint(rand.Intn(256))

		if AreSet(bitmap, bits[b0]) != true {
			t.Errorf("bit %v should be set in bitmap", bits[:1])
		}
		if AreSet(bitmap, bits[b0], bits[b1]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:2])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:3])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:4])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:5])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:6])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5], bits[b6]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:7])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5], bits[b6], bits[b7]) != true {
			t.Errorf("bits %v should be set in bitmap", bits[:8])
		}
	}

	ClrAll(bitmap)

	for i := 0; i < 256; i++ {
		b0 = uint(rand.Intn(256))
		b1 = uint(rand.Intn(256))
		b2 = uint(rand.Intn(256))
		b3 = uint(rand.Intn(256))
		b4 = uint(rand.Intn(256))
		b5 = uint(rand.Intn(256))
		b6 = uint(rand.Intn(256))
		b7 = uint(rand.Intn(256))

		if AreSet(bitmap, bits[b0]) != false {
			t.Errorf("bit %v should not be set in bitmap", bits[:1])
		}
		if AreSet(bitmap, bits[b0], bits[b1]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:2])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:3])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:4])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:5])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:6])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5], bits[b6]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:7])
		}
		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5], bits[b6], bits[b7]) != false {
			t.Errorf("bits %v should not be set in bitmap", bits[:8])
		}
	}

}

// Benchmark checking for 8 bits set.
func BenchmarkAreSet(b *testing.B) {
	var b0, b1, b2, b3, b4, b5, b6, b7 uint
	bitmap := make([]byte, 524288)

	bits := make([]uint, 8)
	for i := 0; i < 8; i++ {
		bits[i] = uint(rand.Intn(524288))
		Set(bits[i], bitmap)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b0 = uint(rand.Intn(8))
		b1 = uint(rand.Intn(8))
		b2 = uint(rand.Intn(8))
		b3 = uint(rand.Intn(8))
		b4 = uint(rand.Intn(8))
		b5 = uint(rand.Intn(8))
		b6 = uint(rand.Intn(8))
		b7 = uint(rand.Intn(8))

		if AreSet(bitmap, bits[b0], bits[b1], bits[b2], bits[b3], bits[b4], bits[b5], bits[b6], bits[b7]) != true {
			b.Errorf("bits %v should be set in bitmap", bits[:8])
		}
	}
}
