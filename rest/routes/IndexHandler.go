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
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "XCD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/East-Caribbean-Dollar-XCD-currency-table.html")

	case "JMD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "JMD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Jamaican-Dollar-JMD-currency-table.html")

	case "BSD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BSD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Bahamian-Dollar-BSD-currency-table.html")

	case "BBD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BBD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Barbadian-Dollar-BBD-currency-table.html")

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
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Bermuda-Dollar-BMD-currency-table.html")

	case "KYD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "KYD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Cayman-Islands-Dollar-KYD-currency-table.html")

	case "AUD":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "AUD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Australian-Dollar-AUD-currency-table.html")

	case "GBP":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "GBP/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/British-Pound-GBP-currency-table.html")

	case "CAD":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "CAD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Canadian-Dollar-CAD-currency-table.html")

	case "DKK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "DKK/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Danish-Krone-DKK-currency-table.html")

	case "EUR":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "EUR/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Euro-EUR-currency-table.html")

	case "HKD":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "HKD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Hong-Kong-Dollar-HKD-currency-table.html")

	case "JPY":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "JPY/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Japanese-Yen-JPY-currency-table.html")

	case "NZD":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "NZD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/New-Zealand-Dollar-NZD-currency-table.html")

	case "NOK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "NOK/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Norwegian-Krone-NOK-currency-table.html")

	case "ZAR":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "ZAR/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/South-African-Rand-ZAR-currency-table.html")

	case "SEK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SEK/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Swedish-Krona-SEK-currency-table.html")

	case "CHF":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "CHF/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Swiss-Franc-CHF-currency-table.html")

	case "TRY":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "TRY/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Turkish-Lira-TRY-currency-table.html")

	case "AED":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "AED/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/United-Arab-Emirates-Dirham-AED-currency-table.html")

	case "CNY":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "CNY/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Chinese-Yuan-CNY-currency-table.html")

	case "THB":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "THB/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Thai-Baht-THB-currency-table.html")

	case "INR":
		c.OnHTML(".colone:nth-child(12) strong",
			func(e *colly.HTMLElement) {
				pair = "INR/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Indian-Rupee-INR-currency-table.html")
	case "SGD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SGD/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Singapore-Dollar-SGD-currency-table.html")

	case "RUB":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "RUB/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Russian-Rouble-RUB-currency-table.html")
	case "KRW":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "KRW/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/South-Korean-Won-KRW-currency-table.html")

	case "ALL":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "ALL/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Albanian-Lek-ALL-currency-table.html")

	case "BGN":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BGN/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Bulgarian-Lev-BGN-currency-table.html")

	case "ARS":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "ARS/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Argentine-Peso-ARS-currency-table.html")

	case "AWG":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "AWG/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Aruba-Florin-AWG-currency-table.html")

	case "BAM":
		c.OnHTML(".table > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(2) > a:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BAM/USD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/currencies/bam-bosnia-herzegovina-convertible-mark")

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
	router.AddRoute("/aud", handler)
	router.AddRoute("/gbp", handler)
	router.AddRoute("/cad", handler)
	router.AddRoute("/dkk", handler)
	router.AddRoute("/eur", handler)
	router.AddRoute("/hkd", handler)
	router.AddRoute("/jpy", handler)
	router.AddRoute("/nzd", handler)
	router.AddRoute("/nok", handler)
	router.AddRoute("/zar", handler)
	router.AddRoute("/sek", handler)
	router.AddRoute("/chf", handler)
	router.AddRoute("/try", handler)
	router.AddRoute("/aed", handler)
	router.AddRoute("/cny", handler)
	router.AddRoute("/thb", handler)
	router.AddRoute("/inr", handler)
	router.AddRoute("/sgd", handler)
	router.AddRoute("/rub", handler)
	router.AddRoute("/krw", handler)
	router.AddRoute("/all", handler)
	router.AddRoute("/bgn", handler)
	router.AddRoute("/ars", handler)
	router.AddRoute("/awg", handler)
	router.AddRoute("/bam", handler)

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
