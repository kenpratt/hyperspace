package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GameConstants) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ship_acceleration":
			z.ShipAcceleration, err = dc.ReadUint16()
			if err != nil {
				return
			}
		case "ship_rotation":
			z.ShipRotation, err = dc.ReadUint16()
			if err != nil {
				return
			}
		case "projectile_speed":
			z.ProjectileSpeed, err = dc.ReadUint16()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z GameConstants) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}
	err = en.WriteString("ship_acceleration")
	if err != nil {
		return
	}
	err = en.WriteUint16(z.ShipAcceleration)
	if err != nil {
		return
	}
	err = en.WriteString("ship_rotation")
	if err != nil {
		return
	}
	err = en.WriteUint16(z.ShipRotation)
	if err != nil {
		return
	}
	err = en.WriteString("projectile_speed")
	if err != nil {
		return
	}
	err = en.WriteUint16(z.ProjectileSpeed)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z GameConstants) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 3)
	o = msgp.AppendString(o, "ship_acceleration")
	o = msgp.AppendUint16(o, z.ShipAcceleration)
	o = msgp.AppendString(o, "ship_rotation")
	o = msgp.AppendUint16(o, z.ShipRotation)
	o = msgp.AppendString(o, "projectile_speed")
	o = msgp.AppendUint16(o, z.ProjectileSpeed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameConstants) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ship_acceleration":
			z.ShipAcceleration, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "ship_rotation":
			z.ShipRotation, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "projectile_speed":
			z.ProjectileSpeed, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z GameConstants) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 17 + msgp.Uint16Size + msgp.StringPrefixSize + 13 + msgp.Uint16Size + msgp.StringPrefixSize + 16 + msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GameError) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "What":
			z.What, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z GameError) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(1)
	if err != nil {
		return
	}
	err = en.WriteString("What")
	if err != nil {
		return
	}
	err = en.WriteString(z.What)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z GameError) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 1)
	o = msgp.AppendString(o, "What")
	o = msgp.AppendString(o, z.What)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameError) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "What":
			z.What, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z GameError) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.StringPrefixSize + len(z.What)
	return
}
