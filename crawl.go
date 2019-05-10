package vanilla

import (
	"golang.org/x/net/html"
	"net/http"
	"scrapper/util"
	"strings"
	"time"
)

func Crawl(url string) *http.Response {
	client:= &http.Client{
		Timeout: 30 * time.Second,
	}

	request:= createRequest(url)
	response, err := client.Do(request)
	util.HandleError(err)

	return response

}
func createRequest(url string) *http.Request {
	request, err := http.NewRequest("GET", url, nil)
	util.HandleError(err)
	request.Header.Set("User-Agent", "Not Firefox")
	return request
}


func FetchDetailUrl(response http.Response) [] string {
	body:= response.Body
	defer body.Close()
	z := html.NewTokenizer(body)
	dtlHrefs := make([] string, 0)
	for  {
		tt := z.Next()
		switch tt {
		case  html.ErrorToken :
			return dtlHrefs
		case html.StartTagToken :
			t := z.Token()
			isLi := t.Data == "li"
			if isLi  && strings.Contains( t.String(), "cntanr"){
				for _, attr := range t.Attr {
					if attr.Key == "data-href" {
						dtlHrefs = append(dtlHrefs, attr.Val)
					}
				}
			}
		}
	}
}


