// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/liujiyong7/smaug/loader`
)

const (
    _entry__quote = 48
)

const (
    _stack__quote = 80
)

const (
    _size__quote = 1760
)

var (
    _pcsp__quote = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x8, 16},
        {0xa, 24},
        {0xc, 32},
        {0xd, 40},
        {0x11, 48},
        {0x6a9, 80},
        {0x6aa, 48},
        {0x6ac, 40},
        {0x6ae, 32},
        {0x6b0, 24},
        {0x6b2, 16},
        {0x6b3, 8},
        {0x6b4, 0},
        {0x6e0, 80},
    }
)

var _cfunc_quote = []loader.CFunc{
    {"_quote_entry", 0,  _entry__quote, 0, nil},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
}
