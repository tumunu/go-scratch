package main

import (
	"fmt"
	"flag"
	"math"
)

// -------
// globals
// the length (metres) of timber available (4.8, 5.2, 6 are common lengths)
const stockLength float64 = 4800

// NZS 3604 - 400mm loaded, 600mm unloaded
const loadedSpacing int = 400 
const unloadedSpacing int = 600

// ------
// helpers
func convertToMilimetres(length float64) int {
	var mm = int(length * 1000)
	return mm
}
func roundToNearest(num float64) int {
	var round = math.Round(num)
	var nearest = int(round)
	return nearest
}
func roundToLowest(num float64)  int {
	var round = math.Trunc(num)
	var lowest = int(round)
	return lowest
}
func calculatePercentage(p int, sum int) float64 {
	var percent = ((float64(sum) * float64(p)) / float64(100))
	return percent
}
func isWholeNumber(num float64) bool {
	if num == float64(int64(num)) {
		return true
	}
	return false
}

// -------
// returns the number of cuts we can get from a single piece of timber stock
func calculateStudLengthsFromStock(length float64) int {
	var numberFromStock = stockLength / length
	var amount = roundToLowest(numberFromStock)
	return amount
}

// returns the number of timber stock required
func calculatePlateLengthsFromStock(length float64) int {
	var numberFromStock = length * 2 / stockLength
	var fullLengths = roundToLowest(numberFromStock)
	var amount = fullLengths + 1 // add an extra length of timber stock.
	return amount
}

// returns a basic number of verticle studs (excludes wall junctions and openings)
// to calulate the number of studs: 
// - we need the wall length
// - whether it is a loaded or unloaded wall
func calculateBasicNumberOfStuds(length float64, isLoaded bool) int {
	var spacing = unloadedSpacing
	if isLoaded {
		spacing = loadedSpacing
	}

	var wallLength = length
	var studs = wallLength / float64(spacing)
	var numberOfStuds = roundToNearest(studs)

	return numberOfStuds
}

// returns the total number of timber required for the build based on the timber stock length
// to calculate the number of lengths of timber:
// - we need to know the height of the wall
// - we need to know the length of the wall
// - we need to know whether the wall is loaded or unloaded
// - we need to add an offcut ammount. 
func calculateNumberOfStockRequired(height float64, length float64, isLoaded bool) int {
	// TODO: allow for both milimetres and metres
	var adjustedHeight = height
	var adjustedLength = length

	var numberOfStuds = calculateBasicNumberOfStuds(adjustedLength, isLoaded)
	fmt.Println("# studs in wall:", numberOfStuds)

	// we can calculate the top, and bottom plates, with just the leanth of the wall
	var requiredPlateStock = calculatePlateLengthsFromStock(adjustedLength)
	fmt.Println("# timber stock for top/bottom plates:", requiredPlateStock)

	// we can calculate the rest of the stock by multiplying the total number of studs by the 
	// total number of studs that can be cut from a single stock length
	var numberOfStudsFromStock = calculateStudLengthsFromStock(adjustedHeight)
	var requiredStudStock = numberOfStuds / numberOfStudsFromStock
	fmt.Println("# timber stock for studs:", requiredStudStock)

	// the basic offcut amount (known as wastage) is usually the rough amount of 10%
	var subTotal = requiredPlateStock + requiredStudStock 
	var offcutAmount = calculatePercentage(10, subTotal)

	fmt.Println(" ")
	fmt.Println("wastage:", offcutAmount)

	var adjustedTotal = float64(subTotal) + offcutAmount
	var stock = roundToNearest(adjustedTotal)

	return stock
}

// ------
func main() {
	var heightPtr = flag.Float64("h", 2.4, "height")
	var lengthPtr = flag.Float64("l", 5.0, "length")
	var isLoadedPtr = flag.Bool("w", false, "loaded")

	flag.Parse()

	var length = *lengthPtr
	var height = *heightPtr
	var isLoaded = *isLoadedPtr

	fmt.Println("Timber stock length (mm):", stockLength)
	fmt.Println(" ")
	fmt.Println("Total # timber stock required:", calculateNumberOfStockRequired(height, length, isLoaded))
	fmt.Println(" ")
}