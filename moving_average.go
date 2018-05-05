package moving_average

import (
	"math/big"
	"errors"
)

var (
	InvalidPeriodErr = errors.New("Provided period is not strictly positive")
	NotEnoughDataPointsErr = errors.New("Not enough data points provided")
)

// Simple implements the Simple Moving Average algorithm for a period of n using provided data.
// Expects data points to be from latest to oldest.
// Returns moving average from latest to oldest.
func Simple(data []*big.Float, n uint) ([]*big.Float, error) {
	if err := validateParameters(uint(len(data)), n); err != nil {
		return nil, err
	}

	var result []*big.Float
	var previousAverage *big.Float

	for i := uint(0); i < n; i++ {
		sumOfNLatestPoints := new(big.Float)

		// Check if we can re-use data from previous data point
		if previousAverage == nil {
			// We have to calculate it from the first n elements of the data
			// We know we always have n elements since we validated the parameters
			for j, value := range data {
				if uint(j) == n {
					break
				}

				// Add the value to the sum of the n latest points
				sumOfNLatestPoints = new(big.Float).Add(sumOfNLatestPoints, value)
			}
		} else {
			// Check if we have enough remaining points to add another entry
			if n + i - 1 > uint(len(data)) {
				break
			}

			// Subtract the latest point used in the previous average
			// Add an older point instead of the removed point
			sumOfNLatestPoints = new(big.Float).Sub(previousAverage, data[i - 1])
			sumOfNLatestPoints = new(big.Float).Add(previousAverage, data[n + i])
		}

		// Calculate the average of the points
		currentAverage := new(big.Float).Quo(sumOfNLatestPoints, big.NewFloat(float64(n)))
		// Add it to the results
		result = append(result, currentAverage)

		previousAverage = currentAverage
	}

	return result, nil
}

// SimpleRev implements the Simple Moving Average algorithm for a period of n using provided data.
// Expects data points to be from oldest to latest.
// Returns moving average from oldest to latest.
func SimpleRev(data []*big.Float, n uint) ([]*big.Float, error) {
	if uint(len(data)) < n {
		return nil, NotEnoughDataPointsErr
	}
}

func validateParameters(dataPointsCount, n uint) error {
	if n < 1 {
		return InvalidPeriodErr
	}

	if dataPointsCount < n {
		return NotEnoughDataPointsErr
	}

	return nil
}

func Cumulative() {

}

func Weighted() {

}