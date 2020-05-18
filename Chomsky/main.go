package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type mapsWithDuplicate struct {
	symbols string
	values  string
}

func main() {

	var myarray []mapsWithDuplicate
	createMyMap(readLines("var13.txt"), &myarray)
	addNewStartingSymbol(&myarray)
	removeEpsilonProduction(&myarray)
	makeUnitSubstitution(&myarray)
	removeNonGeneratingSymbols(&myarray)
	removeUnreachebleSymbols1(&myarray)
	getArrayOfNewSymbols(duplicateCount(getProductionInWrongForm(&myarray)), &myarray)
	fmt.Print(myarray)

}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func createMyMap(lines []string, myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '-' {
				*myarray = append(*myarray, mapsWithDuplicate{lines[i][:j], lines[i][j+1:]})
			}
		}
	}
	return myarray
}

//remove epsilon productions
func removeEpsilonProduction(myArray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	var arrayToRemove []mapsWithDuplicate
	for i := 0; i < len(*myArray); i++ {
		if (*myArray)[i].values == "ε" {
			arrayToRemove = append(arrayToRemove, (*myArray)[i])
			makeSubstitution(myArray, findEpsilonProductionsSymbols(myArray))
			deleteMultipleElements1(myArray, arrayToRemove)
		}
	}
	return myArray
}

func makeSubstitution(myArray *[]mapsWithDuplicate, stringsArray []string) *[]mapsWithDuplicate {
	for i := 0; i < len(*myArray); i++ {
		for j := 0; j < len(stringsArray); j++ {
			if stringContains((*myArray)[i].values, stringsArray[j]) {
				*myArray = append(*myArray, mapsWithDuplicate{(*myArray)[i].symbols, myTrimFunc((*myArray)[i].values, stringsArray[j])})
			}
		}
	}
	return myArray
}

func findEpsilonProductionsSymbols(myArray *[]mapsWithDuplicate) []string {
	var epsilonArray []string
	for i := 0; i < len(*myArray); i++ {
		if (*myArray)[i].values == "ε" {
			epsilonArray = append(epsilonArray, (*myArray)[i].symbols)
		}
	}
	return epsilonArray

}

//unit productions
func makeUnitSubstitution(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	var arrayToDelete []mapsWithDuplicate
	for i := len(*myarray) - 1; i > 0; i-- {
		if isUnitProduction((*myarray)[i]) {
			arrayToDelete = append(arrayToDelete, mapsWithDuplicate{(*myarray)[i].symbols, (*myarray)[i].values})
			for j := 0; j < len(getProductionBySymbol(*myarray, (*myarray)[i].values)); j++ {
				*myarray = append(*myarray, mapsWithDuplicate{(*myarray)[i].symbols, getProductionBySymbol(*myarray, (*myarray)[i].values)[j]})
			}
			deleteMultipleElements1(myarray, arrayToDelete)
			continue
		}
	}
	return myarray
}

func isUnitProduction(myarray mapsWithDuplicate) bool {
	if (len(myarray.symbols) == 1 && len(myarray.values) == 1) &&
		(isNonTerminal(myarray.symbols[0]) && isNonTerminal(myarray.values[0])) {
		return true
	}

	return false
}

func getProductionBySymbol(myarray []mapsWithDuplicate, symbol string) []string {
	var arrayProduction []string
	for i := 0; i < len(myarray); i++ {
		if myarray[i].symbols == symbol {
			arrayProduction = append(arrayProduction, myarray[i].values)
		}
	}
	return arrayProduction
}

func deleteMultipleElements1(baseArray *[]mapsWithDuplicate, arrayToDelete []mapsWithDuplicate) *[]mapsWithDuplicate {
	for i := 0; i < len(*baseArray); i++ {
		url := (*baseArray)[i]
		for _, rem := range arrayToDelete {
			if url == rem {
				*baseArray = append((*baseArray)[:i], (*baseArray)[i+1:]...)
				i--
				break
			}
		}
	}
	return baseArray

}

func myTrimFunc(word string, charToTrim string) string {

	myRunes := []rune(word)
	for i := 0; i < len(myRunes); i++ {
		if string(myRunes[i]) == charToTrim && len(string(myRunes)) > 1 {
			myRunes := append(myRunes[:i], myRunes[i+1:]...)
			return string(myRunes)
		} else if word == charToTrim {
			myRunes[i] = 949
			return string(myRunes)
		}
	}
	return "WORD DOES NOT CONTAIN CHARACTER"
}

func stringContains(word string, char string) bool {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == char {
			return true
		}
	}
	return false
}

func isNonTerminal(char uint8) bool {
	if char >= 65 && char <= 90 {
		return true
	}
	return false
}

func findStartingSymbol(myMap *[]mapsWithDuplicate) string {
	return (*myMap)[0].symbols
}

func addNewStartingSymbol(myMap *[]mapsWithDuplicate) {
	for i := 0; i < len(*myMap); i++ {
		if stringContains((*myMap)[i].values, findStartingSymbol(myMap)) {
			*myMap = append(*myMap, mapsWithDuplicate{"W", findStartingSymbol(myMap)})
			break
		}
	}
}

func hasTerminal(stringArray []string) bool {
	for i := 0; i < len(stringArray); i++ {
		if isTerminalString(stringArray[i]) {
			return true
		}
	}
	return false
}

