package main

import (
	"fmt"
	"math"
)

func calculation(xArr, yArr []float64, operation int) {
	intN := len(xArr)
	n := float64(intN)

	xSum, ySum, x2Sum, xySum := 0.0, 0.0, 0.0, 0.0
	a0, a1, St, Sy, Sr, Sxy, r2, r := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	if operation == 1 { //Power Model
		for i := 0; i < intN; i++ {
			xArr[i] = round(math.Log10(xArr[i]), 4)
			xSum += xArr[i]
			x2Sum += math.Pow(xArr[i], 2)

			yArr[i] = round(math.Log10(yArr[i]), 4)
			ySum += yArr[i]

			xySum += round(xArr[i]*yArr[i], 4)
			fmt.Printf("Row%v: %v\t%v\t%v\t%v\n", i, xArr[i], yArr[i], math.Pow(xArr[i], 2), xArr[i]*yArr[i])
		}
	} else if operation == 2 { //Exponential Model
		for i := 0; i < intN; i++ {
			xSum += xArr[i]
			x2Sum += math.Pow(xArr[i], 2)

			yArr[i] = math.Log(yArr[i])
			ySum += yArr[i]

			xySum += xArr[i] * yArr[i]
			fmt.Printf("Row%v: %v\t%v\t%v\t%v\n", i, xArr[i], yArr[i], math.Pow(xArr[i], 2), xArr[i]*yArr[i])
		}
	} else if operation == 3 { // Saturation-Growth Model
		for i := 0; i < intN; i++ {
			xArr[i] = 1 / xArr[i]
			xSum += xArr[i]
			x2Sum += math.Pow(xArr[i], 2)

			yArr[i] = 1 / yArr[i]
			ySum += yArr[i]

			xySum += xArr[i] * yArr[i]
			fmt.Printf("Row%v: %v\t%v\t%v\t%v\n", i, xArr[i], yArr[i], math.Pow(xArr[i], 2), xArr[i]*yArr[i])
		}

	} else if operation == 4 { // Least Square Regression
		for i := 0; i < intN; i++ {
			xSum += xArr[i]
			ySum += yArr[i]
			xySum += xArr[i] * yArr[i]
			x2Sum += math.Pow(xArr[i], 2)
			fmt.Printf("Row%v: %v\t%v\t%v\t%v\n", i, xArr[i], yArr[i], math.Pow(xArr[i], 2), xArr[i]*yArr[i])
		}

	} else {
		fmt.Println("Invalid Operation!")
		return
	}

	a1 = ((n * xySum) - (xSum * ySum)) / (n*x2Sum - math.Pow(xSum, 2))
	a0 = (ySum / n) - (a1 * (xSum / n))

	for i := 0; i < intN; i++ {
		Sr += math.Pow((yArr[i] - a0 - (a1 * xArr[i])), 2)
		St += math.Pow((yArr[i] - ySum/n), 2)
	}
	Sxy = math.Sqrt(Sr / (n - 2))
	Sy = math.Sqrt(St / (n - 1))
	r2 = (St - Sr) / St
	r = math.Sqrt(r2)

	fmt.Printf("X Sum = %v\tY Sum = %v\tX x Y Sum = %v\tX^2 Sum = %v\n", xSum, ySum, xySum, x2Sum)
	fmt.Printf("a0 = %v\ta1 = %v\n", a0, a1)
	fmt.Printf("St = %v\tSy = %v\n", St, Sy)
	fmt.Printf("Sr = %v\tSxy = %v\n", Sr, Sxy)
	fmt.Printf("r2 = %v\tr = %v\n", r2, r)
}
func integration(a, b float64, n int) {
	x := a
	counter := 0
	var arr []float64
	h := (b - a) / float64(n)
	for x <= b {
		fx := math.Sqrt(64 - (x * x))
		arr = append(arr, fx)
		fmt.Printf("X%v: %v %v \n", counter, x, fx)
		x += h
		counter++
	}
	matr(h, n, arr)
	simpson13(h, n, arr)
	simpson38(h, n, arr)
}
func sum(initial, end, rate int, arr []float64) (sum float64, str string) {
	sum = 0.0
	str = ""
	i := initial
	for i < end {
		sum += arr[i]
		if i == initial {
			str += fmt.Sprintf("%f", arr[i])
		} else {
			str += " + " + fmt.Sprintf("%f", arr[i])
		}
		i += rate
	}
	return sum, str
}
func matr(h float64, n int, arr []float64) {
	sum, str := sum(1, n, 1, arr)
	matrSum := (h * (arr[0] + arr[n] + 2*sum)) / 2
	fmt.Printf("I = (%v/2) x (%v + %v + 2x(%v))", h, arr[0], arr[n], str)
	fmt.Printf("\nMatr Sum: %v\n", matrSum)
}
func simpson13(h float64, n int, arr []float64) {
	sum1, str1 := sum(1, n, 2, arr)
	sum2, str2 := sum(2, n, 2, arr)

	simpson13Sum := (h * (arr[0] + arr[n] + 4*sum1 + 2*sum2)) / 3
	fmt.Printf("I = (%v/3) x (%v + %v + 4x(%v) + 2x(%v))", h, arr[0], arr[n], str1, str2)
	fmt.Printf("\nSimpson 1/3 Sum: %v\n", simpson13Sum)
}
func simpson38(h float64, n int, arr []float64) {
	sum1, str1 := sum(1, n, 3, arr)
	sum2, str2 := sum(2, n, 3, arr)
	sum3, str3 := sum(3, n, 3, arr)

	simpson38Sum := (3 * h / 8) * (arr[0] + arr[n] + 3*sum1 + 3*sum2 + 2*sum3)
	fmt.Printf("I = (3x%v/8) x (%v + %v + 3x(%v + %v) + 2x(%v))", h, arr[0], arr[n], str1, str2, str3)
	fmt.Printf("\nSimpson 3/8 Sum: %v\n", simpson38Sum)
}

