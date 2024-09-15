package main

import "fmt"

// CompressionStrategy defines the interface for compression algorithms.
type CompressionStrategy interface {
	Compress(data []byte) []byte
}

// ZipCompression is one concrete strategy.
type ZipCompression struct{}

func (z ZipCompression) Compress(data []byte) []byte {
	fmt.Println("Compressing using ZIP strategy")
	return data // Placeholder implementation
}

// RarCompression is another concrete strategy.
type RarCompression struct{}

func (r RarCompression) Compress(data []byte) []byte {
	fmt.Println("Compressing using RAR strategy")
	return data // Placeholder implementation
}

// Compressor uses a CompressionStrategy.
type Compressor struct {
	strategy CompressionStrategy
}

func (c *Compressor) SetStrategy(strategy CompressionStrategy) {
	c.strategy = strategy
}

func (c *Compressor) CompressData(data []byte) []byte {
	return c.strategy.Compress(data)
}

func main() {
	data := []byte("Example data")

	compressor := &Compressor{}

	compressor.SetStrategy(ZipCompression{})
	compressor.CompressData(data)

	compressor.SetStrategy(RarCompression{})
	compressor.CompressData(data)
}
