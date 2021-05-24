package util

import (
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

// StringArrayContainsV1 ...
func StringArrayContainsV1(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// StringArrayIncludeDuplicatesV1 ...
func StringArrayIncludeDuplicatesV1(arr []string) bool {
	flg := false

	encountered := map[string]bool{}
	for i := 0; i < len(arr); i++ {
		if encountered[arr[i]] {
			flg = true
			break
		} else {
			encountered[arr[i]] = true
		}
	}
	return flg
}

// MatchStringV1 ...
func MatchStringV1(reg string, s string) bool {
	return regexp.MustCompile(reg).MatchString(s)
}

// MatchNumberStringV1 ...
func MatchNumberStringV1(s string) bool {
	return MatchStringV1("^[0-9]+$", s)
}

// MatchAlphanumericStringV1 ...
func MatchAlphanumericStringV1(s string) bool {
	return MatchStringV1("^[a-zA-Z0-9]+$", s)
}

// MatchAlphanumericSymbolStringV1 ...
func MatchAlphanumericSymbolStringV1(s string) bool {
	// a-zA-Z0-9?!#$,.:;()<>[]{}_'`|~/*+=^&@%\"-
	return MatchStringV1("^[a-zA-Z0-9?!#$,.:;()<>[\\]{}_'`|~/*\\-+=\\^&@%\\\"]+$", s)
}

// MatchAlphanumericSymbolStringV1 ...
func ValidateUserAgent(s string) bool {
	// a-zA-Z0-9?!#$,.:;()<>[]{}_'`|~/*+=^&@%\"-
	return MatchStringV1("^.+?[/\\s][/\\d.:]+$", s)
}

// ContainsForbiddenCharactersFileNameV1 ...
func ContainsForbiddenCharactersFileNameV1(s string) bool {
	// &$@=;: +,?\^`><{}][#%~|
	return MatchStringV1("^[&$@=;: +,?\\\\^`><{}\\][#%~|]+$", s)
}

// SecondsBeforeAndAfterV1 ...
func SecondsBeforeAndAfterV1(c int64, t int64) bool {
	nut := NowTimeV1().UTC().Unix()
	if (nut-t) <= c && c <= (nut+t) {
		return true
	}
	return false
}

// GetFileNameWithoutExt ...
func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// FixedValueFillingV1 ...
func FixedValueFillingV1(fixedValue byte, num int) []byte {
	d := make([]byte, num)
	for i := range d {
		d[i] = fixedValue
	}
	return d
}

// BytesOrderLittleEndianV1 ...
func BytesOrderLittleEndianV1(in []byte) []byte {
	out := make([]byte, len(in))
	_ = copy(out, in)
	for ic, jc := 0, len(out)-1; ic < jc; ic, jc = ic+1, jc-1 {
		out[ic], out[jc] = out[jc], out[ic]
	}
	return out
}

func PrintFields(b interface{}) string {
	val := reflect.ValueOf(b).Elem()
	step := 1
	strUpdate := "set "
	for i := 0; i < val.Type().NumField(); i++ {
		strUpdate += ReplaceItem(val.Type().Field(i).Tag.Get("json")) + " = " + val.Type().Field(i).Tag.Get("json")
		if step < val.Type().NumField() {
			strUpdate += ", "
		}
		step++
	}
	return strUpdate
}

func ReplaceItem(item string) string {
	return strings.ReplaceAll(item, ":", "")
}