func factorial(n int) float64 {
	result := 1.0
	for i := 1; i <= n; i++ {
		result *= float64(i)
	}
	return result
}
func round(number float64, n int) float64 {
	pow := float64(n)
	return math.Round((number * math.Pow(10, pow))) / math.Pow(10, pow)
}
func incrementalSearch(iteration int, start, step float64) {
	val := float64(start)
	xr := 0.0
	for i := 0; i < iteration; i++ {
		fmt.Printf("Iteration %v: \n", i+1)
		fmt.Printf("x:%v\tf(x):%v\n", round(val, 3), round(fx(val), 3))
		for fx(val)*fx(val+step) > 0 {
			val += step
			fmt.Printf("X:%v\tf(x):%v\n", round(val, 3), round(fx(val), 3))
		}
		fmt.Printf("X:%v\tf(x):%v\n", round(val+step, 3), round(fx(val+step), 3))
		xrPrev := xr
		xr = (val + val + step) / 2
		fmt.Printf("Xr = %v\n", xr)
		fmt.Printf("Ea:%v\n", round(ea(xr, xrPrev), 2))
		step /= 10
	}
}
func bisection(start, end float64, iteration int) {
	xu := end
	xl := start
	xr := 0.0
	for i := 0; i < iteration; i++ {
		xrPrev := xr
		xr = (xl + xu) / 2
		fmt.Printf("Xl: %v, Xu: %v, Xr: %v,F(xl): %v, F(xu): %v F(xr): %v, Ea:%v\n", xl, xu, xr, fx(xl), fx(xu), fx(xr), round(ea(xr, xrPrev), 2))
		if fx(xr)*fx(xl) < 0 {
			xu = xr
		} else if fx(xr)*fx(xl) > 0 {
			xl = xr
		} else {
			break
		}
	}
}
func falsePosition(start, end float64, iteration int) {
	xu := end
	xl := start
	xr := 0.0
	for i := 0; i < iteration; i++ {
		xrPrev := xr
		xr = xu - ((fx(xu) * (xl - xu)) / (fx(xl) - fx(xu)))
		fmt.Printf("Xl: %v, Xu: %v, Xr: %v, F(xu): %v, F(xr): %v, Ea:%v\n", xl, xu, xr, fx(xu), fx(xr), round(ea(xr, xrPrev), 2))
		//fmt.Printf("Et: %v\n", (math.Abs(trueValue-xr)/trueValue)*100)
		if fx(xr)*fx(xl) < 0 {
			xu = xr
		} else if fx(xr)*fx(xl) > 0 {
			xl = xr
		} else {
			break
		}
	}
}
func secant(start, end float64, iteration int) {
	xMinusOne := start
	xi := end
	for i := 0; i < iteration; i++ {
		xPlusOne := xi - ((fx(xi) * (xMinusOne - xi)) / (fx(xMinusOne) - fx(xi)))
		fmt.Printf("Xi-1: %v, Xi: %v, f(Xi-1): %v, f(Xi): %v, Xi+1: %v, Ea: %v\n", xMinusOne, xi, fx(xMinusOne), fx(xi), xPlusOne, round(ea(xPlusOne, xi), 2))
		xMinusOne = xi
		xi = xPlusOne
	}
}
func newton(start float64, iteration int) {
	xi := start
	for i := 0; i < iteration; i++ {
		xPlusOne := xi - fx(xi)/fxd(xi)
		fmt.Printf("Xi+1: %v, Xi: %v, f(Xi): %v, f'(Xi): %v, Xi+1: %v, Ea: %v\n", xPlusOne, xi, fx(xi), fxd(xi), xPlusOne, round(ea(xPlusOne, xi), 2))
		xi = xPlusOne
	}
}

func fx(x float64) float64 {
	//return (4.15 * x * x) - (16 * x) + 8 //incremental search
	//return (-0.6 * x * x) + (2.4 * x) + 5.5 //bisection
	//return (-3 * math.Pow(x, 3)) + (19 * x * x) - 20*x - 13 //false position
	//return math.Pow(x, 3) - 2*x*x - 5
	//return -12 - (21 * x) + (18 * x * x) - (2.4 * math.Pow(x, 3)) //secant
	return math.Pow(math.E, -x) * (1.5*math.Sin(x) - 0.5*math.Cos(x))
}
func fxd(x float64) float64 {
	return 1
}
func ea(current, old float64) float64 {
	return math.Abs((current-old)/current) * 100
}
