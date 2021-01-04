package circular

import "fmt"

// Buffer represents a circular buffer
type Buffer struct {
	content []byte
	size    int
}

// NewBuffer creates a circular buffer
func NewBuffer(size int) *Buffer {
	return &Buffer{size: size}
}

// WriteByte writes a byte in the circular buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.full() {
		return fmt.Errorf("cannot write to a full buffer")
	}
	b.content = append(b.content, c)
	return nil
}

// ReadByte reads a byte from the circular buffer
func (b *Buffer) ReadByte() (byte, error) {
	if b.empty() {
		return 0, fmt.Errorf("cannot read from an empty buffer")
	}
	r := b.content[0]
	b.content = b.content[1:]
	return r, nil
}

// Overwrite writes a byte in the circular buffer, overwriting if the buffer is full
func (b *Buffer) Overwrite(c byte) {
	if b.full() {
		b.content = b.content[1:]
	}
	b.content = append(b.content, c)
}

// Reset deletes the content from the circular buffer
func (b *Buffer) Reset() {
	b.content = []byte{}
}

func (b *Buffer) empty() bool {
	return len(b.content) == 0
}

func (b *Buffer) full() bool {
	return len(b.content) == b.size
}
