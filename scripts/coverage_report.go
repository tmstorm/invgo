package main

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/tmstorm/invgo"
	"golang.org/x/net/html"
)

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

	var totalImplemented, totalMethods int
	for _, ep := range invKeys {
		implementedMethods, ok := invgo.ImplementedEndpoints[ep]
		for _, m := range invgateEP[ep] {
			totalMethods++
			if ok && contains(implementedMethods, m) {
				totalImplemented++
			}
		}
	}
	percent := (float64(totalImplemented) / float64(totalMethods)) * 100

	fmt.Fprintln(f, "# API Coverage Report")
	fmt.Fprintf(f, "\n**coverage:** %.2f%% (%d/%d methods implemented)\n", percent, totalImplemented, totalMethods)

	for _, ep := range invKeys {
		fmt.Fprintf(f, "\n### %s\n\n", ep)
		fmt.Fprintln(f, "| Method | Status |")
		fmt.Fprintln(f, "|--------|--------|")

		implementedMethods, ok := invgo.ImplementedEndpoints[ep]
		for _, m := range invgateEP[ep] {
			totalMethods++
			status := "❌"
			if ok && contains(implementedMethods, m) {
				totalImplemented++
				status = "✅"
			}
			fmt.Fprintf(f, "| %s | %s |\n", m, status)
		}
	}
}

func invgateEndpoints() map[string][]string {
	resp, err := http.Get("https://releases.invgate.com/service-desk/api")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	endpoints := map[string][]string{}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" && hasClass(n, "resource-item-title") {
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

			endpoints[endpoint] = methods
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
