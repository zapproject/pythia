package tracker

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/zapproject/pythia/apiOracle"
	"github.com/zapproject/pythia/config"
)

var PSRs = map[int]ValueGenerator{
	1: &SingleSymbol{symbol: "ZAP/USD", granularity: 1000000, transform: MedianAt},
	2: &SingleSymbol{symbol: "BTC/USD", granularity: 1000000, transform: MedianAt},
	3: &SingleSymbol{symbol: "ETH/USD", granularity: 1000000, transform: MedianAt},
	4: &SingleSymbol{symbol: "LTC/USD", granularity: 1000000, transform: MedianAt},
	5: &SingleSymbol{symbol: "XRP/USD", granularity: 1000000, transform: MedianAt},
	6: &SingleSymbol{symbol: "BNB/USD", granularity: 1000000, transform: MedianAt},

	7:  &SingleSymbol{symbol: "DOGE/USD", granularity: 1000000, transform: MedianAt},
	8:  &SingleSymbol{symbol: "BCH/USD", granularity: 1000000, transform: MedianAt},
	9:  &SingleSymbol{symbol: "BSV/USD", granularity: 1000000, transform: MedianAt},
	10: &SingleSymbol{symbol: "EOS/USD", granularity: 1000000, transform: MedianAt},
	11: &SingleSymbol{symbol: "ETC/USD", granularity: 1000000, transform: MedianAt},

	12: &SingleSymbol{symbol: "ADA/USD", granularity: 1000000, transform: MedianAt},
	13: &SingleSymbol{symbol: "DOT/USD", granularity: 1000000, transform: MedianAt},
	14: &SingleSymbol{symbol: "UNI/USD", granularity: 1000000, transform: MedianAt},
	15: &SingleSymbol{symbol: "LINK/USD", granularity: 1000000, transform: MedianAt},

	16: &SingleSymbol{symbol: "EUR/USD", granularity: 1000000, transform: MedianAt},
	17: &SingleSymbol{symbol: "CNY/USD", granularity: 1000000, transform: MedianAt},
	18: &SingleSymbol{symbol: "JPY/USD", granularity: 1000000, transform: MedianAt},
	19: &SingleSymbol{symbol: "INR/USD", granularity: 1000000, transform: MedianAt},
	20: &SingleSymbol{symbol: "GBP/USD", granularity: 1000000, transform: MedianAt},

	21: &SingleSymbol{symbol: "CAD/USD", granularity: 1000000, transform: MedianAt},
	22: &SingleSymbol{symbol: "RUB/USD", granularity: 1000000, transform: MedianAt},
	23: &SingleSymbol{symbol: "KRW/USD", granularity: 1000000, transform: MedianAt},
	24: &SingleSymbol{symbol: "AUD/USD", granularity: 1000000, transform: MedianAt},

	25: &SingleSymbol{symbol: "TRY/USD", granularity: 1000000, transform: MedianAt},
	26: &SingleSymbol{symbol: "CHF/USD", granularity: 1000000, transform: MedianAt},
	27: &SingleSymbol{symbol: "THB/USD", granularity: 1000000, transform: MedianAt},
	28: &SingleSymbol{symbol: "SEK/USD", granularity: 1000000, transform: MedianAt},

	29: &SingleSymbol{symbol: "AED/USD", granularity: 1000000, transform: MedianAt},
	30: &SingleSymbol{symbol: "NOK/USD", granularity: 1000000, transform: MedianAt},
	31: &SingleSymbol{symbol: "SGD/USD", granularity: 1000000, transform: MedianAt},
	32: &SingleSymbol{symbol: "HKD/USD", granularity: 1000000, transform: MedianAt},

	33: &SingleSymbol{symbol: "ZAR/USD", granularity: 1000000, transform: MedianAt},
	34: &SingleSymbol{symbol: "DKK/USD", granularity: 1000000, transform: MedianAt},
	35: &SingleSymbol{symbol: "NZD/USD", granularity: 1000000, transform: MedianAt},
	36: &SingleSymbol{symbol: "XCD/USD", granularity: 1000000, transform: MedianAt},

	37: &SingleSymbol{symbol: "TTD/USD", granularity: 1000000, transform: MedianAt},
	38: &SingleSymbol{symbol: "JMD/USD", granularity: 1000000, transform: MedianAt},
	39: &SingleSymbol{symbol: "HTG/USD", granularity: 1000000, transform: MedianAt},
	40: &SingleSymbol{symbol: "BSD/USD", granularity: 1000000, transform: MedianAt},

	41: &SingleSymbol{symbol: "BMD/USD", granularity: 1000000, transform: MedianAt},
	42: &SingleSymbol{symbol: "KYD/USD", granularity: 1000000, transform: MedianAt},
	43: &SingleSymbol{symbol: "BBD/USD", granularity: 1000000, transform: MedianAt},
	44: &SingleSymbol{symbol: "SRD/USD", granularity: 1000000, transform: MedianAt},
	45: &SingleSymbol{symbol: "BZD/USD", granularity: 1000000, transform: MedianAt},

	46: &SingleSymbol{symbol: "UAH/USD", granularity: 1000000, transform: MedianAt},
	47: &SingleSymbol{symbol: "UYU/USD", granularity: 1000000, transform: MedianAt},
	48: &SingleSymbol{symbol: "VUV/USD", granularity: 1000000, transform: MedianAt},
	49: &SingleSymbol{symbol: "VEF/USD", granularity: 1000000, transform: MedianAt},
	50: &SingleSymbol{symbol: "XOF/USD", granularity: 1000000, transform: MedianAt},
	51: &SingleSymbol{symbol: "YER/USD", granularity: 1000000, transform: MedianAt},

	52:  &SingleSymbol{symbol: "ZMK/USD", granularity: 1000000, transform: MedianAt},
	53:  &SingleSymbol{symbol: "UGX/USD", granularity: 1000000, transform: MedianAt},
	54:  &SingleSymbol{symbol: "TND/USD", granularity: 1000000, transform: MedianAt},
	55:  &SingleSymbol{symbol: "TOP/USD", granularity: 1000000, transform: MedianAt},
	56:  &SingleSymbol{symbol: "TZS/USD", granularity: 1000000, transform: MedianAt},
	57:  &SingleSymbol{symbol: "TWD/USD", granularity: 1000000, transform: MedianAt},
	58:  &SingleSymbol{symbol: "WST/USD", granularity: 1000000, transform: MedianAt},
	59:  &SingleSymbol{symbol: "SAR/USD", granularity: 1000000, transform: MedianAt},
	60:  &SingleSymbol{symbol: "SCR/USD", granularity: 1000000, transform: MedianAt},
	61:  &SingleSymbol{symbol: "SLL/USD", granularity: 1000000, transform: MedianAt},
	62:  &SingleSymbol{symbol: "SKK/USD", granularity: 1000000, transform: MedianAt},
	63:  &SingleSymbol{symbol: "SIT/USD", granularity: 1000000, transform: MedianAt},
	64:  &SingleSymbol{symbol: "SBD/USD", granularity: 1000000, transform: MedianAt},
	65:  &SingleSymbol{symbol: "LKR/USD", granularity: 1000000, transform: MedianAt},
	66:  &SingleSymbol{symbol: "SDD/USD", granularity: 1000000, transform: MedianAt},
	67:  &SingleSymbol{symbol: "SZL/USD", granularity: 1000000, transform: MedianAt},
	68:  &SingleSymbol{symbol: "NAD/USD", granularity: 1000000, transform: MedianAt},
	69:  &SingleSymbol{symbol: "NPR/USD", granularity: 1000000, transform: MedianAt},
	70:  &SingleSymbol{symbol: "ANG/USD", granularity: 1000000, transform: MedianAt},
	71:  &SingleSymbol{symbol: "NIO/USD", granularity: 1000000, transform: MedianAt},
	72:  &SingleSymbol{symbol: "NGN/USD", granularity: 1000000, transform: MedianAt},
	73:  &SingleSymbol{symbol: "OMR/USD", granularity: 1000000, transform: MedianAt},
	74:  &SingleSymbol{symbol: "XPF/USD", granularity: 1000000, transform: MedianAt},
	75:  &SingleSymbol{symbol: "PKR/USD", granularity: 1000000, transform: MedianAt},
	76:  &SingleSymbol{symbol: "PAB/USD", granularity: 1000000, transform: MedianAt},
	77:  &SingleSymbol{symbol: "PGK/USD", granularity: 1000000, transform: MedianAt},
	78:  &SingleSymbol{symbol: "PYG/USD", granularity: 1000000, transform: MedianAt},
	79:  &SingleSymbol{symbol: "PEN/USD", granularity: 1000000, transform: MedianAt},
	80:  &SingleSymbol{symbol: "PHP/USD", granularity: 1000000, transform: MedianAt},
	81:  &SingleSymbol{symbol: "PLN/USD", granularity: 1000000, transform: MedianAt},
	82:  &SingleSymbol{symbol: "QAR/USD", granularity: 1000000, transform: MedianAt},
	83:  &SingleSymbol{symbol: "RON/USD", granularity: 1000000, transform: MedianAt},
	84:  &SingleSymbol{symbol: "RWF/USD", granularity: 1000000, transform: MedianAt},
	85:  &SingleSymbol{symbol: "KZT/USD", granularity: 1000000, transform: MedianAt},
	86:  &SingleSymbol{symbol: "KES/USD", granularity: 1000000, transform: MedianAt},
	87:  &SingleSymbol{symbol: "KWD/USD", granularity: 1000000, transform: MedianAt},
	88:  &SingleSymbol{symbol: "LVL/USD", granularity: 1000000, transform: MedianAt},
	89:  &SingleSymbol{symbol: "LBP/USD", granularity: 1000000, transform: MedianAt},
	90:  &SingleSymbol{symbol: "LSL/USD", granularity: 1000000, transform: MedianAt},
	91:  &SingleSymbol{symbol: "LTL/USD", granularity: 1000000, transform: MedianAt},
	92:  &SingleSymbol{symbol: "MOP/USD", granularity: 1000000, transform: MedianAt},
	93:  &SingleSymbol{symbol: "MKD/USD", granularity: 1000000, transform: MedianAt},
	94:  &SingleSymbol{symbol: "MWK/USD", granularity: 1000000, transform: MedianAt},
	95:  &SingleSymbol{symbol: "MYR/USD", granularity: 1000000, transform: MedianAt},
	96:  &SingleSymbol{symbol: "MVR/USD", granularity: 1000000, transform: MedianAt},
	97:  &SingleSymbol{symbol: "MRO/USD", granularity: 1000000, transform: MedianAt},
	98:  &SingleSymbol{symbol: "MUR/USD", granularity: 1000000, transform: MedianAt},
	99:  &SingleSymbol{symbol: "MXN/USD", granularity: 1000000, transform: MedianAt},
	100: &SingleSymbol{symbol: "MDL/USD", granularity: 1000000, transform: MedianAt},
	101: &SingleSymbol{symbol: "MNT/USD", granularity: 1000000, transform: MedianAt},
	102: &SingleSymbol{symbol: "MAD/USD", granularity: 1000000, transform: MedianAt},
	103: &SingleSymbol{symbol: "HNL/USD", granularity: 1000000, transform: MedianAt},
	104: &SingleSymbol{symbol: "HUF/USD", granularity: 1000000, transform: MedianAt},
	105: &SingleSymbol{symbol: "ISK/USD", granularity: 1000000, transform: MedianAt},
	106: &SingleSymbol{symbol: "IDR/USD", granularity: 1000000, transform: MedianAt},
	107: &SingleSymbol{symbol: "IRR/USD", granularity: 1000000, transform: MedianAt},
	108: &SingleSymbol{symbol: "IQD/USD", granularity: 1000000, transform: MedianAt},
	109: &SingleSymbol{symbol: "ILS/USD", granularity: 1000000, transform: MedianAt},
	110: &SingleSymbol{symbol: "JOD/USD", granularity: 1000000, transform: MedianAt},
	111: &SingleSymbol{symbol: "GTQ/USD", granularity: 1000000, transform: MedianAt},
	112: &SingleSymbol{symbol: "GNF/USD", granularity: 1000000, transform: MedianAt},
	113: &SingleSymbol{symbol: "ALL/USD", granularity: 1000000, transform: MedianAt},
	114: &SingleSymbol{symbol: "BGN/USD", granularity: 1000000, transform: MedianAt},
	115: &SingleSymbol{symbol: "ARS/USD", granularity: 1000000, transform: MedianAt},
	116: &SingleSymbol{symbol: "AWG/USD", granularity: 1000000, transform: MedianAt},
	117: &SingleSymbol{symbol: "BAM/USD", granularity: 1000000, transform: MedianAt},
	118: &SingleSymbol{symbol: "AZN/USD", granularity: 1000000, transform: MedianAt},
	119: &SingleSymbol{symbol: "BYN/USD", granularity: 1000000, transform: MedianAt},
	120: &SingleSymbol{symbol: "AFN/USD", granularity: 1000000, transform: MedianAt},
	121: &SingleSymbol{symbol: "BOB/USD", granularity: 1000000, transform: MedianAt},
	122: &SingleSymbol{symbol: "BWP/USD", granularity: 1000000, transform: MedianAt},
	123: &SingleSymbol{symbol: "BRL/USD", granularity: 1000000, transform: MedianAt},
	124: &SingleSymbol{symbol: "BND/USD", granularity: 1000000, transform: MedianAt},
	125: &SingleSymbol{symbol: "KHR/USD", granularity: 1000000, transform: MedianAt},
	126: &SingleSymbol{symbol: "CLP/USD", granularity: 1000000, transform: MedianAt},
	127: &SingleSymbol{symbol: "COP/USD", granularity: 1000000, transform: MedianAt},
	128: &SingleSymbol{symbol: "CRC/USD", granularity: 1000000, transform: MedianAt},
	129: &SingleSymbol{symbol: "HRK/USD", granularity: 1000000, transform: MedianAt},
	130: &SingleSymbol{symbol: "CUP/USD", granularity: 1000000, transform: MedianAt},
	131: &SingleSymbol{symbol: "CZK/USD", granularity: 1000000, transform: MedianAt},
	132: &SingleSymbol{symbol: "DOP/USD", granularity: 1000000, transform: MedianAt},
	133: &SingleSymbol{symbol: "EGP/USD", granularity: 1000000, transform: MedianAt},
	134: &SingleSymbol{symbol: "SVC/USD", granularity: 1000000, transform: MedianAt},
	135: &SingleSymbol{symbol: "FKP/USD", granularity: 1000000, transform: MedianAt},
	136: &SingleSymbol{symbol: "FJD/USD", granularity: 1000000, transform: MedianAt},
	137: &SingleSymbol{symbol: "GHS/USD", granularity: 1000000, transform: MedianAt},
	138: &SingleSymbol{symbol: "GIP/USD", granularity: 1000000, transform: MedianAt},
	139: &SingleSymbol{symbol: "GGP/USD", granularity: 1000000, transform: MedianAt},
	140: &SingleSymbol{symbol: "GYD/USD", granularity: 1000000, transform: MedianAt},
	141: &SingleSymbol{symbol: "TVD/USD", granularity: 1000000, transform: MedianAt},
	142: &SingleSymbol{symbol: "UZS/USD", granularity: 1000000, transform: MedianAt},
	143: &SingleSymbol{symbol: "VND/USD", granularity: 1000000, transform: MedianAt},
	144: &SingleSymbol{symbol: "SOS/USD", granularity: 1000000, transform: MedianAt},
	145: &SingleSymbol{symbol: "SYP/USD", granularity: 1000000, transform: MedianAt},
}

