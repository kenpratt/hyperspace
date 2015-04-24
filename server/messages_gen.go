package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *RotationData) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "eventId":
			z.EventId, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "direction":
			z.Direction, err = dc.ReadInt8()
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
func (z RotationData) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("eventId")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.EventId)
	if err != nil {
		return
	}
	err = en.WriteString("direction")
	if err != nil {
		return
	}
	err = en.WriteInt8(z.Direction)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z RotationData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "eventId")
	o = msgp.AppendUint64(o, z.EventId)
	o = msgp.AppendString(o, "direction")
	o = msgp.AppendInt8(o, z.Direction)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RotationData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "eventId":
			z.EventId, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "direction":
			z.Direction, bts, err = msgp.ReadInt8Bytes(bts)
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

func (z RotationData) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 7 + msgp.Uint64Size + msgp.StringPrefixSize + 9 + msgp.Int8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Message) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "time":
			z.Time, err = dc.ReadUint64()
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
func (z Message) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("type")
	if err != nil {
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	err = en.WriteString("time")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Time)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Message) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "type")
	o = msgp.AppendString(o, z.Type)
	o = msgp.AppendString(o, "time")
	o = msgp.AppendUint64(o, z.Time)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Message) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "time":
			z.Time, bts, err = msgp.ReadUint64Bytes(bts)
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

func (z Message) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.StringPrefixSize + len(z.Type) + msgp.StringPrefixSize + 4 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *InitMessage) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "time":
			z.Time, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "playerId":
			z.PlayerId, err = dc.ReadString()
			if err != nil {
				return
			}
		case "constants":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Constants = nil
			} else {
				if z.Constants == nil {
					z.Constants = new(GameConstants)
				}
				err = z.Constants.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "state":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.State = nil
			} else {
				if z.State == nil {
					z.State = new(GameState)
				}
				err = z.State.DecodeMsg(dc)
				if err != nil {
					return
				}
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
func (z *InitMessage) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(5)
	if err != nil {
		return
	}
	err = en.WriteString("type")
	if err != nil {
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	err = en.WriteString("time")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Time)
	if err != nil {
		return
	}
	err = en.WriteString("playerId")
	if err != nil {
		return
	}
	err = en.WriteString(z.PlayerId)
	if err != nil {
		return
	}
	err = en.WriteString("constants")
	if err != nil {
		return
	}
	if z.Constants == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Constants.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteString("state")
	if err != nil {
		return
	}
	if z.State == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.State.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *InitMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 5)
	o = msgp.AppendString(o, "type")
	o = msgp.AppendString(o, z.Type)
	o = msgp.AppendString(o, "time")
	o = msgp.AppendUint64(o, z.Time)
	o = msgp.AppendString(o, "playerId")
	o = msgp.AppendString(o, z.PlayerId)
	o = msgp.AppendString(o, "constants")
	if z.Constants == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Constants.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendString(o, "state")
	if z.State == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.State.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InitMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "time":
			z.Time, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "playerId":
			z.PlayerId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "constants":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Constants = nil
			} else {
				if z.Constants == nil {
					z.Constants = new(GameConstants)
				}
				bts, err = z.Constants.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "state":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.State = nil
			} else {
				if z.State == nil {
					z.State = new(GameState)
				}
				bts, err = z.State.UnmarshalMsg(bts)
				if err != nil {
					return
				}
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

func (z *InitMessage) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.StringPrefixSize + len(z.Type) + msgp.StringPrefixSize + 4 + msgp.Uint64Size + msgp.StringPrefixSize + 8 + msgp.StringPrefixSize + len(z.PlayerId) + msgp.StringPrefixSize + 9
	if z.Constants == nil {
		s += msgp.NilSize
	} else {
		s += z.Constants.Msgsize()
	}
	s += msgp.StringPrefixSize + 5
	if z.State == nil {
		s += msgp.NilSize
	} else {
		s += z.State.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UpdateMessage) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "time":
			z.Time, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "state":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.State = nil
			} else {
				if z.State == nil {
					z.State = new(GameState)
				}
				err = z.State.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "lastEvent":
			z.LastEventId, err = dc.ReadUint64()
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
func (z *UpdateMessage) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(4)
	if err != nil {
		return
	}
	err = en.WriteString("type")
	if err != nil {
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	err = en.WriteString("time")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Time)
	if err != nil {
		return
	}
	err = en.WriteString("state")
	if err != nil {
		return
	}
	if z.State == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.State.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteString("lastEvent")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.LastEventId)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UpdateMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 4)
	o = msgp.AppendString(o, "type")
	o = msgp.AppendString(o, z.Type)
	o = msgp.AppendString(o, "time")
	o = msgp.AppendUint64(o, z.Time)
	o = msgp.AppendString(o, "state")
	if z.State == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.State.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendString(o, "lastEvent")
	o = msgp.AppendUint64(o, z.LastEventId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UpdateMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "time":
			z.Time, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "state":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.State = nil
			} else {
				if z.State == nil {
					z.State = new(GameState)
				}
				bts, err = z.State.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "lastEvent":
			z.LastEventId, bts, err = msgp.ReadUint64Bytes(bts)
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

func (z *UpdateMessage) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.StringPrefixSize + len(z.Type) + msgp.StringPrefixSize + 4 + msgp.Uint64Size + msgp.StringPrefixSize + 5
	if z.State == nil {
		s += msgp.NilSize
	} else {
		s += z.State.Msgsize()
	}
	s += msgp.StringPrefixSize + 9 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FireData) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "eventId":
			z.EventId, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "projectileId":
			z.ProjectileId, err = dc.ReadString()
			if err != nil {
				return
			}
		case "created":
			z.Created, err = dc.ReadUint64()
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
func (z FireData) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}
	err = en.WriteString("eventId")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.EventId)
	if err != nil {
		return
	}
	err = en.WriteString("projectileId")
	if err != nil {
		return
	}
	err = en.WriteString(z.ProjectileId)
	if err != nil {
		return
	}
	err = en.WriteString("created")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Created)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z FireData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 3)
	o = msgp.AppendString(o, "eventId")
	o = msgp.AppendUint64(o, z.EventId)
	o = msgp.AppendString(o, "projectileId")
	o = msgp.AppendString(o, z.ProjectileId)
	o = msgp.AppendString(o, "created")
	o = msgp.AppendUint64(o, z.Created)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FireData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "eventId":
			z.EventId, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "projectileId":
			z.ProjectileId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadUint64Bytes(bts)
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

func (z FireData) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 7 + msgp.Uint64Size + msgp.StringPrefixSize + 12 + msgp.StringPrefixSize + len(z.ProjectileId) + msgp.StringPrefixSize + 7 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *AccelerationData) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "eventId":
			z.EventId, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "direction":
			z.Direction, err = dc.ReadInt8()
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
func (z AccelerationData) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("eventId")
	if err != nil {
		return
	}
	err = en.WriteUint64(z.EventId)
	if err != nil {
		return
	}
	err = en.WriteString("direction")
	if err != nil {
		return
	}
	err = en.WriteInt8(z.Direction)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z AccelerationData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "eventId")
	o = msgp.AppendUint64(o, z.EventId)
	o = msgp.AppendString(o, "direction")
	o = msgp.AppendInt8(o, z.Direction)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccelerationData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "eventId":
			z.EventId, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "direction":
			z.Direction, bts, err = msgp.ReadInt8Bytes(bts)
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

func (z AccelerationData) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 7 + msgp.Uint64Size + msgp.StringPrefixSize + 9 + msgp.Int8Size
	return
}
