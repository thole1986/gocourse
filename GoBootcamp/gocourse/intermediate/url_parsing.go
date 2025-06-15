package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	// fmt.Println("Scheme:", parsedURL.Scheme)

	rawURL1 := "https://example.com/path?name=John&age=30"

	parseURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	queryParams := parseURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	// Building URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL:", baseURL.String())

	values := url.Values{}

	// Add key-value pairs to the values object
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	// Encode
	encodedQuery := values.Encode()

	fmt.Println(encodedQuery)

	// Build a URL
	baseURL1 := "https://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery

	fmt.Println(fullURL)
}
