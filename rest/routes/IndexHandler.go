package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

//BalanceHandler handles balance requests
type IndexHandler struct {
}

func collect(symbol string) (string, string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0"),
		colly.MaxDepth(1),
	)

	pair := ""
	value := ""
	end := strings.Index(e.Text, " ")

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

	case "KRW":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "KRW/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/South-Korean-Won-KRW-currency-table.html")

	case "NAD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "NAD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Namibian-Dollar-NAD-currency-table.html")
	case "NPR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "NPR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Nepalese-Rupee-NPR-currency-table.html")
	case "ANG":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "ANG/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Neth-Antilles-Guilder-ANG-currency-table.html")

	case "NIO":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "NIO/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Nicaragua-Cordoba-NIO-currency-table.html")
	case "NGN":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "NGN/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Nigerian-Naira-NGN-currency-table.html")
	case "OMR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "OMR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Omani-Rial-OMR-currency-table.html")
	case "XPF":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "XPF/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Pacific-Franc-XPF-currency-table.html")
	case "PKR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PKR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Pakistani-Rupee-PKR-currency-table.html")
	case "PAB":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PAB/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Panamanian-Balboa-PAB-currency-table.html")
	case "PGK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PGK/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Papua-New-Guinea-Kina-PGK-currency-table.html")
	case "PYG":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PYG/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Paraguayan-Guarani-PYG-currency-table.html")
	case "PEN":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PEN/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Peruvian-Nuevo-Sol-PEN-currency-table.html")
	case "PHP":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PHP/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Philippine-Peso-PHP-currency-table.html")
	case "PLN":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "PLN/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Polish-Zloty-PLN-currency-table.html")
	case "QAR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "QAR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Qatari-Riyal-QAR-currency-table.html")
	case "RON":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "RON/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Romanian-Leu-RON-currency-table.html")
	case "RUB":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "RUB/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Russian-Rouble-RUB-currency-table.html")
	case "RWF":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "RWF/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Rwanda-Franc-RWF-currency-table.html")

	case "WST":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "WST/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Samoa-Tala-WST-currency-table.html")
	case "SAR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SAR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Saudi-Riyal-SAR-currency-table.html")
	case "SCR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SCR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Seychelles-Rupee-SCR-currency-table.html")
	case "SLL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SLL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Sierra-Leone-Leone-SLL-currency-table.html")
	case "SKK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SKK/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Slovak-Koruna-SKK-currency-table.html")
	case "SIT":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SIT/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Slovenian-Tolar-SIT-currency-table.html")
	case "SBD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SBD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Solomon-Islands-Dollar-SBD-currency-table.html")
	case "LKR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "LKR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Sri-Lankan-Rupee-LKR-currency-table.html")
	case "SDD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SDD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Sudanese-Dinar-SDD-currency-table.html")
	case "SZL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "SZL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Swaziland-Lilageni-SZL-currency-table.html")

	case "TWD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "TWD/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Taiwan-Dollar-TWD-currency-table.html")
	case "TZS":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "TZS/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Tanzanian-Shilling-TZS-currency-table.html")
	case "TOP":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "TOP/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Tonga-Paanga-TOP-currency-table.html")
	case "TND":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "TND/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Tunisian-Dinar-TND-currency-table.html")
	case "UGX":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "UGX/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Ugandan-Shilling-UGX-currency-table.html")
	case "UAH":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "UAH/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Ukraine-Hryvnia-UAH-currency-table.html")

	case "UYU":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "UYU/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Uruguayan-New-Peso-UYU-currency-table.html")

	case "VUV":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "VUV/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Venezuelan-bolivar-VUV-currency-table.html")
	case "VEF":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "VEF/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Venezuelan-bolivar-VEF-currency-table.html")

	case "XOF":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "XOF/USD: "
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/West-African-CFA-franc-XOF-currency-table.html")

	case "YER":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "YER/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Yemen-Riyal-YER-currency-table.html")

	case "ZMK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "ZMK/USD: "
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Zambian-Kwacha-ZMK-currency-table.html")

	case "KZT":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "KZT/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Kazakhstan-Tenge-KZT-currency-table.html")

	case "KES":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "KES/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Kenyan-Shilling-KES-currency-table.html")

	case "KWD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "KWD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Kuwaiti-Dinar-KWD-currency-table.html")

	case "LVL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "LVL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Latvian-Lats-LVL-currency-table.html")

	case "LBP":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "LBP/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Lebanese-Pound-LBP-currency-table.html")

	case "LSL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "LSL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Lesotho-Loti-LSL-currency-table.html")

	case "LTL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "LTL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Lithuanian-Litas-LTL-currency-table.html")

	case "MOP":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MOP/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Macau-Pataca-MOP-currency-table.html")

	case "MKD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MKD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Macedonian-Denar-MKD-currency-table.html")

	case "MWK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MWK/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Malawi-Kwacha-MWK-currency-table.html")

	case "MYR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MYR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Malaysian-Ringgit-MYR-currency-table.html")

	case "MVR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MVR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Maldives-Rufiyaa-MVR-currency-table.html")

	case "MRO":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MRO/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Mauritania-Ougulya-MRO-currency-table.html")

	case "MUR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MUR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Mauritius-Rupee-MUR-currency-table.html")

	case "MXN":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MXN/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Mexican-Peso-MXN-currency-table.html")

	case "MDL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MDL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Moldovan-Leu-MDL-currency-table.html")

	case "MNT":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MNT/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Mongolian-Tugrik-MNT-currency-table.html")

	case "MAD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "MAD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Moroccan-Dirham-MAD-currency-table.html")

	case "HNL":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "HNL/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Honduras-Lempira-HNL-currency-table.html")
	case "HUF":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "HUF/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Hungarian-Forint-HUF-currency-table.html")
	case "ISK":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "ISK/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Icelandic-Krona-ISK-currency-table.html")
	case "IDR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "IDR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Indonesian-Rupiah-IDR-currency-table.html")
	case "IRR":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "IRR/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Iran-Rial-IRR-currency-table.html")
	case "IQD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "IQD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Iraqi-Dinar-IQD-currency-table.html")
	case "ILS":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "ILS/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Israeli-Sheqel-ILS-currency-table.html")
	case "JOD":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "JOD/USD"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Jordanian-Dinar-JOD-currency-table.html")
	case "GTQ":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "GTQ"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Guatemala-Quetzal-GTQ-currency-table.html")

	case "GNF":
		c.OnHTML(".coltwo:nth-child(13) strong",
			func(e *colly.HTMLElement) {
				pair = "GNF"
				value = e.Text
			})
		c.Visit("https://www.exchangerates.org.uk/Guinea-Franc-GNF-currency-table.html")

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

	case "AFN":
		c.OnHTML(".datatable > tbody:nth-child(1) > tr:nth-child(4) > td:nth-child(2)",
			func(e *colly.HTMLElement) {
				pair = "AFN/USD"
				value = e.Text[:end]
			})

		c.Visit("https://www.unitconverters.net/currency/afn-to-usd.htm")

	case "AZN":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "AZN/USD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/azn-to-usd-rate?amount=1")

	case "BYN":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "BYN/USD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/byn-to-usd-rate?amount=1")

	case "BOB":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BOB/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Bolivian-Boliviano-BOB-currency-table.html")

	case "BWP":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "BWP/USD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/bwp-to-usd-rate?amount=1")

	case "BRL":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BRL/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Brazilian-Real-BRL-currency-table.html")

	case "BND":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "BND/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Brunei-Dollar-BND-currency-table.html")

	case "KHR":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "KHR/USD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/khr-to-usd-rate?amount=1")

	case "CLP":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "CLP/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Chilean-Peso-CLP-currency-table.html")

	case "COP":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "COP/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Colombian-Peso-COP-currency-table.html")

	case "CRC":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "CRC/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Costa-Rica-Colon-CRC-currency-table.html")

	case "HRK":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "HRK/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Croatian-Kuna-HRK-currency-table.html")

	case "CUP":
		c.OnHTML(".datatable > tbody:nth-child(1) > tr:nth-child(4) > td:nth-child(2)",
			func(e *colly.HTMLElement) {
				pair = "CUP/USD"
				value = e.Text[:end]
			})

		c.Visit("https://www.unitconverters.net/currency/cup-to-usd.htm")

	case "CZK":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "CZK/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Czech-Koruna-CZK-currency-table.html")

	case "DOP":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "DOP/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Dominican-Peso-DOP-currency-table.html")

	case "EGP":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "EGP/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Egyptian-Pound-EGP-currency-table.html")

	case "SVC":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "SVC/USD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/El-Salvador-Colon-SVC-currency-table.html")

	case "FKP":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "FKP/USD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/fkp-to-usd-rate?amount=1")

	case "FJD":
		c.OnHTML("tr.coltwo:nth-child(13) > td:nth-child(4) > strong:nth-child(1)",
			func(e *colly.HTMLElement) {
				pair = "FJD"
				value = e.Text
			})

		c.Visit("https://www.exchangerates.org.uk/Fiji-Dollar-FJD-currency-table.html")

	case "GHS":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "GHS"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/ghs-to-usd-rate?amount=1")

	case "GIP":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "GIP"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/gip-to-usd-rate?amount=1")

	case "GGP":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "GGP"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/ggp-to-usd-rate?amount=1")

	case "GYD":
		c.OnHTML("h3.cc__source-to-target:nth-child(1) > span:nth-child(3)",
			func(e *colly.HTMLElement) {
				pair = "GYD"
				value = e.Text
			})

		c.Visit("https://wise.com/gb/currency-converter/gyd-to-usd-rate?amount=1")
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
	router.AddRoute("/uah", handler)
	router.AddRoute("/uyu", handler)
	router.AddRoute("/vuv", handler)
	router.AddRoute("/vef", handler)
	router.AddRoute("/xof", handler)
	router.AddRoute("/yer", handler)
	router.AddRoute("/zmk", handler)
	router.AddRoute("/ugx", handler)
	router.AddRoute("/tnd", handler)
	router.AddRoute("/top", handler)
	router.AddRoute("/tzs", handler)
	router.AddRoute("/twd", handler)
	router.AddRoute("/wst", handler)
	router.AddRoute("/sar", handler)
	router.AddRoute("/scr", handler)
	router.AddRoute("/sll", handler)
	router.AddRoute("/skk", handler)
	router.AddRoute("/sit", handler)
	router.AddRoute("/sbd", handler)
	router.AddRoute("/lkr", handler)
	router.AddRoute("/sdd", handler)
	router.AddRoute("/szl", handler)
	router.AddRoute("/nad", handler)
	router.AddRoute("/npr", handler)
	router.AddRoute("/ang", handler)
	router.AddRoute("/nio", handler)
	router.AddRoute("/ngn", handler)
	router.AddRoute("/omr", handler)
	router.AddRoute("/xpf", handler)
	router.AddRoute("/pkr", handler)
	router.AddRoute("/pab", handler)
	router.AddRoute("/pgk", handler)
	router.AddRoute("/pyg", handler)
	router.AddRoute("/pen", handler)
	router.AddRoute("/php", handler)
	router.AddRoute("/pln", handler)
	router.AddRoute("/qar", handler)
	router.AddRoute("/ron", handler)
	router.AddRoute("/rwf", handler)
	router.AddRoute("/kzt", handler)
	router.AddRoute("/kes", handler)
	router.AddRoute("/kwd", handler)
	router.AddRoute("/lvl", handler)
	router.AddRoute("/lbp", handler)
	router.AddRoute("/lsl", handler)
	router.AddRoute("/ltl", handler)
	router.AddRoute("/mop", handler)
	router.AddRoute("/mkd", handler)
	router.AddRoute("/mwk", handler)
	router.AddRoute("/myr", handler)
	router.AddRoute("/mvr", handler)
	router.AddRoute("/mro", handler)
	router.AddRoute("/mur", handler)
	router.AddRoute("/mxn", handler)
	router.AddRoute("/mdl", handler)
	router.AddRoute("/mnt", handler)
	router.AddRoute("/mad", handler)
	router.AddRoute("/hnl", handler)
	router.AddRoute("/huf", handler)
	router.AddRoute("/isk", handler)
	router.AddRoute("/idr", handler)
	router.AddRoute("/irr", handler)
	router.AddRoute("/iqd", handler)
	router.AddRoute("/ils", handler)
	router.AddRoute("/jod", handler)
	router.AddRoute("/gtq", handler)
	router.AddRoute("/gnf", handler)
	router.AddRoute("/all", handler)
	router.AddRoute("/bgn", handler)
	router.AddRoute("/ars", handler)
	router.AddRoute("/awg", handler)
	router.AddRoute("/bam", handler)
	router.AddRoute("/azn", handler)
	router.AddRoute("/byn", handler)
	router.AddRoute("/afn", handler)
	router.AddRoute("/bob", handler)
	router.AddRoute("/bwp", handler)
	router.AddRoute("/brl", handler)
	router.AddRoute("/bnd", handler)
	router.AddRoute("/khr", handler)
	router.AddRoute("/clp", handler)
	router.AddRoute("/cop", handler)
	router.AddRoute("/crc", handler)
	router.AddRoute("/hrk", handler)
	router.AddRoute("/cup", handler)
	router.AddRoute("/czk", handler)
	router.AddRoute("/dop", handler)
	router.AddRoute("/egp", handler)
	router.AddRoute("/svc", handler)
	router.AddRoute("/fkp", handler)
	router.AddRoute("/fjd", handler)
	router.AddRoute("/ghs", handler)
	router.AddRoute("/gip", handler)
	router.AddRoute("/ggp", handler)
	router.AddRoute("/gyd", handler)

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
