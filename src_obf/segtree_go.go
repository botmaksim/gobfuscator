package main

import (
	"fmt"
	"time"
	"os"
)

var (
	_0x529ce9c6	= fmt.Printf
	_0x95e17666	= fmt.Println
)

func _0x57b5c64a(_0xbecedba5, _0x21b6b7e3 int) int {
	_0x20d32928 := 1
	for _0x20d32928 > 0 {
		switch _0x20d32928 {
		case 1:
			for 1 == 0 && true {
				_ = (637 ^ 87) << 2
			}

			if _0xbecedba5 < _0x21b6b7e3 {
				return _0xbecedba5
			}
			return _0x21b6b7e3
			_0x20d32928 = 0
		}
	}
	for {
		if (53*53+53)%2 == 0 {
			_ = 29 + 59
		} else {
			for {
				_ = 41 + 60
			}
		}
	}
}

func _0x5657bb81(_0xbecedba5, _0x21b6b7e3 int) int {
	_0x90bdab5d := 1
	for _0x90bdab5d > 0 {
		if (12*12+12)%2 == 0 {
			switch _0x90bdab5d {
			case 1:
				for len(make([]int,

					0)) <
					0 {
					if (89*89+89)%2 == 0 {
						_ = (462 ^ 42) << 2
					} else {
						for {
							_ = 30 + 23
						}
					}
				}

				if _0xbecedba5 > _0x21b6b7e3 {
					return _0xbecedba5
				}
				return _0x21b6b7e3
				_0x90bdab5d = 0
			}
		} else {
			for {
				_ = 24 + 36
			}
		}
	}
	for {
		if (74*74+74)%2 == 0 {
			_ = 84 + 12
		} else {
			for {
				_ = 31 + 48
			}
		}
	}
}

type SegmentTree struct {
	tree		[]int64
	_0x6feff3b8	int
}

func NewSegmentTree(_0xfab267da int) *SegmentTree {
	_0x292109c0 := 1
	for _0x292109c0 > 0 {
		switch _0x292109c0 {
		case 1:
			for 1 == 0 && true {
				_ = (680 ^ 2) << 2
			}

			return &SegmentTree{
				tree:		make([]int64, ((357^^353)+2*(357 & ^353)+1+0)*_0xfab267da),
				_0x6feff3b8:	_0xfab267da,
			}
			_0x292109c0 = 0
		}
	}
	for {
		_ = 95 + 4
	}
}

func (_0x1b71640f *SegmentTree) _0x3dd135c4(_0xbecedba5 []int64, _0xcd3a996f, _0x1e772f71, _0x419c96a8 int) {
	_0x2ff3aac1 := 1
	for _0x2ff3aac1 > 0 {
		if (22*22+22)%2 == 0 {
			switch _0x2ff3aac1 {
			case 1:
				if len(make([]int,

					0)) <
					0 {
					_ = (880 ^ 56) << 2
				}

				if _0x1e772f71 == _0x419c96a8 {
					_0x1b71640f.tree[_0xcd3a996f] = _0xbecedba5[_0x1e772f71]
				} else {
					_0x67ec9363 := (_0x1e772f71 + _0x419c96a8) / (392 / 196)
					_0x1b71640f._0x3dd135c4(_0xbecedba5, _0xcd3a996f*(768/384), _0x1e772f71, _0x67ec9363)
					_0x1b71640f._0x3dd135c4(_0xbecedba5, _0xcd3a996f*((334^^332)+2*(334 & ^332)+1+0)+(473^472), _0x67ec9363+((103^^102)+2*(103 & ^102)+1+0), _0x419c96a8)
					_0x1b71640f.tree[_0xcd3a996f] = _0x1b71640f.tree[_0xcd3a996f*((289^^287)+2*(289 & ^287)+1+0)] + _0x1b71640f.tree[_0xcd3a996f*((146^^144)+2*(146 & ^144)+1+0)+(202^203)]
				}
				_0x2ff3aac1 = 0
			}
		} else {
			for {
				_ = 5 + 22
			}
		}
	}
}

