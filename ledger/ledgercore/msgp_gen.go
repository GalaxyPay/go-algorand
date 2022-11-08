package ledgercore

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// AccountTotals
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//
// AlgoCount
//     |-----> (*) MarshalMsg
//     |-----> (*) CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> (*) Msgsize
//     |-----> (*) MsgIsZero
//
// OnlineRoundParamsData
//           |-----> (*) MarshalMsg
//           |-----> (*) CanMarshalMsg
//           |-----> (*) UnmarshalMsg
//           |-----> (*) CanUnmarshalMsg
//           |-----> (*) Msgsize
//           |-----> (*) MsgIsZero
//
// SPVerificationContext
//           |-----> (*) MarshalMsg
//           |-----> (*) CanMarshalMsg
//           |-----> (*) UnmarshalMsg
//           |-----> (*) CanUnmarshalMsg
//           |-----> (*) Msgsize
//           |-----> (*) MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *AccountTotals) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(4)
	var zb0001Mask uint8 /* 5 bits */
	if ((*z).NotParticipating.Money.MsgIsZero()) && ((*z).NotParticipating.RewardUnits == 0) {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if ((*z).Offline.Money.MsgIsZero()) && ((*z).Offline.RewardUnits == 0) {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if ((*z).Online.Money.MsgIsZero()) && ((*z).Online.RewardUnits == 0) {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).RewardsLevel == 0 {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "notpart"
			o = append(o, 0xa7, 0x6e, 0x6f, 0x74, 0x70, 0x61, 0x72, 0x74)
			// omitempty: check for empty values
			zb0002Len := uint32(2)
			var zb0002Mask uint8 /* 3 bits */
			if (*z).NotParticipating.Money.MsgIsZero() {
				zb0002Len--
				zb0002Mask |= 0x2
			}
			if (*z).NotParticipating.RewardUnits == 0 {
				zb0002Len--
				zb0002Mask |= 0x4
			}
			// variable map header, size zb0002Len
			o = append(o, 0x80|uint8(zb0002Len))
			if (zb0002Mask & 0x2) == 0 { // if not empty
				// string "mon"
				o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
				o = (*z).NotParticipating.Money.MarshalMsg(o)
			}
			if (zb0002Mask & 0x4) == 0 { // if not empty
				// string "rwd"
				o = append(o, 0xa3, 0x72, 0x77, 0x64)
				o = msgp.AppendUint64(o, (*z).NotParticipating.RewardUnits)
			}
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "offline"
			o = append(o, 0xa7, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65)
			// omitempty: check for empty values
			zb0003Len := uint32(2)
			var zb0003Mask uint8 /* 3 bits */
			if (*z).Offline.Money.MsgIsZero() {
				zb0003Len--
				zb0003Mask |= 0x2
			}
			if (*z).Offline.RewardUnits == 0 {
				zb0003Len--
				zb0003Mask |= 0x4
			}
			// variable map header, size zb0003Len
			o = append(o, 0x80|uint8(zb0003Len))
			if (zb0003Mask & 0x2) == 0 { // if not empty
				// string "mon"
				o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
				o = (*z).Offline.Money.MarshalMsg(o)
			}
			if (zb0003Mask & 0x4) == 0 { // if not empty
				// string "rwd"
				o = append(o, 0xa3, 0x72, 0x77, 0x64)
				o = msgp.AppendUint64(o, (*z).Offline.RewardUnits)
			}
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "online"
			o = append(o, 0xa6, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65)
			// omitempty: check for empty values
			zb0004Len := uint32(2)
			var zb0004Mask uint8 /* 3 bits */
			if (*z).Online.Money.MsgIsZero() {
				zb0004Len--
				zb0004Mask |= 0x2
			}
			if (*z).Online.RewardUnits == 0 {
				zb0004Len--
				zb0004Mask |= 0x4
			}
			// variable map header, size zb0004Len
			o = append(o, 0x80|uint8(zb0004Len))
			if (zb0004Mask & 0x2) == 0 { // if not empty
				// string "mon"
				o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
				o = (*z).Online.Money.MarshalMsg(o)
			}
			if (zb0004Mask & 0x4) == 0 { // if not empty
				// string "rwd"
				o = append(o, 0xa3, 0x72, 0x77, 0x64)
				o = msgp.AppendUint64(o, (*z).Online.RewardUnits)
			}
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "rwdlvl"
			o = append(o, 0xa6, 0x72, 0x77, 0x64, 0x6c, 0x76, 0x6c)
			o = msgp.AppendUint64(o, (*z).RewardsLevel)
		}
	}
	return
}

