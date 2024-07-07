package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// lEndian takes a byte slice and returns the little endian int32/64 representation.
func lEndian(val []byte) uint64 {
	// Reverse the byte slice to represent little endian
	for i, j := 0, len(val)-1; i < j; i, j = i+1, j-1 {
		val[i], val[j] = val[j], val[i]
	}

	// Convert reversed byte slice to integer
	buf := bytes.NewReader(val)
	var intValue uint64
	err := binary.Read(buf, binary.BigEndian, &intValue)
	if err != nil {
		return 0
	}

	return intValue
}

func split(data []byte, chunkSize int) [][]byte {
	var splitted [][]byte
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		splitted = append(splitted, chunk)
	}
	return splitted
}

// getIdReversed takes the first 4 bytes of the chunk, reverses them, and converts to a hexadecimal string
func getIdReversed(id []byte) string {
	if len(id) < 4 {
		return "" // or handle the error appropriately
	}
	tmp := make([]byte, 4)
	copy(tmp, id[:4])
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	finalID := ""
	for _, b := range tmp {
		finalID += decimalToHex(b, 2)
	}
	return finalID
}

// decimalToHex converts a byte to a hexadecimal string with minimum length
func decimalToHex(b byte, length int) string {
	return fmt.Sprintf("%0*x", length, b)
}

func subfinder(data []byte, pattern []byte) int {
	patternLen := len(pattern)

	for i := 0; i <= len(data)-patternLen; i++ {
		if data[i] == pattern[0] && bytes.Equal(data[i:i+patternLen], pattern) {
			return i
		}
	}

	return -1 // Return -1 if pattern is not found
}