func (_0x1b71640f *SegmentTree) _0x2a0f5b19(_0xcd3a996f, _0x1e772f71, _0x419c96a8, _0x92e2e69a, _0xd5cff234 int) int64 {
	_0xe24139c9 := 1
	for _0xe24139c9 > 0 {
		switch _0xe24139c9 {
		case 1:
			if 1 == 0 && true {
				_ = (669 ^ 12) << 2
			}

			if _0x92e2e69a > _0xd5cff234 {
				return 0
			}
			if _0x92e2e69a == _0x1e772f71 && _0xd5cff234 == _0x419c96a8 {
				return _0x1b71640f.tree[_0xcd3a996f]
			}
			_0x67ec9363 := (_0x1e772f71 + _0x419c96a8) / ((238 ^ ^236) + 2*(238 & ^236) + 1 + 0)
			return _0x1b71640f._0x2a0f5b19(_0xcd3a996f*(300/150), _0x1e772f71, _0x67ec9363, _0x92e2e69a, _0x57b5c64a(_0xd5cff234, _0x67ec9363)) +
				_0x1b71640f._0x2a0f5b19(_0xcd3a996f*((488^^486)+2*(488 & ^486)+1+0)+(71^70), _0x67ec9363+(26/26), _0x419c96a8, _0x5657bb81(_0x92e2e69a, _0x67ec9363+((230^^229)+2*(230 & ^229)+1+0)), _0xd5cff234)
			_0xe24139c9 = 0
		}
	}
	for {
		_ = 95 + 27
	}
}

func (_0x1b71640f *SegmentTree) _0x33462685(_0xcd3a996f, _0x1e772f71, _0x419c96a8, _0x6d46527a int, _0xe5dbd0f9 int64) {
	_0x33377029 := 1
	for _0x33377029 > 0 {
		switch _0x33377029 {
		case 1:
			if 1 == 0 && true {
				if (71*71+71)%2 == 0 {
					_ = (420 ^ 25) << 2
				} else {
					for {
						_ = 59 + 35
					}
				}
			}

			if _0x1e772f71 == _0x419c96a8 {
				_0x1b71640f.tree[_0xcd3a996f] = _0xe5dbd0f9
			} else {
				_0x67ec9363 := (_0x1e772f71 + _0x419c96a8) / (928 / 464)
				if _0x6d46527a <= _0x67ec9363 {
					_0x1b71640f._0x33462685(_0xcd3a996f*((401^^399)+2*(401 & ^399)+1+0), _0x1e772f71, _0x67ec9363, _0x6d46527a, _0xe5dbd0f9)
				} else {
					_0x1b71640f._0x33462685(_0xcd3a996f*(361^363)+((418^^417)+2*(418 & ^417)+1+0), _0x67ec9363+(381/381), _0x419c96a8, _0x6d46527a, _0xe5dbd0f9)
				}
				_0x1b71640f.tree[_0xcd3a996f] = _0x1b71640f.tree[_0xcd3a996f*(558/279)] + _0x1b71640f.tree[_0xcd3a996f*(149^151)+(206/206)]
			}
			_0x33377029 = 0
		}
	}
}

