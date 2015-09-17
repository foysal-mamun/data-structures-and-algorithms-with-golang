package maths

import "math"

// Factorial uisng recersive
// n! = 1 * 2 * ... * (n - 1) * n for n >= 1
func Factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

// f(n) = 0; when n = 0
// f(n) = 1 when n = 1
// f(n) = f(n-1) + f(n-2) when n > 1
// using space optimize method
func Fibonacci(n int) int {
	a, b, c := 0, 1, 0

	if n == 0 {
		return a
	}

	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return b
}

// n power of x
// x^3 = x * x * x
func Pow(x float64, n int) float64 {

	num := 1.0
	for i := 0; i < n; i++ {
		num *= x
	}

	return num
}

// sin(x) = x/1! - x^3/3! + x^5/5! - x^7/7! + ...
// sin(x), x will be convert to radian first
func Sin(x float64) float64 {

	result, fact, power, sign := 0.0, 0.0, 0.0, 1.0

	// calculate radian
	x = x * 3.14 / 180

	for i := 1.0; i < 8; i += 2 {
		fact = Factorial(i)
		power = Pow(x, int(i))

		result += sign * (power / fact)

		sign *= -1
	}

	return result
}

// Reverse real number
// 1234 => 4321
func ReverseInt(n int) (rev int) {
	for ; n > 0; n /= 10 {
		rev = rev*10 + n%10
	}
	return
}

// 10 base to 8 base
// 10 base 93 = 5 base 135
func DicimalToOctal(n int) int {
	o := 0
	for ; n > 0; n /= 8 {
		o = o*10 + n%8
	}

	return ReverseInt(o)
}

// ascii to number
// "1234" = 1234
func CharacterToNumber(s string) (dec int) {
	//num, _ := strconv.Atoi(s)
	l := len(s)
	for i := 0; i < l; i += 1 {
		dec = dec*10 + int(s[i]) - int('0')
	}

	return
}

// square root of n
// x = x - ( x*x-n ) / 2x; initial x=2.0
func SquareRoot(n float64) float64 {
	err := 0.0000001
	x, y := 2., 0.
	for {
		x = x - (x*x-n)/(2*x)
		if math.Abs(x-y) < err {
			break
		}
		y = x
	}
	return x
}

// Smallest Divisor
func Sdivisor(x int) int {
	if x%2 == 0 {
		return 2
	}

	r := int(SquareRoot(float64(x)))
	d := 3
	for ; x%d != 0 && d < r; d = d + 2 {
	}
	if x%d == 0 {
		return d
	} else {
		return 1
	}
}

// Greatest Common Divisor
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// A Prime number is a whole number greater than 1, whose only two whole-number factors are 1 and itself.
// So Prime number can be divided by evenly only by 1, or itself, and it must be a whole number greater than 1.
// 2, 3, 5, 7, 11, 13 , 17, 19 ...
// The sieve of Eratosthenes way is use here to find prime numbers
// And Daisy - chain concurrency pattern used here
// return a list of nth position primes
func Primes(n int) []int {
	primes := make([]int, n)

	generate := func(ch chan<- int) {
		for i := 2; ; i++ {
			ch <- i
		}
	}

	filter := func(in <-chan int, out chan<- int, prime int) {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}

	prev := make(chan int)
	go generate(prev)
	for i := 0; i < n; i++ {
		primes[i] = <-prev
		next := make(chan int)
		go filter(prev, next, primes[i])
		prev = next
	}

	return primes
}

// return all prime factors of a number n
// f = p1 * p2 * p3 * ... where n > 1 and p1 <= p2 <= pc <= ...
// 60 = 2 * 2 * 3 * 5
func PrimeFactors(n int) []int {

	factors := []int{}
	primes := Primes(int(SquareRoot(float64(n))))

	for _, prime := range primes {

		for n%prime == 0 {
			n = n / prime
			factors = append(factors, prime)
		}
	}

	return factors
}

// Generate Random number using Linear Congruential
func RandomUsingLinearCongruential(x int) int {
	a := 1103515245 //109
	b := 12345      //853
	m := 123456     //4096

	if x >= 0 && x <= m-1 {
		x = (a*x + b) % m
	}
	return x
}
