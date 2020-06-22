# wxregion

获取微信用户信息中原始地区信息为 CN、Beijing、Chaoyang 这种格式的英文（注意不是汉语拼音），某些情况下需要手动切换为中文，本库提供该能力。

```go

// demo

region := WxRegionToCN("Cn", "Beijing", "chaoyang")
fmt.Println(region)
// {中国 北京 朝阳}

region = WxRegionToCN("Cn", "Beijing", "chaoyang2")
fmt.Println(region)
// {中国 北京 chaoyang2}

```