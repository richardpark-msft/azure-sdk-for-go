package internal

import (
	"fmt"
	"net/url"
)

// ConstructAtomPath adds the proper parameters for skip and top
// This is common for the list operations for queues, topics and subscriptions.
func ConstructAtomPath(baseUrl string, skip int, top int) string {
	values := url.Values{}

	if skip > 0 {
		values.Add("$skip", fmt.Sprintf("%d", skip))
	}

	if top > 0 {
		values.Add("$top", fmt.Sprintf("%d", top))
	}

	queryParams := values.Encode()

	if len(queryParams) == 0 {
		return baseUrl
	}

	return fmt.Sprintf("%s?%s", baseUrl, queryParams)
}
