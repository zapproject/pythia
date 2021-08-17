package util

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const API = "json(https://api.currencyscoop.com/v1/convert?from=USD&to=XCD&amount=1&api_key=${CS_KEY}).response.value"

func TestJSONParser(t *testing.T) {
	res, err := testFetch(1000, API)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) > 0 {
		t.Logf("Parsed json properly, size: %d", len(res))
		fmt.Printf("XCD/USD: %v", res)

	} else {
		t.Fatalf("Json not parsed properly: %v\n", res)
		assert.LessOrEqual(t, math.Abs((res[0] - 2.7)), 0.01)
	}
}

func testFetch(_granularity uint, queryString string) ([]float64, error) {

	url, args := ParseQueryString(queryString)
	resp, _ := http.Get(url)

	input, _ := ioutil.ReadAll(resp.Body)
	return ParsePayload(input, args)
}
