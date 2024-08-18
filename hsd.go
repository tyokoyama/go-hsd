// Package hsdはGo言語プログラミングエッセンスのサンプルパッケージです。
package hsd

// StringDistance returns different character count.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// Distance returns different rune count.
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