//these weight functions map values of x between 0 (brand new) and 1 (old) to weights between 0 and 1
//also returns the integral of the weight over the range [0,1]
//weights the oldest data (1) as being 1/3 as important (1/e)
func ExpDecay(x float64) (float64, float64) {
	return 1 / math.Exp(x), 0.63212
}

//weights the oldest data at 0
func LinearDecay(x float64) (float64, float64) {
	return 1 - x, 0.5
}

//weights all data in the time interval evenly
func NoDecay(x float64) (float64, float64) {
	return 1, 1
}

func TimeWeightedAvg(interval time.Duration, weightFn func(float64) (float64, float64)) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
		cfg := config.GetConfig()
		sum := 0.0
		weightSum := 0.0
		numVals := 0
		minTime := at
		maxVolume := 0.0
		for _, api := range apis {
			values := apiOracle.GetRequestValuesForTime(api.Identifier, at, interval)
			for _, v := range values {
				normDelta := at.Sub(v.Created).Seconds() / interval.Seconds()
				weight, _ := weightFn(normDelta)
				sum += v.Price * weight
				weightSum += weight
				numVals += 1
				if minTime.Sub(v.Created).Seconds() > 0 {
					minTime = v.Created
				}
				if v.Volume > maxVolume {
					maxVolume = v.Volume
				}
			}
		}

		// trackerSleepCycle, err := time.ParseDuration(cfg.TrackerSleepCycle)

		// if err != nil {

		// 	fmt.Println(err)
		// }

		// number of APIs * rate * interval
		maxWeight := float64(len(apis)) * (1 / cfg.TrackerSleepCycle.Seconds()) * interval.Seconds()
		//average weight is the integral of the weight fn over [0,1]
		_, avgWeight := weightFn(0)
		targetWeight := maxWeight * avgWeight

		var result apiOracle.PriceInfo
		result.Price = sum / weightSum

		//use the highest volume seen over all values. works well when the time averaging window is equal to the interval of volume reporting
		// ie, 24 hour average on an api that returns 24hr volume
		result.Volume = maxVolume
		// if math.Min(weightSum/targetWeight, 1.0) < .5{
		// 	values,_ := apiOracle.GetNearestTwoRequestValue(apis[0].Identifier, at)
		// 	fmt.Println("not enough data for time series, series starts : ", values.Created)
		// }
		//fmt.Println("Time Weighted: ", result)
		return result, math.Min(weightSum/targetWeight, 1.0)
	}
}

