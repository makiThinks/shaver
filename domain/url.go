// url: domain model
package domain

// core: change url to a short one
func ToShortUrl(realUrl string) (string, error) {
	// check if exsit, return old one
	existShort := getExistShortUrl(realUrl)
	if existShort == "" {
		return "newFakeShortUrl", nil
	} else {
		return existShort, nil
	}
}

// core:ge real url from the short one
// TODO: add real logic
func ToRealUrl(shortUrl string) (string, error) {
	return "fakeRealUrl", nil
}

// TODO: get existed url if existed
func getExistShortUrl(realUrl string) string {
	if (len(realUrl) & 1) == 1 {
		return "existFakeShortUrl"
	} else {
		return ""
	}
}
