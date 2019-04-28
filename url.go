/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package csn

import (
	"strings"
	"fmt"
)

// https://cn.chiasenhac.vn/mp3/chinese/c-pop/hong-dau-sinh-nam-quoc~dong-le~ts35b07wqhqnt9.html
func ExtractUrlInfo(url string) (CSNUrlInfo, error) {
	var ret CSNUrlInfo
	toks := strings.Split(url, "/")
	fmt.Printf("%#v\n", toks)

	return ret, nil
}
