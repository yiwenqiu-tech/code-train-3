package _map

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var domainCountMap = make(map[string]int)

	for _, cpdomain := range cpdomains {
		countAndDomain := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(countAndDomain[0])
		fullDomain := countAndDomain[1]

		domainCountMap[fullDomain] += count
		tmpDomain := fullDomain
		firstIndex := strings.Index(tmpDomain, ".")
		for firstIndex != -1 {
			tmpDomain = tmpDomain[firstIndex+1:]
			domainCountMap[tmpDomain] += count
			firstIndex = strings.Index(tmpDomain, ".")
		}
	}

	var res []string
	for domain, count := range domainCountMap {
		res = append(res, strconv.Itoa(count)+" "+domain)
	}
	return res
}
