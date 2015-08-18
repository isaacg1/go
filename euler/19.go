package main

func sundays(jan1 int, leap int) (int, int) {
	day := jan1
	sunday_count := 0
	month_lens := []int{31, 28 + leap, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for _, m_len := range month_lens {
		if day%7 == 0 {
			sunday_count += 1
		}
		day += m_len
	}
	return sunday_count, day % 7
}

func main() {
	jan1 := 2
	sunday_total := 0
	for year := 1901; year <= 2000; year++ {
		var leap int
		if year%4 == 0 {
			leap = 1
		} else {
			leap = 0
		}
		var sunday_count int
		sunday_count, jan1 = sundays(jan1, leap)
		sunday_total += sunday_count
	}
	println(sunday_total)
}
