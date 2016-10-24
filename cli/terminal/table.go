/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package terminal

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Table interface {
	Add(row ...string)
	Print()
}

type PrintableTable struct {
	headers       []string
	headerPrinted bool
	maxSizes      []int
	rows          [][]string
}

func NewTable(headers []string) Table {
	return &PrintableTable{
		headers:  headers,
		maxSizes: make([]int, len(headers)),
	}
}

func (t *PrintableTable) Add(rows ...string) {
	singleLines := make([]string, len(rows))
	for i, r := range rows {
		single := strings.Replace(r, "\n", " ", -1)
		singleLines[i] = strings.TrimSpace(single)
	}
	t.rows = append(t.rows, singleLines)
}

func (t *PrintableTable) Print() {
	for _, row := range append(t.rows, t.headers) {
		t.calculateMaxSize(row)
	}

	if t.headerPrinted == false {
		t.printHeader()
		t.headerPrinted = true
	}

	for _, line := range t.rows {
		t.printRow(line)
	}

	t.rows = [][]string{}
}

func (t *PrintableTable) calculateMaxSize(row []string) {
	for index, value := range row {
		cellLength := utf8.RuneCountInString(value)
		if t.maxSizes[index] < cellLength {
			t.maxSizes[index] = cellLength
		}
	}
}

func (t *PrintableTable) printHeader() {
	output := ""
	for col, value := range t.headers {
		output = output + t.cellValue(col, value)
	}
	fmt.Println(output)
}

func (t *PrintableTable) printRow(row []string) {
	output := ""
	for columnIndex, value := range row {
		output = output + t.cellValue(columnIndex, value)
	}
	fmt.Printf("%s\n", output)
}

func (t *PrintableTable) cellValue(col int, value string) string {
	padding := ""
	delim := ""
	if col < len(t.headers)-1 {
		delim = "| "
		padding = strings.Repeat(" ", t.maxSizes[col]-utf8.RuneCountInString(value))
	}
	return fmt.Sprintf("%s%s   " + delim, value, padding)
}