func isTerminalString(string2 string) bool {
	if len(string2) == 1 && isTerminalChar(string2[0]) {
		return true
	} else if len(string2) == 1 && !isTerminalChar(string2[0]) {
		return false
	} else if len(string2) > 1 {
		for j := 0; j < len(string2); j++ {
			if isTerminalChar(string2[j]) && j == len(string2)-1 {
				continue
			} else if !isTerminalChar(string2[j]) && j == len(string2)-1 {
				return false
			}
		}
	}
	return true
}

func isTerminalChar(char uint8) bool {
	if char >= 97 && char <= 122 {
		return true
	}
	return false
}

func RandomStringGenerator(length int, letterRunes []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func createAccessibleRunes(arrayToDelete map[string]bool) []rune {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 0; i < len(letterRunes); i++ {
		url := letterRunes[i]
		for rem := range arrayToDelete {
			if string(url) == rem {
				letterRunes = append(letterRunes[:i], letterRunes[i+1:]...)
				i--
				break
			}
		}
	}
	return letterRunes
}

// non generating symbols
func removeNonGeneratingSymbols(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	var arrayOfNonGeneratingSymbols []mapsWithDuplicate
	for i := 0; i < len(*myarray); i++ {
		if hasTerminal(getProductionBySymbol( *myarray, (*myarray)[i].symbols)) {
			continue
		}
		arrayOfNonGeneratingSymbols = append(arrayOfNonGeneratingSymbols, mapsWithDuplicate{(*myarray)[i].symbols, (*myarray)[i].values})
	}
	return deleteMultipleElements1(myarray, arrayOfNonGeneratingSymbols)
}

//unreacheble symbols
func removeUnreachebleSymbols1(myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	var arrayToDelete []mapsWithDuplicate
	for k := range compareProductions(getMyArrayProductions(myarray), getStaringSymbolsProduction(myarray)) {
		if compareProductions(getMyArrayProductions(myarray), getStaringSymbolsProduction(myarray))[k] == false {
			for _, n := range getProductionBySymbol(*myarray, k) {
				arrayToDelete = append(arrayToDelete, mapsWithDuplicate{k, n})
			}
		}
	}
	return deleteMultipleElements1(myarray, arrayToDelete)
}

func compareProductions(myMap map[string]bool, startingSymbolMap map[string]bool) map[string]bool {
	for k := range myMap {
		for i := range startingSymbolMap {
			if k == i {
				myMap[k] = true
			}
		}
	}
	return myMap
}

func getMyArrayProductions(myarray *[]mapsWithDuplicate) map[string]bool {
	myMap := make(map[string]bool)
	for i := 0; i < len(*myarray); i++ {
		myMap[(*myarray)[i].symbols] = false
	}
	return myMap
}

func getStaringSymbolsProduction(myarray *[]mapsWithDuplicate) map[string]bool{
	myMap := make(map[string]bool)
	for i:=0;i<len(getProductionBySymbol(*myarray,findStartingSymbol(myarray)));i++ {
		for _,v := range getProductionBySymbol(*myarray,findStartingSymbol(myarray))[i]{
			if isNonTerminal(uint8(v)) {
				myMap[findStartingSymbol(myarray)] = true
				myMap[string(uint8(v))] = true
			}
		}
	}
	return myMap
}

func getArrayOfNewSymbols(mymap map[string]int, myarray *[]mapsWithDuplicate) *[]mapsWithDuplicate {
	myFirstMap := make(map[string]string)
	myMap := make(map[string]string)
	for k := range mymap {
		myFirstMap[k[0:len(k)-1]] = RandomStringGenerator(1, createAccessibleRunes(getStaringSymbolsProduction(myarray)))
		myMap[(*myarray)[getTerminalByProduction(*myarray, (*myarray)[getTerminalByProduction(*myarray, k)].values)].values] =
			(*myarray)[getTerminalByProduction(*myarray, (*myarray)[getTerminalByProduction(*myarray, k)].values)].values[0 : len(k)-1]
		*myarray = append(*myarray, mapsWithDuplicate{RandomStringGenerator(1, createAccessibleRunes(getStaringSymbolsProduction(myarray))), k[0 : len(k)-1]})
	}

	for i := 0; i < len(*myarray); i++ {
		for k, v := range createArrayOfNewSymbols(myMap, myFirstMap) {
			if (*myarray)[i].values == k {
				(*myarray)[i].values = v
			}
		}

	}

	return myarray
}

func createArrayOfNewSymbols(myMap map[string]string, myFirstMap map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range myMap {
		for k, v := range myFirstMap {
			if value == k {
				newMap[key] = v + key[len(k):]
			}
		}
	}
	return newMap
}

func getTerminalByProduction(myArray []mapsWithDuplicate, input string) int {
	for i := 0; i < len(myArray); i++ {
		if myArray[i].values == input {
			return i
		}
	}
	return 0
}

func duplicateCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}

func getProductionInWrongForm(myarray *[]mapsWithDuplicate) []string {
	var wrongFormArray []string
	for k := range getMyArrayProductions(myarray) {
		for _, v := range getProductionBySymbol(*myarray, k) {
			if (len(v) == 2 && isTerminalChar(v[0]) && isNonTerminal(v[1])) ||
				(len(v) == 2 && isNonTerminal(v[0]) && isTerminalChar(v[1])) || len(v) > 2 {
				wrongFormArray = append(wrongFormArray, v)
			}
		}
	}
	return wrongFormArray
}
