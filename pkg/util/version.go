package util

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

type VersionCompareRes int

const (
	VersionCompareResBig   VersionCompareRes = iota + 1 /// 大
	VersionCompareResSmall                              /// 小
	VersionCompareResEqual                              /// 相等
)

func CompareLittleVer(verA, verB string) VersionCompareRes {
	bytesA := []byte(verA)
	bytesB := []byte(verB)

	lenA := len(bytesA)
	lenB := len(bytesB)
	fmt.Println(lenA, lenB)

	if lenA > lenB {
		return VersionCompareResBig
	}

	if lenA < lenB {
		return VersionCompareResSmall
	}

	// 如果长度相等则按byte位进行比较

	return compareByBytes(bytesA, bytesB)
}

// 按byte位进行比较小版本号
func compareByBytes(verA, verB []byte) VersionCompareRes {
	for index := range verA {
		if verA[index] > verB[index] {
			return VersionCompareResBig
		}
		if verA[index] < verB[index] {
			return VersionCompareResSmall
		}
	}
	return VersionCompareResEqual
}

func Compare(verA, verB string) VersionCompareRes {
	a := strings.Split(verA, ".")
	b := strings.Split(verB, ".")
	for idx, ver := range a {
		av := cast.ToInt(ver)
		bv := 0
		if len(b) >= idx {
			bv = cast.ToInt(b[idx])
		}
		if av > bv {
			return VersionCompareResBig
		}
		if av < bv {
			return VersionCompareResSmall
		}
	}
	return VersionCompareResEqual
}
