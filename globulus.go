package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var header = true

func main() {
	var err error
	flag.Parse()
	fd := os.Stdin

	if len(flag.Args()) > 0 {
		fd, err = os.Open(flag.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(file *os.File) { _ = file.Close() }(fd)

	r := csv.NewReader(fd)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Scan how many times values appear in the columns
	columnCountDistinctValues := make(map[int]map[string]int)
	for nr, record := range records {
		if header && nr == 0 {
			continue
		}
		for i, v := range record {
			if columnCountDistinctValues[i] == nil {
				columnCountDistinctValues[i] = make(map[string]int)
				columnCountDistinctValues[i][v] = 0
			}
			columnCountDistinctValues[i][v] = columnCountDistinctValues[i][v] + 1
		}
	}

	// Scan count data types of the columns
	columnDataTypes := make(map[int]map[string]int)
	for nr, record := range records {
		if header && nr == 0 {
			continue
		}
		for columnNumber, value := range record {
			if columnDataTypes[columnNumber] == nil {
				columnDataTypes[columnNumber] = make(map[string]int)
				for _, typeName := range []string{"empyty", "int", "float", "bool", "other"} {
					columnDataTypes[columnNumber][typeName] = 0
				}
			}
			if value == "" {
				columnDataTypes[columnNumber]["empty"] = columnDataTypes[columnNumber]["empty"] + 1
			} else if _, err := strconv.ParseInt(value, 10, 64); err == nil {
				columnDataTypes[columnNumber]["int"] = columnDataTypes[columnNumber]["int"] + 1
			} else if _, err := strconv.ParseFloat(value, 64); err == nil {
				columnDataTypes[columnNumber]["float"] = columnDataTypes[columnNumber]["float"] + 1
			} else if _, err := strconv.ParseBool(value); err == nil {
				columnDataTypes[columnNumber]["bool"] = columnDataTypes[columnNumber]["bool"] + 1
			} else {
				columnDataTypes[columnNumber]["other"] = columnDataTypes[columnNumber]["other"] + 1
			}
		}
	}

	// Score the columns by the data types contained
	typeScores := map[string]float64{
		"empty": 0,
		"float": 1,
		"int":   2,
		"bool":  3,
		"other": 5,
	}

	columnTypeScore := make([]float64, r.FieldsPerRecord)
	for columnNumber, typeCounts := range columnDataTypes {
		score := 0.0
		for typeName, tc := range typeCounts {
			score += typeScores[typeName] * float64(tc)
		}
		columnTypeScore[columnNumber] = score
	}
	//fmt.Fprintf(os.Stderr,"%v\n", columnDataTypes)
	//fmt.Fprintf(os.Stderr,"%v\n", columnTypeScore)

	// Group columns by count of different values in them
	// Initialise empty matrix
	groupByCountValues := map[int][]int{}
	for i := range groupByCountValues {
		groupByCountValues[i] = []int{}
	}
	// Scan the value counts, collect max ever
	max := 0
	for columnNumber, values := range columnCountDistinctValues {
		if len(values) > max {
			max = len(values)
		}
		groupByCountValues[len(values)] = append(groupByCountValues[len(values)], columnNumber)
	}

	// Compute a new ordering of the columns
	// get a descending ordered list of counts
	counts := make([]int, 0, len(groupByCountValues))
	for k := range groupByCountValues {
		counts = append(counts, k)
	}
	sort.Ints(counts)

	// get full list of columns in the new order
	var newOrder = make([]int, r.FieldsPerRecord)
	var columnNumber = 0
	for _, count := range counts {
		colList := groupByCountValues[count]
		// Tie-breaker - order these columns by their data type scores
		sort.Slice(colList, func(i, j int) bool {
			return columnTypeScore[i] < columnTypeScore[j]
		})

		for _, newColumnNumber := range colList {
			newOrder[columnNumber] = newColumnNumber
		}
	}

	// Now re-order the data
	var reordered [][]string
	for _, record := range records {
		newrow := make([]string, len(record))
		for i := 0; i < r.FieldsPerRecord; i++ {
			newrow[i] = record[newOrder[i]]
		}
		reordered = append(reordered, newrow)
	}

	w := csv.NewWriter(os.Stdout)
	err = w.WriteAll(reordered)
	if err != nil {
		log.Fatal(err)
	}

	var explainTable [][]string
	if header {
		explainTable = append(explainTable, records[0])
	}
	valueCounts := make([]string, r.FieldsPerRecord)
	for columnNumber, values := range columnCountDistinctValues {
		valueCounts[newOrder[columnNumber]] = strconv.Itoa(len(values))
	}
	explainTable = append(explainTable, valueCounts)

	scores := make([]string, r.FieldsPerRecord)
	for columnNumber := 0; columnNumber < r.FieldsPerRecord; columnNumber++ {
		scores[newOrder[columnNumber]] = fmt.Sprintf("%f", columnTypeScore[columnNumber])
	}
	explainTable = append(explainTable, scores)

	we := csv.NewWriter(os.Stdout)
	err = we.WriteAll(explainTable)
	if err != nil {
		log.Fatal(err)
	}



}