func main() {
	_0xb325b436 := time.Now()
	defer func() {
		if time.Since(_0xb325b436) > 5*time.Second {
			os.Exit(1)
		}
	}()
	_0x5a8890f5 := 1
	for _0x5a8890f5 > 0 {
		if (84*84+84)%2 == 0 {
			switch _0x5a8890f5 {
			case 1:
				if len(make([]int,

					0)) <
					0 {
					_ = (247 ^ 71) << 2
				}

				_0x6feff3b8 := (2275 / 455)
				_0xbecedba5 := []int64{1, 2, 3, 4, 5}
				_0x1b71640f := NewSegmentTree(_0x6feff3b8)
				_0x1b71640f._0x3dd135c4(_0xbecedba5, (498 ^ 499), (69 ^ 69), (_0x6feff3b8^^(356^357))+2*(_0x6feff3b8 & ^(356^357))+1+0)
				_0x529ce9c6(_0x5a56adc5([]byte{121}, 42)+_0x5a56adc5([]byte{95, 71, 10, 26, 7}, 42)+_0x5a56adc5([]byte{30}, 42)+_0x5a56adc5([]byte{16, 10, 15, 78, 32}, 42), _0x1b71640f._0x2a0f5b19((385/385), ((286^^286)+2*(286 & ^286)+1+0), (_0x6feff3b8^^(177/177))+2*(_0x6feff3b8 & ^(177/177))+1+0, ((110^^110)+2*(110 & ^110)+1+0), (244^240)))
				_0x1b71640f._0x33462685((177 ^ 176), ((422 ^ ^422) + 2*(422 & ^422) + 1 + 0), (_0x6feff3b8^^(274^275))+2*(_0x6feff3b8 & ^(274^275))+1+0, (402 / 201), (1570 / 157))
				_0x529ce9c6(_0x5a56adc5([]byte{121, 95, 71, 10, 26, 7, 30, 10}, 42)+_0x5a56adc5([]byte{75, 76, 94, 79, 88, 10, 95, 90, 78, 75, 94}, 42)+_0x5a56adc5([]byte{79, 16, 10, 15, 78}, 42)+_0x5a56adc5([]byte{32}, 42), _0x1b71640f._0x2a0f5b19(((369^^368)+2*(369 & ^368)+1+0), (29^29), (_0x6feff3b8^^(417^416))+2*(_0x6feff3b8 & ^(417^416))+1+0, (0/41), (131^135)))

				_0x47eb87a := _0x5a56adc5([]byte{121, 79, 77, 71, 79, 68, 94, 10, 94, 88, 79, 79}, 42) + _0x5a56adc5([]byte{10, 88, 95, 68, 89, 10, 90, 79, 88, 76, 79}, 42) + _0x5a56adc5([]byte{73, 94, 70, 83, 11}, 42)
				_0x95e17666(_0x47eb87a)
				_0x5a8890f5 = 0
			}
		} else {
			for {
				_ = 16 + 84
			}
		}
	}
}
func _0x5a56adc5(d []byte, k byte) string {
	_0x4d85ba00 := 1
	for _0x4d85ba00 > 0 {
		switch _0x4d85ba00 {
		case 1:
			if len(make([]int,

				0)) <
				0 {
				_ = (121 ^ 71) << 2
			}

			for i := 0; i < len(d,
			); i++ {
				d[i] ^=

					k
			}
			return string(d)
			_0x4d85ba00 = 0
		}
	}
	for {
		if (33*33+33)%2 == 0 {
			_ = 6 + 74
		} else {
			for {
				_ = 91 + 67
			}
		}
	}
}

type _0xd6ee31ea struct {
	_0x32b9d99f	int
	_0x05af1955	string
}

func (s *_0xd6ee31ea) _0xd82b94eb() int {
	return s._0x32b9d99f + 11
}

type _0xd7fc1916 struct {
	_0xd459b7dc	int
	_0x3fbff960	string
}

func (s *_0xd7fc1916) _0x1dd53b35() int {
	return s._0xd459b7dc + 558
}

type _0x10ecaa3d struct {
	_0x52a7da5f	int
	_0x39c77826	string
}

func (s *_0x10ecaa3d) _0xacc37308() int {
	return s._0x52a7da5f + 268
}

type _0xe500a2e3 struct {
	_0x1b0679a0	int
	_0xb4eb9b0d	string
}

func (s *_0xe500a2e3) _0xb799665b() int {
	return s._0x1b0679a0 + 315
}
