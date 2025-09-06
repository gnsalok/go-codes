package test

import "fmt"

//WhatTime will take sec as parameter and return time slice as string.
func WhatTime(seconds int) (string, error) {
	if seconds >= 0 && seconds <= 86399 {
		hConst := 3600
		mConst := 60
		hours := seconds / hConst
		seconds = seconds % hConst
		mins := seconds / mConst
		seconds = seconds % mConst
		return fmt.Sprintf("%d:%d:%d", hours, mins, seconds), nil
	} else {
		return "", fmt.Errorf("Invalid Input, Enter input between 0 to 86399.")

	}
}
