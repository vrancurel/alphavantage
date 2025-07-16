package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToAnalytics(t *testing.T) {
	var buf = `{
  "meta_data": {
    "symbols": "STOCK1,STOCK2",
    "min_dt": "2023-03-03",
    "max_dt": "2024-03-01",
    "ohlc": "Close",
    "interval": "DAILY"
  },
  "payload": {
    "RETURNS_CALCULATIONS": {
      "MIN": {
        "STOCK1": -0.09004237288135597,
        "STOCK2": -0.048020086833708175
      },
      "MAX": {
        "STOCK1": 0.2653534183082271,
        "STOCK2": 0.04692683515290419
      },
      "MEAN": {
        "STOCK1": 0.0018726725586940209,
        "STOCK2": 0.0007860999918119856
      },
      "MEDIAN": {
        "STOCK1": 0,
        "STOCK2": 0.0005495443688787738
      },
      "CUMULATIVE_RETURN": {
        "STOCK1": 0.43025540275049434,
        "STOCK2": 0.19588108390715786
      },
      "VARIANCE": {
        "STOCK1": 0.0009215484660468207,
        "STOCK2": 0.00014059454171378498
      },
      "STDDEV": {
        "STOCK1": 0.030357016751433607,
        "STOCK2": 0.011857256921977569
      },
      "MAX_DRAWDOWN": {
        "STOCK1": {
          "max_drawdown": -0.22888015717092336,
          "drawdown_range": {
            "start_drawdown": "2023-03-03",
            "end_drawdown": "2023-03-15"
          }
        },
        "STOCK2": {
          "max_drawdown": -0.08959022652074322,
          "drawdown_range": {
            "start_drawdown": "2023-07-31",
            "end_drawdown": "2023-08-07"
          }
        }
      },
      "HISTOGRAM": {
        "STOCK1": {
          "bin_count": [
            0,
            6,
            116,
            119,
            8,
            0,
            0,
            0,
            1,
            0
          ],
          "bin_edges": [
            -0.15,
            -0.1,
            -0.05,
            0,
            0.05,
            0.1,
            0.15,
            0.2,
            0.25,
            0.3,
            0.35
          ]
        },
        "STOCK2": {
          "bin_count": [
            0,
            0,
            120,
            130,
            0,
            0,
            0,
            0,
            0,
            0
          ],
          "bin_edges": [
            -0.15,
            -0.1,
            -0.05,
            0,
            0.05,
            0.1,
            0.15,
            0.2,
            0.25,
            0.3,
            0.35
          ]
        }
      },
      "AUTOCORRELATION": {
        "STOCK1": -0.09020369371098456,
        "STOCK2": 0.09237716246134561
      },
      "COVARIANCE": {
        "index": [
          "STOCK1",
          "STOCK2"
        ],
        "covariance": [
          [
            0.0009215485
          ],
          [
            4.40494e-05,
            0.0001405945
          ]
        ]
      },
      "CORRELATION": {
        "index": [
          "STOCK1",
          "STOCK2"
        ],
        "correlation": [
          [
            1
          ],
          [
            0.1223760538,
            1
          ]
        ]
      }
    }
  }
}`

	analytics, err := toAnalytics([]byte(buf))
	assert.NoError(t.Fatalf, err)
	assert.EqualStrings(t, "STOCK1,STOCK2", analytics.MetaData.Symbols)

	expectedMinSTOCK2 := -0.048020086833708175
	expectedMinSTOCK1 := -0.09004237288135597
	expectedMaxSTOCK2 := 0.04692683515290419
	expectedMaxSTOCK1 := 0.2653534183082271
	expectedMeanSTOCK2 := 0.0007860999918119856
	expectedMeanSTOCK1 := 0.0018726725586940209
	expectedMedianSTOCK2 := 0.0005495443688787738
	expectedMedianSTOCK1 := 0.0
	expectedCumulativeReturnSTOCK2 := 0.19588108390715786
	expectedCumulativeReturnSTOCK1 := 0.43025540275049434
	expectedVarianceSTOCK2 := 0.00014059454171378498
	expectedVarianceSTOCK1 := 0.0009215484660468207
	expectedStdDevSTOCK2 := 0.011857256921977569
	expectedStdDevSTOCK1 := 0.030357016751433607
	expectedMaxDrawdownSTOCK2 := -0.08959022652074322
	expectedMaxDrawdownStartSTOCK2 := "2023-07-31"
	expectedMaxDrawdownEndSTOCK2 := "2023-08-07"
	expectedMaxDrawdownSTOCK1 := -0.22888015717092336
	expectedMaxDrawdownStartSTOCK1 := "2023-03-03"
	expectedMaxDrawdownEndSTOCK1 := "2023-03-15"
	//expectedHistogramBinCountSTOCK2 := []int{0, 0, 120, 130, 0, 0, 0, 0, 0, 0}
	//expectedHistogramBinEdgesSTOCK2 := []float64{-0.15, -0.1, -0.05, 0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35}
	//expectedHistogramBinCountSTOCK1 := []int{0, 6, 116, 119, 8, 0, 0, 0, 1, 0}
	//expectedHistogramBinEdgesSTOCK1 := []float64{-0.15, -0.1, -0.05, 0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35}
	expectedAutocorrelationSTOCK2 := 0.09237716246134561
	expectedAutocorrelationSTOCK1 := -0.09020369371098456
	//expectedCorrelationSTOCK1STOCK2 := []float64{1}
	//expectedCorrelationSTOCK2STOCK1 := []float64{0.1223760538, 1}

	metaData := analytics.MetaData
	assert.EqualStrings(t, "STOCK1,STOCK2", metaData.Symbols)
	assert.EqualStrings(t, "2023-03-03", metaData.MinDt)
	assert.EqualStrings(t, "2024-03-01", metaData.MaxDt)
	assert.EqualStrings(t, "Close", metaData.OHLC)
	assert.EqualStrings(t, "DAILY", metaData.Interval)

	payload := analytics.Payload
	returnsCalculations := payload["RETURNS_CALCULATIONS"]

	min := returnsCalculations.Min
	assert.EqualFloat64(t, expectedMinSTOCK2, min["STOCK2"])
	assert.EqualFloat64(t, expectedMinSTOCK1, min["STOCK1"])

	max := returnsCalculations.Max
	assert.EqualFloat64(t, expectedMaxSTOCK2, max["STOCK2"])
	assert.EqualFloat64(t, expectedMaxSTOCK1, max["STOCK1"])

	mean := returnsCalculations.Mean
	assert.EqualFloat64(t, expectedMeanSTOCK2, mean["STOCK2"])
	assert.EqualFloat64(t, expectedMeanSTOCK1, mean["STOCK1"])

	median := returnsCalculations.Median
	assert.EqualFloat64(t, expectedMedianSTOCK2, median["STOCK2"])
	assert.EqualFloat64(t, expectedMedianSTOCK1, median["STOCK1"])

	cumulativeReturn := returnsCalculations.CumulativeReturn
	assert.EqualFloat64(t, expectedCumulativeReturnSTOCK2, cumulativeReturn["STOCK2"])
	assert.EqualFloat64(t, expectedCumulativeReturnSTOCK1, cumulativeReturn["STOCK1"])

	variance := returnsCalculations.Variance
	assert.EqualFloat64(t, expectedVarianceSTOCK2, variance["STOCK2"])
	assert.EqualFloat64(t, expectedVarianceSTOCK1, variance["STOCK1"])

	stdDev := returnsCalculations.StdDev
	assert.EqualFloat64(t, expectedStdDevSTOCK2, stdDev["STOCK2"])
	assert.EqualFloat64(t, expectedStdDevSTOCK1, stdDev["STOCK1"])

	maxDrawdown := returnsCalculations.Drawdown
	maxDrawdownSTOCK2 := maxDrawdown["STOCK2"]
	assert.EqualFloat64(t, expectedMaxDrawdownSTOCK2, maxDrawdownSTOCK2.MaxDrawdown)
	drawdownRangeSTOCK2 := maxDrawdownSTOCK2.DrawdownRange
	assert.EqualStrings(t, expectedMaxDrawdownStartSTOCK2, drawdownRangeSTOCK2.StartDrawdown)
	assert.EqualStrings(t, expectedMaxDrawdownEndSTOCK2, drawdownRangeSTOCK2.EndDrawdown)
	maxDrawdownSTOCK1 := maxDrawdown["STOCK1"]
	assert.EqualFloat64(t, expectedMaxDrawdownSTOCK1, maxDrawdownSTOCK1.MaxDrawdown)
	drawdownRangeSTOCK1 := maxDrawdownSTOCK1.DrawdownRange
	assert.EqualStrings(t, expectedMaxDrawdownStartSTOCK1, drawdownRangeSTOCK1.StartDrawdown)
	assert.EqualStrings(t, expectedMaxDrawdownEndSTOCK1, drawdownRangeSTOCK1.EndDrawdown)

	//histogram := returnsCalculations.Histogram
	//histogramSTOCK2 := histogram["STOCK2"]
	//assert.EqualIntArray(t, expectedHistogramBinCountSTOCK2, histogramSTOCK2.BinCount)
	//assert.EqualFloat64Array(t, expectedHistogramBinEdgesSTOCK2, histogramSTOCK2.BinEdges)
	//histogramSTOCK1 := histogram["STOCK1"]
	//assert.EqualIntArray(t, expectedHistogramBinCountSTOCK1, histogramSTOCK1.BinCount)
	//assert.EqualFloat64Array(t, expectedHistogramBinEdgesSTOCK1, histogramSTOCK1.BinEdges)

	autocorrelation := returnsCalculations.Autocorrelation
	assert.EqualFloat64(t, expectedAutocorrelationSTOCK2, autocorrelation["STOCK2"])
	assert.EqualFloat64(t, expectedAutocorrelationSTOCK1, autocorrelation["STOCK1"])

	correlation := returnsCalculations.Correlation
	index := correlation.Index
	assert.EqualStrings(t, "STOCK1", index[0])
	assert.EqualStrings(t, "STOCK2", index[1])
	//correlationMatrix := correlation.Correlation
	//correlationMatrixSTOCK1 := correlationMatrix[0]
	//assert.EqualFloat64Matrix(t, expectedCorrelationSTOCK1STOCK2, correlationMatrixSTOCK1)
	//correlationMatrixSTOCK2 := correlationMatrix[1]
	//assert.EqualFloat64Matric(t, expectedCorrelationSTOCK2STOCK1, correlationMatrixSTOCK2)
}
