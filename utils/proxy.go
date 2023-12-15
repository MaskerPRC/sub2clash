package utils

import (
	"go.uber.org/zap"
	"strings"
	"sub2clash/logger"
	"sub2clash/model"
	"sub2clash/parser"
)

func GetContryName(countryKey string) string {
	// 创建一个切片包含所有的国家映射
	countryMaps := []map[string]string{
		model.CountryFlag,
		model.CountryChineseName,
		model.CountryISO,
		model.CountryEnglishName,
	}

	// 对每一个映射进行检查
	for i, countryMap := range countryMaps {
		if i == 2 {
			// 对ISO匹配做特殊处理
			// 根据常用分割字符分割字符串
			splitChars := []string{"-", "_", " "}
			key := make([]string, 0)
			for _, splitChar := range splitChars {
				slic := strings.Split(countryKey, splitChar)
				for _, v := range slic {
					if len(v) == 2 {
						key = append(key, v)
					}
				}
			}
			// 对每一个分割后的字符串进行检查
			for _, v := range key {
				// 如果匹配到了国家
				if country, ok := countryMap[strings.ToUpper(v)]; ok {
					return country
				}
			}
		}
		for k, v := range countryMap {
			if strings.Contains(countryKey, k) {
				return v
			}
		}
	}
	return "其他地区"
}

// 添加到某个节点组
func AddNewGroup(sub *model.Subscription, insertGroup string, autotest bool, lazy bool) {
	var newGroup model.ProxyGroup
	if !autotest {
		newGroup = model.ProxyGroup{
			Name:          insertGroup,
			Type:          "select",
			Proxies:       []string{},
			IsCountryGrop: true,
			Size:          1,
		}
	} else {
		newGroup = model.ProxyGroup{
			Name:          insertGroup,
			Type:          "url-test",
			Proxies:       []string{},
			IsCountryGrop: true,
			Url:           "https://www.gstatic.com/generate_204",
			Interval:      300,
			Tolerance:     50,
			Lazy:          lazy,
			Size:          1,
		}
	}
	sub.ProxyGroups = append(sub.ProxyGroups, newGroup)
}

// 添加到某个节点组
func AddToGroup(sub *model.Subscription, proxy model.Proxy, insertGroup string) bool {
	for i := range sub.ProxyGroups {
		group := &sub.ProxyGroups[i]

		if group.Name == insertGroup {
			group.Proxies = append(group.Proxies, proxy.Name)
			group.Size++
			return true
		}
	}
	return false
}

func AddAllNewProxies(
	sub *model.Subscription, autotest bool,
	lazy bool, clashType model.ClashType, proxies ...model.Proxy,
) {
	proxyTypes := model.GetSupportProxyTypes(clashType)

	// 遍历每个代理节点，添加节点
	for _, proxy := range proxies {
		// 跳过无效类型
		if !proxyTypes[proxy.Type] {
			continue
		}
		sub.Proxies = append(sub.Proxies, proxy)
	}
}

func ParseProxy(proxies ...string) []model.Proxy {
	var result []model.Proxy
	for _, proxy := range proxies {
		if proxy != "" {
			var proxyItem model.Proxy
			var err error
			// 解析节点
			if strings.HasPrefix(proxy, "ss://") {
				proxyItem, err = parser.ParseSS(proxy)
			}
			if strings.HasPrefix(proxy, "trojan://") {
				proxyItem, err = parser.ParseTrojan(proxy)
			}
			if strings.HasPrefix(proxy, "vmess://") {
				proxyItem, err = parser.ParseVmess(proxy)
			}
			if strings.HasPrefix(proxy, "vless://") {
				proxyItem, err = parser.ParseVless(proxy)
			}
			if strings.HasPrefix(proxy, "ssr://") {
				proxyItem, err = parser.ParseShadowsocksR(proxy)
			}
			if err == nil {
				result = append(result, proxyItem)
			} else {
				logger.Logger.Debug(
					"parse proxy failed", zap.String("proxy", proxy), zap.Error(err),
				)
			}
		}
	}
	return result
}
