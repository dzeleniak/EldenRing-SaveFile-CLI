package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"unicode/utf16"
)

const (
	slot1_start  = 0x00000310
	slot2_start  = 0x00280320
	slot3_start  = 0x500330
	slot4_start  = 0x780340
	slot5_start  = 0xA00350
	slot6_start  = 0xC80360
	slot7_start  = 0xF00370
	slot8_start  = 0x1180380
	slot9_start  = 0x1400390
	slot10_start = 0x16803A0

	slot1_end  = 0x0028030F + 1
	slot2_end  = 0x050031F + 1
	slot3_end  = 0x78032F + 1
	slot4_end  = 0xA0033F + 1
	slot5_end  = 0xC8034F + 1
	slot6_end  = 0xF0035F + 1
	slot7_end  = 0x118036F + 1
	slot8_end  = 0x140037F + 1
	slot9_end  = 0x168038F + 1
	slot10_end = 0x190039F + 1

	slot1_name_start  = 0x1901d0e
	slot2_name_start  = 0x1901f5a
	slot3_name_start  = 0x19021a6
	slot4_name_start  = 0x19023f2
	slot5_name_start  = 0x190263e
	slot6_name_start  = 0x190288a
	slot7_name_start  = 0x1902ad6
	slot8_name_start  = 0x1902d22
	slot9_name_start  = 0x1902f6e
	slot10_name_start = 0x19031ba
	nameLength        = 32
)

var (
	slotLocations = [][]int{
		{slot1_start, slot1_end},
		{slot2_start, slot2_end},
		{slot3_start, slot3_end},
		{slot4_start, slot4_end},
		{slot5_start, slot5_end},
		{slot6_start, slot6_end},
		{slot7_start, slot7_end},
		{slot8_start, slot8_end},
		{slot9_start, slot9_end},
		{slot10_start, slot10_end},
	}
	slotNameLocations = []int{
		slot1_name_start,
		slot2_name_start,
		slot3_name_start,
		slot4_name_start,
		slot5_name_start,
		slot6_name_start,
		slot7_name_start,
		slot8_name_start,
		slot9_name_start,
		slot10_name_start,
	}
)

type saveFile struct {
	filename   string
	data       []byte
	characters map[string]*character
}

type SaveFile interface {
	Load()
	LoadCharacters() error
}

func (f *saveFile) Load() error {
	file, err := os.ReadFile(f.filename)
	if err != nil {
		return err
	}

	f.data = file

	return nil
}

func (f *saveFile) getCharacterName(slot int) (string, error) {
	nameBytes := f.data[slotNameLocations[slot] : slotNameLocations[slot]+nameLength]
	nameHex := make([]uint16, len(nameBytes)/2)

	if err := binary.Read(bytes.NewReader(nameBytes), binary.LittleEndian, &nameHex); err != nil {
		return "", err
	}

	return string(utf16.Decode(nameHex)), nil
}

func (f *saveFile) LoadCharacters() error {
	for i, v := range slotLocations {
		data := f.data[v[0]:v[1]]
		name, err := f.getCharacterName(i)

		if err != nil {
			return err
		}

		c := &character{
			data: data,
			name: name,
		}

		f.characters[name] = c
		c.Load()
	}

	return nil
}
