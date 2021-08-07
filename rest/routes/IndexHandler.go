package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	// "github.com/zapproject/pythia/common"
	// "github.com/zapproject/pythia/db"
	// "github.com/zapproject/pythia/rest"
)

//BalanceHandler handles balance requests
type IndexHandler struct {
}

func collect(symbol string) (string, string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0"),
		colly.MaxDepth(1),
	)

	//   c.OnResponse(func(r *colly.Response){
	// 	fmt.Println("Visited: ", r.Request.URL)
	//   })
	pair := ""
	value := ""

	switch symbol {
	case "XCD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)", func(e *colly.HTMLElement) {
			pair = "XCD/USD: "
			value = e.Text
		})

		c.Visit("https://www.exchangerates.org.uk/East-Caribbean-Dollar-XCD-currency-table.html")

	case "JMD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)", func(e *colly.HTMLElement) {
			pair = "JMD/USD: "
			value = e.Text
		})

		c.Visit("https://www.exchangerates.org.uk/Jamaican-Dollar-JMD-currency-table.html")

	case "BSD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BSD/USD: "
				value = e.Text + ".00000"
			})

		c.Visit("https://www.exchangerates.org.uk/Bahamian-Dollar-BSD-currency-table.html")

	case "BBD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)", func(e *colly.HTMLElement) {
			pair = "BBD/USD: "
			value = e.Text + ".00000"
		})

		c.Visit("https://www.exchangerates.org.uk/Bermuda-Dollar-BMD-currency-table.html")

	case "TTD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "TTD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Trinidad-Tobago-Dollar-TTD-currency-table.html")

	case "BZD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BZD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Belize-Dollar-BZD-currency-table.html")

	case "HTG":
		c.OnHTML("tr.colone:nth-child(14) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "HTG/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Haiti-Gourde-HTG-currency-table.html")

	case "SRD":
		c.OnHTML("#currency-notes > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2)",
			func(e *colly.HTMLElement) {
				pair = "SRD/USD: "
				value = e.Text
			})

		c.Visit("https://www.centralbank.org.bz/")

	case "BMD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BMD/USD: "
				value = e.Text + ".00000"
			})

		c.Visit("https://www.exchangerates.org.uk/Bermuda-Dollar-BMD-currency-table.html")

	case "KYD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "KYD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Cayman-Islands-Dollar-KYD-currency-table.html")
	}
	return pair, value
}

func BuildEndpoints(router *Router, handler *IndexHandler) {
	router.AddRoute("/xcd", handler)
	router.AddRoute("/ttd", handler)
	router.AddRoute("/jmd", handler)
	router.AddRoute("/bbd", handler)
	router.AddRoute("/bsd", handler)
	router.AddRoute("/bzd", handler)
	router.AddRoute("/htg", handler)
	router.AddRoute("/srd", handler)
	router.AddRoute("/bmd", handler)
	router.AddRoute("/kyd", handler)
}

//Incoming implementation for  handler
func (h *IndexHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	// DB := ctx.Value(common.DBContextKey).(db.DB)

	URL := req.URL
	symbol := URL.Path
	slash := strings.Index(symbol, "/")
	if slash != -1 {
		symbol = symbol[slash+1:]
	}

	p, v := collect(strings.ToUpper(symbol))

	// v, err := DB.Get(db.GasKey)
	// if err != nil {
	// 	log.Printf("Problem reading Gas from DB: %v\n", err)
	// 	return 500, `{"error": "Could not read Gas from DB"}`
	// }
	return 200, fmt.Sprintf(
		`{
			"pair": "%s",
			"value": "%s"
		 }`,
		p, v)
}
