package main

import (
	"fmt"
	"strings"
)

type rectangle struct {
	width, height int
}

type square struct {
	len int
}

//increase the specified number of times
func (r *rectangle) increaseIn(num int) {
	r.width *= num
	r.height *= num
}

//draw rectangle
func drawRec(width, height int) {
	row := strings.Repeat("O", width) + "\n"
	column := strings.Repeat(row, height)
	fmt.Print(column)
}

//draw vertical or horizontal rectangle vertical==true horizontal==false
func (r rectangle) drawVariableRec(vertical bool) {
	//arrange the width and height values so that the width is always less than the height
	if r.width > r.height {
		r.width, r.height = r.height, r.width
	}
	if vertical {
		drawRec(r.width, r.height)
		return
	}
	drawRec(r.height, r.width)
	return
}
func (r rectangle) getArea() int {
	area := r.width * r.height
	return area
}
func (r rectangle) isBigger(rec rectangle) bool {
	return r.getArea() < rec.getArea()
}
func (r rectangle) getNumPossibleSquares(sq square) int {
	forWidth := r.width / sq.len
	forHeight := r.height / sq.len
	res := 0
	switch {
	case forWidth > 0 && forHeight > 0:
		res = forWidth * forHeight
		return res
	case forWidth == 0 && forHeight > 0:
		res = forHeight
		return res

	case forWidth > 0 && forHeight == 0:
		res = forWidth
		return res
	}
	return res
}
func PrettyDivisor() {
	fmt.Print(" -------------//------------- \n")
}

func main() {

	recFirst := rectangle{
		width:  3,
		height: 2,
	}
	recSecond := rectangle{
		width:  5,
		height: 5,
	}
	square := square{len: 2}

	fmt.Print("Print vertical rectangle \n")
	recFirst.drawVariableRec(true)
	PrettyDivisor()

	fmt.Print("Print horizontal rectangle \n")
	recFirst.drawVariableRec(false)
	PrettyDivisor()

	increaseNum := 2
	fmt.Printf("Increase rec in %v \n", increaseNum)
	recFirst.increaseIn(increaseNum)
	recFirst.drawVariableRec(true)
	PrettyDivisor()

	res := "No"
	if recFirst.isBigger(recSecond) {
		res = "Yes"
	}
	fmt.Printf("Is first rectangle bigger then second after increase? - %v \n", res)
	PrettyDivisor()

	res2 := recSecond.getNumPossibleSquares(square)
	fmt.Printf("The square is placed - %v times  \n", res2)

}