func (_ *AccountTotals) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountTotals)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccountTotals) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			var zb0003 int
			var zb0004 bool
			zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Online")
					return
				}
				if zb0003 > 0 {
					zb0003--
					bts, err = (*z).Online.Money.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online", "struct-from-array", "Money")
						return
					}
				}
				if zb0003 > 0 {
					zb0003--
					(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online", "struct-from-array", "RewardUnits")
						return
					}
				}
				if zb0003 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0003)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Online")
					return
				}
				if zb0004 {
					(*z).Online = AlgoCount{}
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online")
						return
					}
					switch string(field) {
					case "mon":
						bts, err = (*z).Online.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Online", "Money")
							return
						}
					case "rwd":
						(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Online", "RewardUnits")
							return
						}
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Online")
							return
						}
					}
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			var zb0005 int
			var zb0006 bool
			zb0005, zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0005, zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Offline")
					return
				}
				if zb0005 > 0 {
					zb0005--
					bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline", "struct-from-array", "Money")
						return
					}
				}
				if zb0005 > 0 {
					zb0005--
					(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline", "struct-from-array", "RewardUnits")
						return
					}
				}
				if zb0005 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0005)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Offline")
					return
				}
				if zb0006 {
					(*z).Offline = AlgoCount{}
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline")
						return
					}
					switch string(field) {
					case "mon":
						bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Offline", "Money")
							return
						}
					case "rwd":
						(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Offline", "RewardUnits")
							return
						}
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Offline")
							return
						}
					}
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			var zb0007 int
			var zb0008 bool
			zb0007, zb0008, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0007, zb0008, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
					return
				}
				if zb0007 > 0 {
					zb0007--
					bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "struct-from-array", "Money")
						return
					}
				}
				if zb0007 > 0 {
					zb0007--
					(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "struct-from-array", "RewardUnits")
						return
					}
				}
				if zb0007 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0007)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
					return
				}
				if zb0008 {
					(*z).NotParticipating = AlgoCount{}
				}
				for zb0007 > 0 {
					zb0007--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
						return
					}
					switch string(field) {
					case "mon":
						bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "Money")
							return
						}
					case "rwd":
						(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "RewardUnits")
							return
						}
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
							return
						}
					}
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "RewardsLevel")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = AccountTotals{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "online":
				var zb0009 int
				var zb0010 bool
				zb0009, zb0010, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0009, zb0010, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Online")
						return
					}
					if zb0009 > 0 {
						zb0009--
						bts, err = (*z).Online.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "Online", "struct-from-array", "Money")
							return
						}
					}
					if zb0009 > 0 {
						zb0009--
						(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Online", "struct-from-array", "RewardUnits")
							return
						}
					}
					if zb0009 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0009)
						if err != nil {
							err = msgp.WrapError(err, "Online", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "Online")
						return
					}
					if zb0010 {
						(*z).Online = AlgoCount{}
					}
					for zb0009 > 0 {
						zb0009--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "Online")
							return
						}
						switch string(field) {
						case "mon":
							bts, err = (*z).Online.Money.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "Online", "Money")
								return
							}
						case "rwd":
							(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Online", "RewardUnits")
								return
							}
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "Online")
								return
							}
						}
					}
				}
			case "offline":
				var zb0011 int
				var zb0012 bool
				zb0011, zb0012, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0011, zb0012, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Offline")
						return
					}
					if zb0011 > 0 {
						zb0011--
						bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "Offline", "struct-from-array", "Money")
							return
						}
					}
					if zb0011 > 0 {
						zb0011--
						(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Offline", "struct-from-array", "RewardUnits")
							return
						}
					}
					if zb0011 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0011)
						if err != nil {
							err = msgp.WrapError(err, "Offline", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "Offline")
						return
					}
					if zb0012 {
						(*z).Offline = AlgoCount{}
					}
					for zb0011 > 0 {
						zb0011--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "Offline")
							return
						}
						switch string(field) {
						case "mon":
							bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "Offline", "Money")
								return
							}
						case "rwd":
							(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Offline", "RewardUnits")
								return
							}
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "Offline")
								return
							}
						}
					}
				}
			case "notpart":
				var zb0013 int
				var zb0014 bool
				zb0013, zb0014, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0013, zb0014, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "NotParticipating")
						return
					}
					if zb0013 > 0 {
						zb0013--
						bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating", "struct-from-array", "Money")
							return
						}
					}
					if zb0013 > 0 {
						zb0013--
						(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating", "struct-from-array", "RewardUnits")
							return
						}
					}
					if zb0013 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0013)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "NotParticipating")
						return
					}
					if zb0014 {
						(*z).NotParticipating = AlgoCount{}
					}
					for zb0013 > 0 {
						zb0013--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating")
							return
						}
						switch string(field) {
						case "mon":
							bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "NotParticipating", "Money")
								return
							}
						case "rwd":
							(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "NotParticipating", "RewardUnits")
								return
							}
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "NotParticipating")
								return
							}
						}
					}
				}
			case "rwdlvl":
				(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RewardsLevel")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *AccountTotals) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountTotals)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AccountTotals) Msgsize() (s int) {
	s = 1 + 7 + 1 + 4 + (*z).Online.Money.Msgsize() + 4 + msgp.Uint64Size + 8 + 1 + 4 + (*z).Offline.Money.Msgsize() + 4 + msgp.Uint64Size + 8 + 1 + 4 + (*z).NotParticipating.Money.Msgsize() + 4 + msgp.Uint64Size + 7 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *AccountTotals) MsgIsZero() bool {
	return (((*z).Online.Money.MsgIsZero()) && ((*z).Online.RewardUnits == 0)) && (((*z).Offline.Money.MsgIsZero()) && ((*z).Offline.RewardUnits == 0)) && (((*z).NotParticipating.Money.MsgIsZero()) && ((*z).NotParticipating.RewardUnits == 0)) && ((*z).RewardsLevel == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z *AlgoCount) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Money.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).RewardUnits == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "mon"
			o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
			o = (*z).Money.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "rwd"
			o = append(o, 0xa3, 0x72, 0x77, 0x64)
			o = msgp.AppendUint64(o, (*z).RewardUnits)
		}
	}
	return
}

