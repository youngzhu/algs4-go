package pq_test

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/sorting/pq"
	"github.com/youngzhu/algs4-go/util"
)

// A priority queue client that reads transations, and prints out the M largest ones
const M = 5

func ExampleMinPQ_topM() {
	minPQ := pq.NewMinPQN(M)
	topM(minPQ, "amount")

	// Output:
	// Thompson    2/27/2000  4747.08
	// vonNeumann  2/12/1994  4732.35
	// vonNeumann  1/11/1999  4409.74
	// Hoare       8/18/1992  4381.21
	// vonNeumann  3/26/2002  4121.85
}

func ExampleMaxPQ_topM() {
	maxPQ := pq.NewMaxPQN(M)
	topM(maxPQ, "date")

	// Output:
	// Turing      6/17/1990   644.08
	// Turing      2/11/1991  2156.86
	// Hoare       8/18/1992  4381.21
	// Hoare       5/10/1993  3229.27
	// Turing     10/12/1993  3532.36
}

func topM(priorityQueue pq.PriorityQueue, order string) {
	in := util.NewInReadLines("testdata/tinyBatch.txt")

	for !in.IsEmpty() {
		line := in.ReadString()
		var trans pq.Item
		if order == "date" {
			trans = newDateOrderTrans(line)
		} else {
			trans = newAmountOrderTrans(line)
		}

		priorityQueue.Insert(trans)

		// remove minimum/maximum if M+1 entries on the PQ
		if priorityQueue.Size() > M {
			priorityQueue.Delete()
		}
		// top M entries are on the PQ
	}

	// print entries on PQ in reverse order
	stack := fund.NewStack()
	for !priorityQueue.IsEmpty() {
		stack.Push(priorityQueue.Delete())
	}
	for _, v := range stack.Iterator() {
		fmt.Println(v)
	}
}

type (
	DateOrderTrans   Transaction // order by date
	AmountOrderTrans Transaction // order by amount
)

func newDateOrderTrans(s string) DateOrderTrans {
	who, when, amount := parseTransaction(s)
	return DateOrderTrans{who, when, amount}
}

func newAmountOrderTrans(s string) AmountOrderTrans {
	who, when, amount := parseTransaction(s)
	return AmountOrderTrans{who, when, amount}
}

func (t DateOrderTrans) CompareTo(x pq.Item) int {
	tt := x.(DateOrderTrans)
	return t.when.CompareTo(tt.when)
}

func (t AmountOrderTrans) CompareTo(x pq.Item) int {
	tt := x.(AmountOrderTrans)
	if t.amount < tt.amount {
		return -1
	} else if t.amount > tt.amount {
		return 1
	} else {
		return 0
	}
}

func (t DateOrderTrans) String() string {
	return fmt.Sprintf("%-10s %10s %8.2f", t.who, t.when, t.amount)
}

func (t AmountOrderTrans) String() string {
	return fmt.Sprintf("%-10s %10s %8.2f", t.who, t.when, t.amount)
}

// Data type for commercial transactions
type Transaction struct {
	who    string // customer
	when   Date   // date
	amount float64
}

func parseTransaction(s string) (string, Date, float64) {
	datas := strings.Fields(s)
	who := datas[0]
	when := newDate(datas[1])

	amount, err := strconv.ParseFloat(datas[2], 64)
	if err != nil {
		panic("invalid amount: " + datas[2])
	}

	if math.IsNaN(amount) || math.IsInf(amount, 0) {
		panic("Amount cannot be NaN or infinite")
	}

	return who, when, amount
}

var days = [...]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// An immutable data type for dates
type Date struct {
	month int // between 1 and 12
	day   int // between 1 and days[month]
	year  int
}

func newDate(s string) Date {
	fields := strings.Split(s, "/")
	if len(fields) != 3 {
		panic("Invalid date: " + s)
	}
	month, _ := strconv.Atoi(fields[0])
	day, _ := strconv.Atoi(fields[1])
	year, _ := strconv.Atoi(fields[2])

	if !isValidDate(month, day, year) {
		panic("Invalid date: " + s)
	}

	return Date{month, day, year}
}

// is the given date valid?
func isValidDate(m, d, y int) bool {
	if m < 1 || m > 12 {
		return false
	}
	if d < 1 || d > days[m] {
		return false
	}
	if m == 2 && d == 29 && !isLeapYear(y) {
		return false
	}
	return true
}

// is y a leap year?
func isLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%100 == 0 {
		return false
	}
	return y%4 == 0
}

func (d Date) CompareTo(x Date) int {
	if d.year < x.year {
		return -1
	}
	if d.year > x.year {
		return 1
	}
	if d.month < x.month {
		return -1
	}
	if d.month > x.month {
		return 1
	}
	if d.day < x.day {
		return -1
	}
	if d.day > x.day {
		return 1
	}
	return 0
}

func (d Date) String() string {
	return fmt.Sprintf("%d/%d/%d", d.month, d.day, d.year)
}
