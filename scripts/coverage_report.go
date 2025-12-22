package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/tmstorm/invgo"
	"golang.org/x/net/html"
)

const invgateDocsLink = "https://releases.invgate.com/service-desk/api"

type ResourceItem struct {
	Name    string   `json:"name"`
	Link    string   `json:"link"`
	Methods []string `json:"methods"`
}

type Endpoints struct {
	CoveragePercent  float64        `json:"coverage_percent"`
	TotalImplemented int            `json:"total_implemented"`
	TotalMethods     int            `json:"total_methods"`
	Endpoints        []ResourceItem `json:"endpoints"`
}

func main() {
	invgateEP := invgateEndpoints()
	var invKeys []string
	for k := range invgateEP {
		invKeys = append(invKeys, k)
	}
	sort.Strings(invKeys)

	f, err := os.Create("API_COVERAGE.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	implemented := Endpoints{}
	for _, ep := range invKeys {
		implementedMethods, ok := invgo.ImplementedEndpoints[ep]
		resource := ResourceItem{
			Name: invgateEP[ep].Name,
			Link: invgateEP[ep].Link,
		}
		for _, m := range invgateEP[ep].Methods {
			implemented.TotalMethods++
			if ok && contains(implementedMethods, m) {
				resource.Methods = append(resource.Methods, m)
				implemented.TotalImplemented++
			}
		}
		if len(resource.Methods) > 0 {
			implemented.Endpoints = append(implemented.Endpoints, resource)
		}
	}
	implemented.CoveragePercent = (float64(implemented.TotalImplemented) / float64(implemented.TotalMethods)) * 100

	fmt.Fprintln(f, "# API Coverage Report")
	fmt.Fprintf(f, "\n**coverage:** %.2f%% (%d/%d methods implemented)\n", implemented.CoveragePercent, implemented.TotalImplemented, implemented.TotalMethods)

	for _, ep := range invKeys {
		fmt.Fprintf(f, "\n### [%s](%s)\n\n", ep, fmt.Sprintf("%s/%s", invgateDocsLink, invgateEP[ep].Link))
		fmt.Fprintln(f, "| Method | Status |")
		fmt.Fprintln(f, "|--------|--------|")

		implementedMethods, ok := invgo.ImplementedEndpoints[ep]
		for _, m := range invgateEP[ep].Methods {
			status := "❌"
			if ok && contains(implementedMethods, m) {
				status = "✅"
			}
			fmt.Fprintf(f, "| %s | %s |\n", m, status)
		}
	}

	data, err := json.MarshalIndent(implemented, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("api_coverage.json", data, 0o644)
	if err != nil {
		panic(err)
	}
}

func invgateEndpoints() map[string]ResourceItem {
	resp, err := http.Get(invgateDocsLink)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	endpoints := map[string]ResourceItem{}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	for n := range doc.Descendants() {
		resource := ResourceItem{}
		if n.Type == html.ElementNode && n.Data == "a" && hasClass(n, "resource-item-title") {
			for _, a := range n.Attr {
				if a.Key == "href" {
					resource.Link = a.Val
				}
			}
			var endpoint string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "b" && c.FirstChild != nil {
					endpoint = strings.TrimSpace(c.FirstChild.Data)
					break
				}
			}
			if endpoint == "" {
				continue
			}

			sib := n.NextSibling
			for sib != nil && !(sib.Type == html.ElementNode && sib.Data == "div" && hasClass(sib, "resource-item-submenu")) {
				sib = sib.NextSibling
			}

			methods := []string{}
			if sib != nil {
				for c := sib.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "a" && hasClass(c, "resource-item") {
						for gc := c.FirstChild; gc != nil; gc = gc.NextSibling {
							if gc.Type == html.ElementNode && gc.Data == "div" && gc.FirstChild != nil {
								method := strings.TrimSpace(gc.FirstChild.Data)
								methods = append(methods, method)
							}
						}
					}
				}
			}

			resource.Name = endpoint
			resource.Methods = methods
			endpoints[endpoint] = resource
		}
	}

	return endpoints
}

func hasClass(n *html.Node, class string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			classes := strings.Fields(attr.Val)
			if ok := slices.Contains(classes, class); ok {
				return ok
			}
		}
	}
	return false
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
