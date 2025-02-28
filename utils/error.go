/*
 * @Author: Edwater
 * @Description:
 * @File: error
 * @Version: 1.0.0
 * @Data: 2025/2/20 16:12
 */

package utils

import (
	"errors"
	"log"
)

var (
	ErrUnsupportedExt = errors.New("unsupported extension")
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
