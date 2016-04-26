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
	"github.com/codegangsta/cli"
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
	c             *cli.Context
}

func NewTable(c *cli.Context, headers []string) Table {
	return &PrintableTable{
		headers:  headers,
		maxSizes: make([]int, len(headers)),
		c:        c,
	}
}

func (t *PrintableTable) Add(row ...string) {
	t.rows = append(t.rows, row)
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
	fmt.Fprintln(t.c.App.Writer, output)
}

func (t *PrintableTable) printRow(row []string) {
	output := ""
	for columnIndex, value := range row {
		if columnIndex == 0 {
			value = value
		}

		output = output + t.cellValue(columnIndex, value)
	}
	fmt.Fprintf(t.c.App.Writer, "%s\n", output)
}

func (t *PrintableTable) cellValue(col int, value string) string {
	padding := ""
	if col < len(t.headers)-1 {
		padding = strings.Repeat(" ", t.maxSizes[col]-utf8.RuneCountInString(value))
	}
	return fmt.Sprintf("%s%s   ", value, padding)
}
