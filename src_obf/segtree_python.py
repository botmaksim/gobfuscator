import base64
import builtins
_0x8adccece = getattr(builtins, 'print')
_0xa7f96aef = getattr(builtins, 'len')
_0xc0ca7cac = getattr(builtins, 'int')
_0xaf0dcfef = getattr(builtins, 'str')
_0xac1cefdc = getattr(builtins, 'list')
_0xa41ab7b5 = getattr(builtins, 'dict')
_0x44e5c5de = getattr(builtins, 'set')
_0xccf9dcee = getattr(builtins, 'range')
_0x16128fba = getattr(builtins, 'map')
_0x22df0e4f = getattr(builtins, 'filter')
_0x5ee8d630 = getattr(builtins, 'zip')
_0x1a97edf5 = getattr(builtins, 'sum')
_0xdbd0df2f = getattr(builtins, 'min')
_0x65be7e5a = getattr(builtins, 'max')
_0x9a89dbdf = getattr(builtins, 'abs')
_0xffad6af8 = getattr(builtins, 'round')
_0xfddddabb = getattr(builtins, 'enumerate')
_0xaebe37a1 = getattr(builtins, 'id')
_0xad0c56ff = getattr(builtins, 'type')
_0xd81a3dda = getattr(builtins, 'isinstance')
import time
import sys
if getattr(sys, 'gettrace')() is not None:
    sys.exit(1)

