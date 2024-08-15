
// Code generated by Makefile, DO NOT EDIT.

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package avx2

import (
    `github.com/liujiyong7/smaug/loader`
)

func Use() {
    loader.WrapGoC(_text_f64toa, _cfunc_f64toa, []loader.GoC{{"_f64toa", &S_f64toa, &F_f64toa}}, "avx2", "avx2/f64toa.c")
    loader.WrapGoC(_text_f32toa, _cfunc_f32toa, []loader.GoC{{"_f32toa", &S_f32toa, &F_f32toa}}, "avx2", "avx2/f32toa.c")
    loader.WrapGoC(_text_get_by_path, _cfunc_get_by_path, []loader.GoC{{"_get_by_path", &S_get_by_path, &F_get_by_path}}, "avx2", "avx2/get_by_path.c")
    loader.WrapGoC(_text_html_escape, _cfunc_html_escape, []loader.GoC{{"_html_escape", &S_html_escape, &F_html_escape}}, "avx2", "avx2/html_escape.c")
    loader.WrapGoC(_text_i64toa, _cfunc_i64toa, []loader.GoC{{"_i64toa", &S_i64toa, &F_i64toa}}, "avx2", "avx2/i64toa.c")
    loader.WrapGoC(_text_lspace, _cfunc_lspace, []loader.GoC{{"_lspace", &S_lspace, &F_lspace}}, "avx2", "avx2/lspace.c")
    loader.WrapGoC(_text_quote, _cfunc_quote, []loader.GoC{{"_quote", &S_quote, &F_quote}}, "avx2", "avx2/quote.c")
    loader.WrapGoC(_text_skip_array, _cfunc_skip_array, []loader.GoC{{"_skip_array", &S_skip_array, &F_skip_array}}, "avx2", "avx2/skip_array.c")
    loader.WrapGoC(_text_skip_number, _cfunc_skip_number, []loader.GoC{{"_skip_number", &S_skip_number, &F_skip_number}}, "avx2", "avx2/skip_number.c")
    loader.WrapGoC(_text_skip_one, _cfunc_skip_one, []loader.GoC{{"_skip_one", &S_skip_one, &F_skip_one}}, "avx2", "avx2/skip_one.c")
    loader.WrapGoC(_text_skip_object, _cfunc_skip_object, []loader.GoC{{"_skip_object", &S_skip_object, &F_skip_object}}, "avx2", "avx2/skip_object.c")
    loader.WrapGoC(_text_skip_one_fast, _cfunc_skip_one_fast, []loader.GoC{{"_skip_one_fast", &S_skip_one_fast, &F_skip_one_fast}}, "avx2", "avx2/skip_one_fast.c")
    loader.WrapGoC(_text_u64toa, _cfunc_u64toa, []loader.GoC{{"_u64toa", &S_u64toa, &F_u64toa}}, "avx2", "avx2/u64toa.c")
    loader.WrapGoC(_text_unquote, _cfunc_unquote, []loader.GoC{{"_unquote", &S_unquote, &F_unquote}}, "avx2", "avx2/unquote.c")
    loader.WrapGoC(_text_validate_one, _cfunc_validate_one, []loader.GoC{{"_validate_one", &S_validate_one, &F_validate_one}}, "avx2", "avx2/validate_one.c")
    loader.WrapGoC(_text_validate_utf8, _cfunc_validate_utf8, []loader.GoC{{"_validate_utf8", &S_validate_utf8, &F_validate_utf8}}, "avx2", "avx2/validate_utf8.c")
    loader.WrapGoC(_text_validate_utf8_fast, _cfunc_validate_utf8_fast, []loader.GoC{{"_validate_utf8_fast", &S_validate_utf8_fast, &F_validate_utf8_fast}}, "avx2", "avx2/validate_utf8_fast.c")
    loader.WrapGoC(_text_vnumber, _cfunc_vnumber, []loader.GoC{{"_vnumber", &S_vnumber, &F_vnumber}}, "avx2", "avx2/vnumber.c")
    loader.WrapGoC(_text_vsigned, _cfunc_vsigned, []loader.GoC{{"_vsigned", &S_vsigned, &F_vsigned}}, "avx2", "avx2/vsigned.c")
    loader.WrapGoC(_text_vunsigned, _cfunc_vunsigned, []loader.GoC{{"_vunsigned", &S_vunsigned, &F_vunsigned}}, "avx2", "avx2/vunsigned.c")
    loader.WrapGoC(_text_vstring, _cfunc_vstring, []loader.GoC{{"_vstring", &S_vstring, &F_vstring}}, "avx2", "avx2/vstring.c")
    loader.WrapGoC(_text_value, _cfunc_value, []loader.GoC{{"_value", &S_value, &F_value}}, "avx2", "avx2/value.c")
    loader.WrapGoC(_text_parse_with_padding, _cfunc_parse_with_padding, []loader.GoC{{"_parse_with_padding", &S_parse_with_padding, &F_parse_with_padding}}, "avx2", "avx2/parser.c")
    loader.WrapGoC(_text_lookup_small_key, _cfunc_lookup_small_key, []loader.GoC{{"_lookup_small_key", &S_lookup_small_key, &F_lookup_small_key}}, "avx2", "avx2/lookup.c")
}
