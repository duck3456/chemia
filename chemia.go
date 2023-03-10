package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"sort"
	"reflect"
)

var (
	num1 int = 0
	num2 int = 1
	num3 int = 1
	num4 int = 1

	array1 []string
	array1FirstHalf []string
	array1SecondHalf []string
	array1Final []string
	array2 []string
	array3 []string
	array4 []string
	array5 []string
	
	result bool = false

	finalResult string

	tries int = 0

	input string

	isFirst bool
	canReturn bool
)

const (
	MAX_NUM = 6
)

func main() {
	fmt.Println()
	fmt.Println("Program ten uzgadnia reakcje wymiany i podwójnej wymiany; reakcja musi zawierać 2 substraty i 2 produkty")
	fmt.Println("Reakcja musi być zapisana w specjalnym formacie, bez spacji, a zamiast strzałki jest _")
	fmt.Println("Np. NaOH+N2O5_NaNO3+H2O")
	fmt.Println("Program nie działa z nawiasami, więc np. zamiast (OH)2 trzeba zapisać O2H2, a zamiast Ca(NO3)2 - CaN2O6")
	fmt.Println("Program nie uzgadnia reakcji, gdzie współczynnik stechiometryczny może być większy niż 6")
	fmt.Println("Więc jeśli czekasz ponad kilka sekund i program nie uzgodnił reakcji, to albo reakcja jest zapisana źle, albo wymaga uzgodnienia z większymi współczynnikami stechiometrycznymi")
	fmt.Println()

	fmt.Print("Wpisz reakcję: ")
	fmt.Scan(&input)

	/*fmt.Println()
	fmt.Println(array1)
	fmt.Println(array1FirstHalf)
	fmt.Println(array1SecondHalf)
	fmt.Println(array1Final)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(result)*/

	for !result {
		tries++
		num1++

		if num1 > MAX_NUM {
			num1 = 1
			num2++
		}

		if num2 > MAX_NUM {
			num2 = 1
			num3++
		}

		if num3 > MAX_NUM {
			num3 = 1
			num4++
		}

		if num4 > MAX_NUM {
			num4 = 1
		}

		
		array1 = isolateElements(input)

		array1FirstHalf = makeFirstSmallerArray(array1)
		array1SecondHalf = makeSecondSmallerArray(array1)

		array1Final = multiplyElementsRandomly(array1FirstHalf, array1SecondHalf, []int{num1, num2, num3, num4})

		array2 = isolateElementsEvenMore(array1Final)
		array3 = countElements(array2)
		array4 = makeFirstSmallerArray(array3)
		array5 = makeSecondSmallerArray(array3)

		result = EqualIgnoringOrder(array4, array5)

		if result {
			for i := 2; i <= MAX_NUM; i++ {
				if (num1%i == 0 && num2%i == 0 && num3%i == 0 && num4%i == 0) {
					num1 = num1/i
					num2 = num2/i
					num3 = num3/i
					num4 = num4/i
				}
			}

			finalResult = fmt.Sprint(num1) + array1[0] + "+" + fmt.Sprint(num2) + array1[1] + "_" + fmt.Sprint(num3) + array1[3] + "+" + fmt.Sprint(num4) + array1[4]
			finalResult = strings.ReplaceAll(finalResult, "+1", "+")
			finalResult = strings.ReplaceAll(finalResult, "_1", "_")
			
			if string(finalResult[0]) == "1" {
				finalResult = finalResult[1:]
			}

			fmt.Println("Reakcja uzgodniona!", finalResult)
			fmt.Println()
			fmt.Println("Próby uzgodnienia podjęte przez program:", tries)
			fmt.Scan(&input)
			break
		}
	}
}

func IsUpper(s string) bool {
    for _, r := range s {
        if unicode.IsUpper(r) && unicode.IsLetter(r) {
            return true
        }
    }
    return false
}

func IsLetter(s string) bool {
    for _, r := range s {
        if unicode.IsLetter(r) {
            return true
        }
    }
    return false
}

func isolateElements(expression string) []string {
	
	elements := []string{}

	//1. isolate the elements like this: [Na +H2SO4 , _NaSO4 +H2]
	lastStopIndex := 0

	for i := 0; i < len(expression); i++ {
		
		if (string(expression[i]) == "+" || string(expression[i]) == "_" || i == len(expression)-1) {
			elements = append(elements, expression[lastStopIndex:i+1])
			lastStopIndex = i

			if string(expression[i]) == "_" {
				elements = append(elements, ",")
			}
		}

	}

	//elements[len(elements)-1] += string(expression[len(expression)-1]) //add the last character, which got lost in the code above

	//2. remove "+" and "_" from the array: [Na H2SO4 , NaSO4 H2]
	for i := 0; i < len(elements); i++ {
		elements[i] = strings.ReplaceAll(elements[i], "+", "")
		elements[i] = strings.ReplaceAll(elements[i], "_", "")
	}

	return elements
}