func VolumeWeightedAPIs(processor IndexProcessor) IndexProcessor {
	return func(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
		var results []apiOracle.PriceInfo
		totalConfidence := 0.0
		for _, api := range apis {
			value, confidence := processor([]*IndexTracker{api}, at)
			//fmt.Println("vwAPI's : ",value, "  : ", confidence)
			if confidence > 0 {
				results = append(results, value)
				totalConfidence += confidence
			}
		}
		return VolumeWeightedAvg(results), totalConfidence / float64(len(results))
	}
}

func getLatest(apis []*IndexTracker, at time.Time) ([]apiOracle.PriceInfo, float64) {
	var values []apiOracle.PriceInfo
	totalConf := 0.0

	for _, api := range apis {

		b, _ := apiOracle.GetNearestTwoRequestValue(api.Identifier, at)

		if b != nil {
			//penalize values more than 5 minutes old
			totalConf += math.Min(5/at.Sub(b.Created).Minutes(), 1.0)
			values = append(values, b.PriceInfo)
		}
	}

	return values, totalConf / float64(len(apis))
}

func MedianAt(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	values, confidence := getLatest(apis, at)

	if confidence == 0 {
		return apiOracle.PriceInfo{}, 0
	}

	return Median(values), confidence
}

func ManualEntry(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {

	vals, confidence := getLatest(apis, at)

	if confidence == 0 {
		return apiOracle.PriceInfo{}, 0
	}

	for _, val := range vals {
		// fmt.Println(int64(val.Volume),time.Now().Unix())
		if int64(val.Volume) < clck.Now().Unix() {
			fmt.Println("Pulled Timestamp: ", val.Volume)
			fmt.Println("Warning: Manual Data Entry is expired, please update")
			return apiOracle.PriceInfo{}, 0
		}
	}
	return Median(vals), confidence
}

func MedianAtEOD(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {
	now := clck.Now().UTC()

	fmt.Println(now)
	d := 24 * time.Hour
	eod := now.Truncate(d)

	fmt.Println(MedianAt(apis, eod))

	return MedianAt(apis, eod)
}

func Median(values []apiOracle.PriceInfo) apiOracle.PriceInfo {
	var result apiOracle.PriceInfo
	// fmt.Println("Values: ", values)
	sort.Slice(values, func(i, j int) bool {
		return values[i].Price < values[j].Price
	})

	result.Price = values[len(values)/2].Price
	for _, val := range values {
		result.Volume += val.Volume
	}
	// fmt.Println("Result: ", result)
	return result
}

func MeanAt(apis []*IndexTracker, at time.Time) (apiOracle.PriceInfo, float64) {

	values, confidence := getLatest(apis, at)

	if confidence == 0 {
		return apiOracle.PriceInfo{}, 0
	}

	return Mean(values), confidence
}

func Mean(vals []apiOracle.PriceInfo) apiOracle.PriceInfo {
	var result apiOracle.PriceInfo
	priceSum := 0.0
	volSum := 0.0
	for _, val := range vals {
		priceSum += val.Price
		volSum += val.Volume
	}
	result.Price = priceSum / float64(len(vals))
	result.Volume = volSum
	return result
}

func VolumeWeightedAvg(vals []apiOracle.PriceInfo) apiOracle.PriceInfo {
	priceSum := 0.0
	volSum := 0.0
	for _, val := range vals {
		priceSum += val.Price * val.Volume
		volSum += val.Volume
	}
	if volSum > 0 {
		return apiOracle.PriceInfo{Price: priceSum / volSum, Volume: volSum}
	}
	//if there was no volume data, just do a normal average instead
	priceSum = 0
	for _, val := range vals {
		priceSum += val.Price
	}
	return apiOracle.PriceInfo{Price: priceSum / float64(len(vals)), Volume: 0}
}
