//
// Used to Generate UUID. It's standard by RFC4122
// Current package use VariantRFC4122 as Default Variant
//
// Copyright (C) 2013 by Hydra <xxeaglenet@gmail.com>
//
// [https://github.com/eahydra/gouuid]
package uuid

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
	"net"
	"time"
)

// The variant field determines the layout of the UUID.
const (
	VariantNCS       byte = 0x80 // Reserved, NCS backward compatibility.
	VariantRFC4122   byte = 0x40 // The variant specified in RFC4122
	VariantMicrosoft byte = 0x20 // Reserved, Microsoft Corporation backward compatibility
	VariantReserved  byte = 0x00 // Reserved for future definition.
)

// The UUID Generate Algorithm Version
const (
	VerTimeStamp     byte = 1 // The time-based version
	VerNameBasedMD5  byte = 3 // The name-based version by MD5
	VerRandom        byte = 4 // The randomly or pseudorandomly generated version
	VerNameBasedSHA1 byte = 5 // The name-based version by SHA1
)

// Some Reserved Namespace UUID
var (
	NamespaceDNS  = &UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NamespaceURL  = &UUID{0x6b, 0xa7, 0xb8, 0x11, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NamespaceOID  = &UUID{0x6b, 0xa7, 0xb8, 0x12, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	NamespaceX500 = &UUID{0x6b, 0xa7, 0xb8, 0x14, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
)

var cacheNodeID []byte

type UUID [16]byte

// Convert UUID to String, format is {x-x-x-x-x}
func (u *UUID) String() string {
	return fmt.Sprintf("{%x-%x-%x-%x-%x}", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// Convert UUID to byte slice, it's Little Endian
func (u *UUID) Byte() []byte {
	return u[:]
}

// Get Gernerate Version
func (u *UUID) Version() byte {
	return (u[6] & 0xF0) >> 4
}

// Get Generate Variant
func (u *UUID) Variant() byte {
	d := ^u[8]
	if d&VariantNCS != 0 {
		return VariantNCS
	} else if d&VariantRFC4122 != 0 {
		return VariantRFC4122
	} else if d&VariantMicrosoft != 0 {
		return VariantMicrosoft
	}
	return VariantReserved
}

func (u *UUID) setVariantAndVersion(variant byte, ver byte) {
	switch variant {
	case VariantNCS:
		u[8] &= 0x7F
	case VariantRFC4122:
		u[8] &= 0x3F
		u[8] |= 0x80
	case VariantMicrosoft:
		u[8] &= 0x1F
		u[8] |= 0xC0
	case VariantReserved:
		u[8] &= 0x1F
		u[8] |= 0xE0
	}
	u[6] = (u[6] & 0x0F) | (ver << 4)
}

// Compare u1 and u2 whether equ or less
func Compare(u1 *UUID, u2 *UUID) int {
	return bytes.Compare(u1.Byte(), u2.Byte())
}

// Parse from byte slice to UUID
func Parse(b []byte) *UUID {
	if len(b) != 16 {
		return nil
	}
	var u UUID
	copy(u[:], b[0:16])
	return &u
}

// Generate one UUID standard by RFC4122
func NewUUID() *UUID {
	return NewUUIDByTime()
}

// Generate with timestamp.
func NewUUIDByTime() *UUID {
	t := uint64(time.Now().UnixNano()/100 + 0x01b21dd213814000)
	var b [2]byte
	rand.Read(b[:])
	clockSeq := binary.LittleEndian.Uint16(b[:])

	timeLow := uint32(t & (0x100000000 - 1))
	timeMid := uint16((t >> 32) & 0xFFFF)
	timeHighVer := uint16((t >> 48) & 0x0FFF)
	clockSeq &= 0x3FFF

	var u UUID
	binary.LittleEndian.PutUint32(u[0:4], timeLow)
	binary.LittleEndian.PutUint16(u[4:6], timeMid)
	binary.LittleEndian.PutUint16(u[6:8], timeHighVer)
	binary.LittleEndian.PutUint16(u[8:10], clockSeq)
	copy(u[10:16], cacheNodeID)
	u.setVariantAndVersion(VariantRFC4122, VerTimeStamp)
	return &u
}

func getNodeID() []byte {
	var d [6]byte
	inet, err := net.Interfaces()
	if err == nil {
		var set bool
		for _, v := range inet {
			if len(v.HardwareAddr.String()) != 0 {
				copy(d[:], []byte(v.HardwareAddr))
				set = true
				break
			}
		}
		if set {
			return d[:]
		}
	}
	rand.Read(d[:])
	d[0] |= 0x01
	return d[:]
}

func hashUUID(h hash.Hash, un *UUID, name string) []byte {
	h.Write(un[:])
	h.Write([]byte(name))
	return h.Sum(nil)
}

// Generate base namespace with MD5
func NewUUIDByMd5(un *UUID, name string) *UUID {
	b := hashUUID(md5.New(), un, name)
	var u UUID
	copy(u[:], b[:])
	u.setVariantAndVersion(VariantRFC4122, VerNameBasedMD5)
	return &u
}

// Generate base namespace with SHA1
func NewUUIDBySHA1(un *UUID, name string) *UUID {
	b := hashUUID(sha1.New(), un, name)
	var u UUID
	copy(u[:], b[:16])
	u.setVariantAndVersion(VariantRFC4122, VerNameBasedSHA1)
	return &u
}

// Generate UUI with random
func NewUUIDByRandom() *UUID {
	var u UUID
	rand.Read(u[:])
	u.setVariantAndVersion(VariantRFC4122, VerRandom)
	return &u
}

func init() {
	cacheNodeID = getNodeID()
}