func multiplyElementsRandomly(firstHalf []string, secondHalf []string, numbers []int) []string {
	elements := []string{}

	for i := 0; i < len(firstHalf); i++ {
		
		num := numbers[i]

		for j := 0; j < num; j++ {
			elements = append(elements, firstHalf[i])
		}
	}

	elements = append(elements, ",")

	for i := 0; i < len(secondHalf); i++ {
		
		num := numbers[i+2]

		for j := 0; j < num; j++ {
			elements = append(elements, secondHalf[i])
		}
	}

	return elements
}

func isolateElementsEvenMore(elements []string) []string {
	
	//1. isolate the elements even more like this: [Na H2 S O4 , Na S O4 H2]
	evenMoreIsolatedElements := []string{}

	elementsString := strings.Join(elements, "")

	lastStopIndex := 0
	isFirst = true

	for i := 0; i < len(elementsString); i++ {
		
		if (IsUpper(string(elementsString[i])) || string(elementsString[i]) == ",") {

			if (!isFirst) {
				evenMoreIsolatedElements = append(evenMoreIsolatedElements, elementsString[lastStopIndex:i])
				lastStopIndex = i
			}
			isFirst = false
		}
		if i == len(elementsString)-1 {
			evenMoreIsolatedElements = append(evenMoreIsolatedElements, elementsString[lastStopIndex:i+1])
		}
	}

	return evenMoreIsolatedElements
}

func countElements(evenMoreIsolatedElements []string) []string {
	
	//1. make an array like this: [Na H H S O O O O , Na S O O O O H H]
	countedElements := []string{}

	for i := 0; i < len(evenMoreIsolatedElements); i++ {

		if (IsLetter(string(evenMoreIsolatedElements[i][len(evenMoreIsolatedElements[i])-1:])) == false && string(evenMoreIsolatedElements[i][len(evenMoreIsolatedElements[i])-1:]) != ",") {
			//fmt.Println(string(evenMoreIsolatedElements[i][len(evenMoreIsolatedElements[i])-1:]))

			num, _ := strconv.Atoi(string(evenMoreIsolatedElements[i][len(evenMoreIsolatedElements[i])-1:]))
			numTwoDigit, _ := strconv.Atoi(string(evenMoreIsolatedElements[i][len(evenMoreIsolatedElements[i])-2:]))

			if numTwoDigit == 0 {
				for j := 0; j < num; j++ {
					countedElements = append(countedElements, evenMoreIsolatedElements[i])
				}
			} else {
				for j := 0; j < numTwoDigit; j++ {
					countedElements = append(countedElements, evenMoreIsolatedElements[i])
				}
			}
		} else {
			countedElements = append(countedElements, evenMoreIsolatedElements[i])
		}
	}

	for i := 0; i < len(countedElements); i++ {
		for j := 0; j < 10; j++ {
			countedElements[i] = strings.ReplaceAll(countedElements[i], fmt.Sprint(j), "")
		}
	}

	return countedElements
}

func makeFirstSmallerArray(countedElements []string) []string {
	
	//1. make an array like this: [Na H H S O O O O]
	firstSmallerArray := []string{}

	for i := 0; i < len(countedElements); i++ {
		if string(countedElements[i]) == "," {
			return firstSmallerArray
		} else {
			firstSmallerArray = append(firstSmallerArray, countedElements[i])
		}
	}

	return nil
}

func makeSecondSmallerArray(countedElements []string) []string {
	
	secondSmallerArray := []string{}
	canReturn = false

	for i := 0; i < len(countedElements); i++ {
		if canReturn {
			secondSmallerArray = append(secondSmallerArray, countedElements[i])
		}

		if string(countedElements[i]) == "," {
			canReturn = true
		}
	}

	return secondSmallerArray
}

//stackoverflow.com/questions/73094902/how-to-compare-two-arrays-is-same-with-any-orders-value-in-golang
func EqualIgnoringOrder(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    sort.Strings(a)
    sort.Strings(b)

    return reflect.DeepEqual(a, b)
}