package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GameState) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "time":
			z.Time, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "ships":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Ships == nil && msz > 0 {
				z.Ships = make(map[string]*Ship, msz)
			} else if len(z.Ships) > 0 {
				for key, _ := range z.Ships {
					delete(z.Ships, key)
				}
			}
			for msz > 0 {
				msz--
				var xvk string
				var bzg *Ship
				xvk, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					bzg = nil
				} else {
					if bzg == nil {
						bzg = new(Ship)
					}
					err = bzg.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Ships[xvk] = bzg
			}
		case "projectiles":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Projectiles == nil && msz > 0 {
				z.Projectiles = make(map[string]*Projectile, msz)
			} else if len(z.Projectiles) > 0 {
				for key, _ := range z.Projectiles {
					delete(z.Projectiles, key)
				}
			}
			for msz > 0 {
				msz--
				var bai string
				var cmr *Projectile
				bai, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					cmr = nil
				} else {
					if cmr == nil {
						cmr = new(Projectile)
					}
					err = cmr.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Projectiles[bai] = cmr
			}
		case "asteroids":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Asteroids == nil && msz > 0 {
				z.Asteroids = make(map[string]*Asteroid, msz)
			} else if len(z.Asteroids) > 0 {
				for key, _ := range z.Asteroids {
					delete(z.Asteroids, key)
				}
			}
			for msz > 0 {
				msz--
				var ajw string
				var wht *Asteroid
				ajw, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					wht = nil
				} else {
					if wht == nil {
						wht = new(Asteroid)
					}
					err = wht.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Asteroids[ajw] = wht
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
func (z *GameState) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(4)
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
	err = en.WriteString("ships")
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Ships)))
	if err != nil {
		return
	}
	for xvk, bzg := range z.Ships {
		err = en.WriteString(xvk)
		if err != nil {
			return
		}
		if bzg == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = bzg.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	err = en.WriteString("projectiles")
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Projectiles)))
	if err != nil {
		return
	}
	for bai, cmr := range z.Projectiles {
		err = en.WriteString(bai)
		if err != nil {
			return
		}
		if cmr == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = cmr.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	err = en.WriteString("asteroids")
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Asteroids)))
	if err != nil {
		return
	}
	for ajw, wht := range z.Asteroids {
		err = en.WriteString(ajw)
		if err != nil {
			return
		}
		if wht == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = wht.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GameState) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 4)
	o = msgp.AppendString(o, "time")
	o = msgp.AppendUint64(o, z.Time)
	o = msgp.AppendString(o, "ships")
	o = msgp.AppendMapHeader(o, uint32(len(z.Ships)))
	for xvk, bzg := range z.Ships {
		o = msgp.AppendString(o, xvk)
		if bzg == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = bzg.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	o = msgp.AppendString(o, "projectiles")
	o = msgp.AppendMapHeader(o, uint32(len(z.Projectiles)))
	for bai, cmr := range z.Projectiles {
		o = msgp.AppendString(o, bai)
		if cmr == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = cmr.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	o = msgp.AppendString(o, "asteroids")
	o = msgp.AppendMapHeader(o, uint32(len(z.Asteroids)))
	for ajw, wht := range z.Asteroids {
		o = msgp.AppendString(o, ajw)
		if wht == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = wht.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GameState) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "time":
			z.Time, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "ships":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Ships == nil && msz > 0 {
				z.Ships = make(map[string]*Ship, msz)
			} else if len(z.Ships) > 0 {
				for key, _ := range z.Ships {
					delete(z.Ships, key)
				}
			}
			for msz > 0 {
				var xvk string
				var bzg *Ship
				msz--
				xvk, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					bzg = nil
				} else {
					if bzg == nil {
						bzg = new(Ship)
					}
					bts, err = bzg.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
				z.Ships[xvk] = bzg
			}
		case "projectiles":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Projectiles == nil && msz > 0 {
				z.Projectiles = make(map[string]*Projectile, msz)
			} else if len(z.Projectiles) > 0 {
				for key, _ := range z.Projectiles {
					delete(z.Projectiles, key)
				}
			}
			for msz > 0 {
				var bai string
				var cmr *Projectile
				msz--
				bai, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					cmr = nil
				} else {
					if cmr == nil {
						cmr = new(Projectile)
					}
					bts, err = cmr.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
				z.Projectiles[bai] = cmr
			}
		case "asteroids":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Asteroids == nil && msz > 0 {
				z.Asteroids = make(map[string]*Asteroid, msz)
			} else if len(z.Asteroids) > 0 {
				for key, _ := range z.Asteroids {
					delete(z.Asteroids, key)
				}
			}
			for msz > 0 {
				var ajw string
				var wht *Asteroid
				msz--
				ajw, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					wht = nil
				} else {
					if wht == nil {
						wht = new(Asteroid)
					}
					bts, err = wht.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
				z.Asteroids[ajw] = wht
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

func (z *GameState) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.Uint64Size + msgp.StringPrefixSize + 5 + msgp.MapHeaderSize
	if z.Ships != nil {
		for xvk, bzg := range z.Ships {
			_ = bzg
			s += msgp.StringPrefixSize + len(xvk)
			if bzg == nil {
				s += msgp.NilSize
			} else {
				s += bzg.Msgsize()
			}
		}
	}
	s += msgp.StringPrefixSize + 11 + msgp.MapHeaderSize
	if z.Projectiles != nil {
		for bai, cmr := range z.Projectiles {
			_ = cmr
			s += msgp.StringPrefixSize + len(bai)
			if cmr == nil {
				s += msgp.NilSize
			} else {
				s += cmr.Msgsize()
			}
		}
	}
	s += msgp.StringPrefixSize + 9 + msgp.MapHeaderSize
	if z.Asteroids != nil {
		for ajw, wht := range z.Asteroids {
			_ = wht
			s += msgp.StringPrefixSize + len(ajw)
			if wht == nil {
				s += msgp.NilSize
			} else {
				s += wht.Msgsize()
			}
		}
	}
	return
}
