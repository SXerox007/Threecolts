package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type countUniqueUrlsScenario struct {
	url    []string
	output int
}

var countUniqueUrlsTest = []countUniqueUrlsScenario{
	{
		url:    []string{"https://example.com", "https://example.com/"},
		output: 1,
	},
	{
		url:    []string{"https://example.com", "http://example.com"},
		output: 2,
	},
	{
		url:    []string{"https://example.com?", "https://example.com"},
		output: 1,
	},
	{
		url:    []string{"https://example.com?a=1&b=2", "https://example.com?b=2&a=1"},
		output: 1,
	},
}

type countUniqueUrlsPerTopLevelDomainScenario struct {
	url    []string
	output map[string]int
}

var countUniqueUrlsPerTopLevelDomainTest = []countUniqueUrlsPerTopLevelDomainScenario{
	{
		url:    []string{"https://example.com"},
		output: map[string]int{"example.com": 1},
	},
	{
		url:    []string{"https://example.com", "https://subdomain.example.com"},
		output: map[string]int{"example.com": 2},
	},
}

// Count Unique URL
func TestCountUniqueUrls(t *testing.T) {
	t.Cleanup(func() {
		t.Log("TestCountUniqueUrls:")
		for _, sc := range countUniqueUrlsTest {
			output := CountUniqueUrls(sc.url)
			if !cmp.Equal(sc.output, output, nil) {
				t.Error(cmp.Diff(sc.output, output, nil))
			}
			t.Log("Input:", sc.url, " Output:", output)
		}

	})
}

func TestCountUniqueUrlsPerTopLevelDomain(t *testing.T) {
	t.Cleanup(func() {
		t.Log("TestingCountUniqueUrlsPerTopLevelDomain:")
		for _, sc := range countUniqueUrlsPerTopLevelDomainTest {
			output := CountUniqueUrlsPerTopLevelDomain(sc.url)
			if !cmp.Equal(sc.output, output, nil) {
				t.Error(cmp.Diff(sc.output, output, nil))
			}
			t.Log("Input:", sc.url, " Output:", output)
		}

	})
}
