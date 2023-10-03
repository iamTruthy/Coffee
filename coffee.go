package coffee

import "strings"

func WithLatte(s string) string {
	return "Your Order:" + strings.ToUpper (s) 
}
