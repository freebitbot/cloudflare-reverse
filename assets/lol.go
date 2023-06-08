package main

import (
	"fmt"
	"strings"
)

func cryptData(color string, elem string) string {
	var f func(string, int, func(int) string) string
	f = func(item string, i int, chartCAt func(int) string) string {
		x := map[string]interface{}{
			"IWBEV": "G7Ttuo4rxW6mUa+k2C5eQhN$pqEM8gyFVfcZBSA3dlvRDnY9bPOziwHXJjI-1s0LK",
			"YZHcU": func(progressOld, progressNew int) bool {
				return progressOld < progressNew
			},
			"YjAfQ": func(name string, initialValue interface{}) bool {
				return name == initialValue
			},
			"qDpRl": "SZSSH",
			"fOTVL": func(loaded, total int) bool {
				return loaded > total
			},
			"uQsHx": func(x, y int) bool {
				return x < y
			},
			"qvVdy": func(fun func(int) string, obj int) string {
				return fun(obj)
			},
			"oCnYT": func(leftDiags, columns int) int {
				return leftDiags | columns
			},
			"cbqgM": func(left, right int) int {
				return left << right
			},
			"ZRTzU": func(data, type int) bool {
				return data == type
			},
			"tcoyh": func(fractionalPosition, typeVal int) bool {
				return fractionalPosition > typeVal
			},
			"gPlDx": func(h, i int) int {
				return h - i
			},
			"moDPL": func(fn, useBinaryString int) int {
				return fn & useBinaryString
			},
			"AqyRT": func(replacement func(string) string, css string) string {
				return replacement(css)
			},
			"SPMAy": "myPvT",
			"oWFwx": func(left, right int) int {
				return left << right
			},
			"tHnDI": func(h, i int) int {
				return h - i
			},
			"spysu": func(name, requiredFrom string) bool {
				return name == requiredFrom
			},
			"nJDoD": func(fun func(int) string, obj int) string {
				return fun(obj)
			},
			"HBUgW": func(left, right int) int {
				return left << right
			},
			"POcMj": func(name, requiredFrom string) bool {
				return name == requiredFrom
			},
			"Ygywr": func(fun func(int) string, obj int) string {
				return fun(obj)
			},
			"dIgEk": func(name, days string) bool {
				return name == days
			},
			"IAfOy": func(h, i int) int {
				return h - i
			},
			"ifVeK": func(name, days string) bool {
				return name == days
			},
			"bgpVs": func(x, y int) bool {
				return x < y
			},
		}

		for k, v := range x {
			if k == item {
				switch val := v.(type) {
				case string:
					return val
				case func(int) string:
					return val(i)
				}
			}
		}

		return ""
	}

	colors := strings.Split(color, "")
	elements := strings.Split(elem, "")

	var result string
	for i := 0; i < len(elements); i++ {
		colorIndex := i % len(colors)
		element := elements[i]

		converted := f(element, i, func(num int) string {
			return colors[(num+i)%len(colors)]
		})

		result += converted
	}

	return result
}

func main() {
color := "ABC"
elem := "XYZ"
crypted := cryptData(color, elem)
fmt.Println(crypted)
}