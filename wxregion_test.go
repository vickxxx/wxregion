package wxregion

import (
	"fmt"
	"testing"
)

func TestWxRegionToCN(t *testing.T) {
	fmt.Println("hello test")
	region := WxRegionToCN("Cn", "Beijing", "chaoyang2")
	fmt.Println(region)
}
