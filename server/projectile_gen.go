package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Projectile) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "i":
			z.Id, err = dc.ReadString()
			if err != nil {
				return
			}
		case "z":
			z.Alive, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "p":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Position = nil
			} else {
				if z.Position == nil {
					z.Position = new(Point)
				}
				err = z.Position.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "v":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Velocity = nil
			} else {
				if z.Velocity == nil {
					z.Velocity = new(Vector)
				}
				err = z.Velocity.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "c":
			z.Created, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "o":
			z.Owner, err = dc.ReadString()
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
func (z *Projectile) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(6)
	if err != nil {
		return
	}
	err = en.WriteString("i")
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		return
	}
	err = en.WriteString("z")
	if err != nil {
		return
	}
	err = en.WriteBool(z.Alive)
	if err != nil {
		return
	}
	err = en.WriteString("p")
	if err != nil {
		return
	}
	if z.Position == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Position.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteString("v")
	if err != nil {
		return
	}
	if z.Velocity == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Velocity.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteString("c")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Created)
	if err != nil {
		return
	}
	err = en.WriteString("o")
	if err != nil {
		return
	}
	err = en.WriteString(z.Owner)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Projectile) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 6)
	o = msgp.AppendString(o, "i")
	o = msgp.AppendString(o, z.Id)
	o = msgp.AppendString(o, "z")
	o = msgp.AppendBool(o, z.Alive)
	o = msgp.AppendString(o, "p")
	if z.Position == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Position.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendString(o, "v")
	if z.Velocity == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Velocity.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendString(o, "c")
	o = msgp.AppendUint64(o, z.Created)
	o = msgp.AppendString(o, "o")
	o = msgp.AppendString(o, z.Owner)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Projectile) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "i":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "z":
			z.Alive, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "p":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Position = nil
			} else {
				if z.Position == nil {
					z.Position = new(Point)
				}
				bts, err = z.Position.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "v":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Velocity = nil
			} else {
				if z.Velocity == nil {
					z.Velocity = new(Vector)
				}
				bts, err = z.Velocity.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "c":
			z.Created, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "o":
			z.Owner, bts, err = msgp.ReadStringBytes(bts)
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

func (z *Projectile) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 1 + msgp.StringPrefixSize + len(z.Id) + msgp.StringPrefixSize + 1 + msgp.BoolSize + msgp.StringPrefixSize + 1
	if z.Position == nil {
		s += msgp.NilSize
	} else {
		s += z.Position.Msgsize()
	}
	s += msgp.StringPrefixSize + 1
	if z.Velocity == nil {
		s += msgp.NilSize
	} else {
		s += z.Velocity.Msgsize()
	}
	s += msgp.StringPrefixSize + 1 + msgp.Uint64Size + msgp.StringPrefixSize + 1 + msgp.StringPrefixSize + len(z.Owner)
	return
}
