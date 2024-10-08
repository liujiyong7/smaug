// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/liujiyong7/smaug/loader`
)

const (
    _entry__lookup_small_key = 96
)

const (
    _stack__lookup_small_key = 56
)

const (
    _size__lookup_small_key = 810
)

var (
    _pcsp__lookup_small_key = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x8, 16},
        {0xa, 24},
        {0xc, 32},
        {0xd, 40},
        {0xe, 48},
        {0x31c, 56},
        {0x31d, 48},
        {0x31f, 40},
        {0x321, 32},
        {0x323, 24},
        {0x325, 16},
        {0x326, 8},
        {0x32a, 0},
    }
)

var _cfunc_lookup_small_key = []loader.CFunc{
    {"_lookup_small_key_entry", 0,  _entry__lookup_small_key, 0, nil},
    {"_lookup_small_key", _entry__lookup_small_key, _size__lookup_small_key, _stack__lookup_small_key, _pcsp__lookup_small_key},
}