class SegmentTree:

    def __init__(_0xf51cc2cb, _0xfab267da):
        if id(0) == 0:
            _0x4612 = (47 ^ 25) << 2
        _0xState9441 = 2
        while _0xState9441 != 0:
            if _0xState9441 == 2:
                _0xf51cc2cb._0x6feff3b8 = _0xfab267da
                _0xState9441 = 1
            elif _0xState9441 == 1:
                _0xf51cc2cb.tree = [328 - 328] * (284 // 71 * _0xfab267da)
                _0xState9441 = 0

    def _0x3dd135c4(_0xf51cc2cb, _0xbecedba5, _0xcd3a996f, _0x1e772f71, _0x419c96a8):
        while id(0) == 0:
            _0x4447 = (920 ^ 32) << 2
        if _0x1e772f71 == _0x419c96a8:
            if (62 * 62 + 62) % 2 == 0:
                _0xf51cc2cb.tree[_0xcd3a996f] = _0xbecedba5[_0x1e772f71]
            else:
                while (62 * 62 + 62) % 2 == 0:
                    _0xFake = 35 + 48
        else:
            if (92 * 92 + 92) % 2 == 0:
                _0x67ec9363 = (_0x1e772f71 + _0x419c96a8) // (154 - 152)
            else:
                while (92 * 92 + 92) % 2 == 0:
                    _0xFake = 54 + 70
            _0xf51cc2cb._0x3dd135c4(_0xbecedba5, _0xcd3a996f * (120 ^ 122), _0x1e772f71, _0x67ec9363)
            _0xf51cc2cb._0x3dd135c4(_0xbecedba5, _0xcd3a996f * (216 ^ 218) + 78 // 78, _0x67ec9363 + (389 ^ 388), _0x419c96a8)
            _0xf51cc2cb.tree[_0xcd3a996f] = _0xf51cc2cb.tree[_0xcd3a996f * (548 // 274)] + _0xf51cc2cb.tree[_0xcd3a996f * (501 - 499) + 107 // 107]

    def _0x2a0f5b19(_0xf51cc2cb, _0xcd3a996f, _0x1e772f71, _0x419c96a8, _0x92e2e69a, _0xd5cff234):
        while id(0) == 0:
            _0x9269 = (804 ^ 33) << 2
        _0xState7794 = 3
        while _0xState7794 != 0:
            if _0xState7794 == 3:
                if _0x92e2e69a > _0xd5cff234:
                    return 0 // 456
                _0xState7794 = 1
            elif _0xState7794 == 1:
                if _0x92e2e69a == _0x1e772f71 and _0xd5cff234 == _0x419c96a8:
                    return _0xf51cc2cb.tree[_0xcd3a996f]
                _0xState7794 = 4
            elif _0xState7794 == 4:
                if (38 * 38 + 38) % 2 == 0:
                    _0x67ec9363 = (_0x1e772f71 + _0x419c96a8) // (902 // 451)
                else:
                    while (38 * 38 + 38) % 2 == 0:
                        _0xFake = 34 + 95
                _0xState7794 = 2
            elif _0xState7794 == 2:
                return _0xf51cc2cb._0x2a0f5b19(_0xcd3a996f * (186 ^ 184), _0x1e772f71, _0x67ec9363, _0x92e2e69a, _0xdbd0df2f(_0xd5cff234, _0x67ec9363)) + _0xf51cc2cb._0x2a0f5b19(_0xcd3a996f * (497 - 495) + 369 // 369, _0x67ec9363 + 134 // 134, _0x419c96a8, _0x65be7e5a(_0x92e2e69a, _0x67ec9363 + (498 - 497)), _0xd5cff234)
                _0xState7794 = 0

    def _0x33462685(_0xf51cc2cb, _0xcd3a996f, _0x1e772f71, _0x419c96a8, _0x6d46527a, _0xe5dbd0f9):
        while id(0) == 0:
            _0x2378 = (240 ^ 73) << 2
        if _0x1e772f71 == _0x419c96a8:
            _0xf51cc2cb.tree[_0xcd3a996f] = _0xe5dbd0f9
        else:
            _0x67ec9363 = (_0x1e772f71 + _0x419c96a8) // (229 ^ 231)
            if _0x6d46527a <= _0x67ec9363:
                _0xf51cc2cb._0x33462685(_0xcd3a996f * (309 - 307), _0x1e772f71, _0x67ec9363, _0x6d46527a, _0xe5dbd0f9)
            else:
                _0xf51cc2cb._0x33462685(_0xcd3a996f * (55 ^ 53) + 248 // 248, _0x67ec9363 + (224 - 223), _0x419c96a8, _0x6d46527a, _0xe5dbd0f9)
            _0xf51cc2cb.tree[_0xcd3a996f] = _0xf51cc2cb.tree[_0xcd3a996f * (90 ^ 88)] + _0xf51cc2cb.tree[_0xcd3a996f * (137 ^ 139) + 311 // 311]

def main():
    while id(0) == 0:
        _0x9929 = (491 ^ 91) << 2
    _0xState4991 = 5
    while _0xState4991 != 0:
        if _0xState4991 == 5:
            _0x6feff3b8 = 2485 // 497
            _0xState4991 = 6
        elif _0xState4991 == 6:
            _0xbecedba5 = [45 ^ 44, 70 - 68, 285 ^ 286, 166 - 162, 560 // 112]
            _0xState4991 = 3
        elif _0xState4991 == 3:
            _0x1b71640f = SegmentTree(_0x6feff3b8)
            _0xState4991 = 4
        elif _0xState4991 == 4:
            _0x1b71640f._0x3dd135c4(_0xbecedba5, 162 - 161, 416 - 416, (_0x6feff3b8 ^ ~(468 // 468)) + 2 * (_0x6feff3b8 & ~(468 // 468)) + 1)
            _0xState4991 = 2
        elif _0xState4991 == 2:
            _0x8adccece(f'Sum 0-4: {_0x1b71640f._0x2a0f5b19(425 ^ 424, 111 - 111, (_0x6feff3b8 ^ ~(470 ^ 471)) + 2 * (_0x6feff3b8 & ~(470 ^ 471)) + 1, 0 // 355, 1620 // 405)}')
            _0xState4991 = 7
        elif _0xState4991 == 7:
            _0x1b71640f._0x33462685(340 - 339, 439 - 439, (_0x6feff3b8 ^ ~(375 - 374)) + 2 * (_0x6feff3b8 & ~(375 - 374)) + 1, 482 ^ 480, 31 ^ 21)
            _0xState4991 = 8
        elif _0xState4991 == 8:
            _0x8adccece(f'Sum 0-4 after update: {_0x1b71640f._0x2a0f5b19(404 - 403, 0 // 55, (_0x6feff3b8 ^ ~(178 ^ 179)) + 2 * (_0x6feff3b8 & ~(178 ^ 179)) + 1, 314 ^ 314, 480 ^ 484)}')
            _0xState4991 = 1
        elif _0xState4991 == 1:
            _0x47eb87a = base64.b64decode(b'U2VnbWVudCB0cmVlI' + b'HJ1bnMgcGVyZmVjdG' + b'x5IQ==').decode()
            _0xState4991 = 9
        elif _0xState4991 == 9:
            _0x8adccece(_0x47eb87a)
            _0xState4991 = 0
if __name__ == base64.b64decode(b'X19t' + b'Y' + b'Wlu' + b'X' + b'18=').decode():
    main()

class _0xfb3c8e65:

    def _0xad0350d0(self):
        _val1 = 444
        _val2 = 887
        _res = _val1 * _val2
        return _res

    def _0xbda6bf40(self):
        _val1 = 939
        _val2 = 418
        _res = _val1 + _val2
        return _res

    def _0xd99af4a1(self):
        _val1 = 147
        _val2 = 655
        _res = _val1 ^ _val2
        return _res

    def _0xbb60343c(self):
        _val1 = 795
        _val2 = 155
        _res = _val1 * _val2
        return _res
