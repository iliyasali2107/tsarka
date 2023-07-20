package service

import (
	"fmt"
	"strconv"
	"time"
)

type IINService interface {
	Check(iin string) (string, bool)
}

type iinService struct{}

func NewIINService() IINService {
	return &iinService{}
}

var monthMap = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func (is *iinService) Check(iin string) (string, bool) {
	if len(iin) != 12 {
		return "", false
	}

	iinInt := []int{}

	for _, s := range iin {
		if s > '9' || s < '0' {
			return "", false
		}

		iinInt = append(iinInt, int(s)-48)
	}

	sexCentury, err := strconv.Atoi(string(iin[6]))
	if err != nil && sexCentury < 6 && sexCentury != 0 {
		return "", false
	}

	if sexCentury > 6 || sexCentury == 0 {
		return "", false
	}

	currentYear := time.Now().Year()

	year, _ := strconv.Atoi(iin[:2])

	sex, century := sextury(sexCentury)
	year = century*100 + year
	if year > currentYear {
		return "", false
	}

	month, _ := strconv.Atoi(iin[2:4])
	if month < 1 || month > 12 {
		return "", false
	}

	var maxDays int
	if month == 2 && isLeapYear(year) {
		maxDays = 29
	} else {
		maxDays = monthMap[month]
	}
	day, _ := strconv.Atoi(iin[4:6])
	if day < 1 || day > 31 {
		return "", false
	}

	if day > maxDays {
		return "", false
	}

	last, _ := strconv.Atoi(string(iin[11]))

	vesRazryada := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	vesRazryadaDva := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}
	sum := 0

	for i := 0; i < 11; i++ {
		sum += (vesRazryada[i] * iinInt[i])
	}

	if sum%11 == 10 {
		sum = 0
		for i := 1; i < 12; i++ {
			sum += (vesRazryadaDva[i] * iinInt[i])
		}
	}

	if sum%11 == 10 || sum%11 != last {
		return "", false
	}

	return fmt.Sprintf("it is a %s, born in %d.%d.%d", sex, day, month, year), true
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

func sextury(n int) (sex string, century int) {
	if n%2 == 0 {
		if n == 2 {
			century = 18
		} else if n == 4 {
			century = 19
		} else {
			century = 20
		}
		sex = "female"
	} else {
		if n == 1 {
			century = 18
		} else if n == 3 {
			century = 19
		} else {
			century = 20
		}
		sex = "male"
	}
	return
}