func (_ *AlgoCount) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*AlgoCount)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AlgoCount) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Money.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Money")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "RewardUnits")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = AlgoCount{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "mon":
				bts, err = (*z).Money.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Money")
					return
				}
			case "rwd":
				(*z).RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RewardUnits")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *AlgoCount) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*AlgoCount)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AlgoCount) Msgsize() (s int) {
	s = 1 + 4 + (*z).Money.Msgsize() + 4 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *AlgoCount) MsgIsZero() bool {
	return ((*z).Money.MsgIsZero()) && ((*z).RewardUnits == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z *OnlineRoundParamsData) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 4 bits */
	if (*z).OnlineSupply == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).CurrentProtocol.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).RewardsLevel == 0 {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "online"
			o = append(o, 0xa6, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65)
			o = msgp.AppendUint64(o, (*z).OnlineSupply)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "proto"
			o = append(o, 0xa5, 0x70, 0x72, 0x6f, 0x74, 0x6f)
			o = (*z).CurrentProtocol.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "rwdlvl"
			o = append(o, 0xa6, 0x72, 0x77, 0x64, 0x6c, 0x76, 0x6c)
			o = msgp.AppendUint64(o, (*z).RewardsLevel)
		}
	}
	return
}

func (_ *OnlineRoundParamsData) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*OnlineRoundParamsData)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OnlineRoundParamsData) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			(*z).OnlineSupply, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "OnlineSupply")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "RewardsLevel")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).CurrentProtocol.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "CurrentProtocol")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = OnlineRoundParamsData{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "online":
				(*z).OnlineSupply, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "OnlineSupply")
					return
				}
			case "rwdlvl":
				(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RewardsLevel")
					return
				}
			case "proto":
				bts, err = (*z).CurrentProtocol.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "CurrentProtocol")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *OnlineRoundParamsData) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*OnlineRoundParamsData)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *OnlineRoundParamsData) Msgsize() (s int) {
	s = 1 + 7 + msgp.Uint64Size + 7 + msgp.Uint64Size + 6 + (*z).CurrentProtocol.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *OnlineRoundParamsData) MsgIsZero() bool {
	return ((*z).OnlineSupply == 0) && ((*z).RewardsLevel == 0) && ((*z).CurrentProtocol.MsgIsZero())
}

// MarshalMsg implements msgp.Marshaler
func (z *SPVerificationContext) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(4)
	var zb0001Mask uint8 /* 5 bits */
	if (*z).OnlineTotalWeight.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).LastAttestedRound.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Version.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).VotersCommitment.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "pw"
			o = append(o, 0xa2, 0x70, 0x77)
			o = (*z).OnlineTotalWeight.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "spround"
			o = append(o, 0xa7, 0x73, 0x70, 0x72, 0x6f, 0x75, 0x6e, 0x64)
			o = (*z).LastAttestedRound.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "v"
			o = append(o, 0xa1, 0x76)
			o = (*z).Version.MarshalMsg(o)
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "vc"
			o = append(o, 0xa2, 0x76, 0x63)
			o = (*z).VotersCommitment.MarshalMsg(o)
		}
	}
	return
}

func (_ *SPVerificationContext) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*SPVerificationContext)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SPVerificationContext) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).LastAttestedRound.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "LastAttestedRound")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).VotersCommitment.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VotersCommitment")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).OnlineTotalWeight.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "OnlineTotalWeight")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Version.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Version")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = SPVerificationContext{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "spround":
				bts, err = (*z).LastAttestedRound.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "LastAttestedRound")
					return
				}
			case "vc":
				bts, err = (*z).VotersCommitment.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VotersCommitment")
					return
				}
			case "pw":
				bts, err = (*z).OnlineTotalWeight.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "OnlineTotalWeight")
					return
				}
			case "v":
				bts, err = (*z).Version.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Version")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *SPVerificationContext) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*SPVerificationContext)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SPVerificationContext) Msgsize() (s int) {
	s = 1 + 8 + (*z).LastAttestedRound.Msgsize() + 3 + (*z).VotersCommitment.Msgsize() + 3 + (*z).OnlineTotalWeight.Msgsize() + 2 + (*z).Version.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *SPVerificationContext) MsgIsZero() bool {
	return ((*z).LastAttestedRound.MsgIsZero()) && ((*z).VotersCommitment.MsgIsZero()) && ((*z).OnlineTotalWeight.MsgIsZero()) && ((*z).Version.MsgIsZero())
}
