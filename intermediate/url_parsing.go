package intermediate

import (
	"fmt"
	"net/url"
)

func main() {
	// The struc url as below:
	// [scheme://][userinfo@]host:[:port][/path][?query]["#fragment"]
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
	fmt.Println("Raw Query:", parsedURL.RawPath)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURL1 := "https://example.com/path?name=John&age=30"
	parsedURL1, err := url.Parse(rawURL1)

	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}

	queryParams := parsedURL1.Query()
	// Return a map (key, value) pair
	fmt.Println(queryParams)

	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	// Build URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}
	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "39")
	baseURL.RawQuery = query.Encode()
	fmt.Println("Built URL: ", baseURL.String())

	// Create values objects
	values := url.Values{}
	// Add key-value pairs to the values object
	values.Add("name", "Tho Le")
	values.Add("age", "30")
	values.Add("city", "HCMC")
	values.Add("country", "Vietnam")
	// Encode into URL format
	encodedQuery := values.Encode()

	fmt.Println(encodedQuery)

	// Build a URL
	baseURL1 := "https://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery
	fmt.Println(fullURL)

}
