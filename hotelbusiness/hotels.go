//go:build !solution

package hotelbusiness

import (
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

type Pair struct {
	num int
	idx int
}

func ComputeLoad(guests []Guest) []Load {
	array := make([]Pair, 0)
	for _, item := range guests {
		array = append(array, Pair{+1, item.CheckInDate})
		array = append(array, Pair{-1, item.CheckOutDate})
	}

	sort.SliceStable(array, func(i, j int) bool {
		return array[i].idx < array[j].idx
	})
	result := make([]Load, 0)
	currentIdx := -1
	currentNum := 0
	for _, item := range array {
		if item.idx != currentIdx && currentIdx != -1 {
			if len(result) == 0 || currentNum != result[len(result)-1].GuestCount {
				result = append(result, Load{currentIdx, currentNum})
			}
		}
		currentIdx = item.idx
		currentNum += item.num
	}
	if currentIdx != -1 {
		if len(result) == 0 || currentNum != result[len(result)-1].GuestCount {
			result = append(result, Load{currentIdx, currentNum})
		}
	}
	return result
}
