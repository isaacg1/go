package main

func pow10(n int) int {
	prod := 1
	for i := 0; i < n; i++ {
		prod *= 10
	}
	return prod
}

func champ(d_pos int) int {
	pos := d_pos - 1
	block := 9
	d_len := 1
	for d_len*block < pos {
		pos -= d_len * block
		d_len++
		block *= 10
	}
	indexed_num := block/10 + pos/d_len + 1
	cutoff := pow10(pos%d_len + 1)
	for indexed_num >= cutoff {
		indexed_num /= 10
	}
	return indexed_num % 10
}

func main() {
	prod := 1
	for i := 0; i <= 6; i++ {
		prod *= champ(pow10(i))
	}
	println(prod)
}
