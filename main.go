package main

import (
	"fmt"
	"net/url"
	"strings"
)

/**
* This function counts how many unique normalized valid URLs were passed to the function
*
* Accepts a list of URLs
*
* Example:
*
* input: ['https://example.com']
* output: 1
*
* Notes:
*  - assume none of the URLs have authentication information (username, password).
*
* Normalized URL:
*  - process in which a URL is modified and standardized: https://en.wikipedia.org/wiki/URL_normalization
*
#    For example.
#    These 2 urls are the same:
#    input: ["https://example.com", "https://example.com/"]
#    output: 1
#
#    These 2 are not the same:
#    input: ["https://example.com", "http://example.com"]
#    output 2
#
#    These 2 are the same:
#    input: ["https://example.com?", "https://example.com"]
#    output: 1
#
#    These 2 are the same:
#    input: ["https://example.com?a=1&b=2", "https://example.com?b=2&a=1"]
#    output: 1
*/

func normalizeUrl(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return ""
	}
	parsedURL.RawQuery = ""
	return parsedURL.String()
}

func cleanPath(path string) string {
	parts := []string{}
	for _, p := range strings.Split(path, "/") {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
		} else {
			parts = append(parts, p)
		}
	}
	return "/" + strings.Join(parts, "/")
}

func CountUniqueUrls(urls []string) int {
	seen := make(map[string]bool)
	for _, urlStr := range urls {
		URL := strings.TrimRight(urlStr, "?")
		u, err := url.Parse(URL)
		if err != nil {
			continue
		}
		//normalizedURL := normalizeUrl(urlStr)
		u.RawQuery = ""
		u.Fragment = ""
		u.Path = cleanPath(u.Path)
		seen[u.String()] = true
	}
	return len(seen)
}

/**
 * This function counts how many unique normalized valid URLs were passed to the function per top level domain
 *
 * A top level domain is a domain in the form of example.com. Assume all top level domains end in .com
 * subdomain.example.com is not a top level domain.
 *
 * Accepts a list of URLs
 *
 * Example:
 *
 * input: ["https://example.com"]
 * output: Hash["example.com" => 1]
 *
 * input: ["https://example.com", "https://subdomain.example.com"]
 * output: Hash["example.com" => 2]
 *
 */

func getMainDomain(inputURL string) string {
	inputURL = "https://" + inputURL
	u, err := url.Parse(inputURL)
	if err != nil {
		return ""
	}
	host := u.Hostname()
	if strings.Count(host, ".") > 1 {
		lastDot := strings.LastIndex(host, ".")
		secondLastDot := strings.LastIndex(host[:lastDot-1], ".")
		return host[secondLastDot+1:]
	}
	return host
}

func CountUniqueUrlsPerTopLevelDomain(urls []string) map[string]int {
	seen := make(map[string]int)
	for _, urlStr := range urls {
		u, err := url.Parse(urlStr)
		if err != nil {
			continue
		}
		u.RawQuery = ""
		u.Fragment = ""
		u.Path = cleanPath(u.Path)
		host := strings.TrimSuffix(u.Host, ".com")
		if host != "" {
			host += ".com"
			host = getMainDomain(host)
			seen[host]++
		}
	}
	return seen
}

func problem1() {
	urls := []string{"https://example.com", "https://example.com/"} // 1
	fmt.Println("Input:", urls)
	fmt.Println(CountUniqueUrls(urls))

	urls = []string{"https://example.com", "http://example.com"} // 2
	fmt.Println("Input:", urls)
	fmt.Println(CountUniqueUrls(urls))

	urls = []string{"https://example.com?", "https://example.com"} //1
	fmt.Println("Input:", urls)
	fmt.Println(CountUniqueUrls(urls))

	urls = []string{"https://example.com?a=1&b=2", "https://example.com?b=2&a=1"} //1
	fmt.Println("Input:", urls)
	fmt.Println(CountUniqueUrls(urls))
}

func problem2() {
	urls := []string{"https://example.com"} // Hash["example.com" => 1]
	fmt.Println(CountUniqueUrlsPerTopLevelDomain(urls))

	urls = []string{"https://example.com", "https://subdomain.example.com"} // Hash["example.com" => 2]
	fmt.Println(CountUniqueUrlsPerTopLevelDomain(urls))
}

func main() {
	problem1()
	problem2()
}
