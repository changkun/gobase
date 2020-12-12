// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package algo

import "changkun.de/x/pkg/common"

// BinarySearch implements binary search algorithm.
// It returns the index if element x is found in arr
// Otherwise it returns -1.
func BinarySearch(arr []interface{}, x interface{}, less common.Less) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if less(arr[mid], x) {
			l = mid + 1
		} else if less(x, arr[mid]) {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
