package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Point) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "x":
			z.X, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "y":
			z.Y, err = dc.ReadFloat64()
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
func (z Point) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("x")
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.X)
	if err != nil {
		return
	}
	err = en.WriteString("y")
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Y)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Point) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "x")
	o = msgp.AppendFloat64(o, z.X)
	o = msgp.AppendString(o, "y")
	o = msgp.AppendFloat64(o, z.Y)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Point) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "x":
			z.X, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "y":
			z.Y, bts, err = msgp.ReadFloat64Bytes(bts)
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

func (z Point) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 1 + msgp.Float64Size + msgp.StringPrefixSize + 1 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Vector) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "x":
			z.X, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "y":
			z.Y, err = dc.ReadFloat64()
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
func (z Vector) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("x")
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.X)
	if err != nil {
		return
	}
	err = en.WriteString("y")
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Y)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Vector) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "x")
	o = msgp.AppendFloat64(o, z.X)
	o = msgp.AppendString(o, "y")
	o = msgp.AppendFloat64(o, z.Y)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Vector) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "x":
			z.X, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "y":
			z.Y, bts, err = msgp.ReadFloat64Bytes(bts)
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

func (z Vector) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 1 + msgp.Float64Size + msgp.StringPrefixSize + 1 + msgp.Float64Size
	return
}
