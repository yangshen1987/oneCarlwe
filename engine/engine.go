package engine

import (
	"crawler/fetch"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,m := range seeds{
		requests = append(requests, m)
	}
	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetch.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher wrong url is %s err is %s", r.Url, err.Error())
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items{
			log.Printf("get it item %s",item)
		}

	}
}
