package main


//ProviceList
import (
	"crawler/engine"
	"crawler/weike/parser"
)

func main()  {
	engine.Run(engine.Request{
		Url:"http://weike46.com/",
		ParserFunc: parser.ProviceList,
	})
}
