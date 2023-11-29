package cache

import (
	"fmt"
	"strconv"
)

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}

func VideoCacheKey(id uint) string {
	return fmt.Sprintf("video:%s", strconv.Itoa(int(id)))
}

func QueryVideoHistoryKey(uid uint) string {
	return fmt.Sprintf("history:video:user%s", strconv.Itoa(int(uid)))
}
