package main

import (
	"fmt"
	"time"
)

const rows int = 11

type placeholder [rows]string

var zero = placeholder{
	"    .n~~%x.    ",
	"  x88X   888.  ",
	" X888X   8888L ",
	"X8888X   88888 ",
	"88888X   88888X",
	"88888X   88888X",
	"88888X   88888f",
	"48888X   88888 ",
	" ?888X   8888\" ",
	"  \"88X   88*`  ",
	"    ^\"===\"`    ",
}

var one = placeholder{
	"        oe     ",
	"      .@88     ",
	"  ==*88888     ",
	"     88888     ",
	"     88888     ",
	"     88888     ",
	"     88888     ",
	"     88888     ",
	"     88888     ",
	"     88888     ",
	"  '**%%%%%%**  ",
}

var two = placeholder{
	"  .--~*teu.    ",
	"  dF     988Nx ",
	" d888b   `8888>",
	" ?8888>  98888F",
	"  \"**\"  x88888~",
	"       d8888*` ",
	"     z8**\"`   :",
	"   :?.....  ..F",
	"  <\"\"888888888~",
	"  8:  \"888888* ",
	"  \"\"    \"**\"`  ",
}

var three = placeholder{
	"  .x~~\"*Weu.   ",
	"  d8Nu.  9888c ",
	"  88888  98888 ",
	"  \"***\"  9888% ",
	"       ..@8*\"  ",
	"    ````\"8Weu  ",
	"   ..    ?8888L",
	" :@88N   '8888N",
	" *8888~  '8888F",
	" '*8\"`   9888% ",
	"   `~===*%\"`   ",
}

var four = placeholder{
	"        xeee   ",
	"       d888R   ",
	"      d8888R   ",
	"     @ 8888R   ",
	"   .P  8888R   ",
	"  :F   8888R   ",
	" x\"    8888R   ",
	"d8eeeee88888eer",
	"       8888R   ",
	"       8888R   ",
	"    \"*%%%%%%**~",
}

var five = placeholder{
	"  cuuu....uK   ",
	"  888888888    ",
	"  8*888**\"     ",
	"  >  .....     ",
	"  Lz\"  ^888Nu  ",
	"  F     '8888k ",
	"  ..     88888>",
	" @888L   88888 ",
	"'8888F   8888F ",
	" %8F\"   d888\"  ",
	"  ^\"===*%\"`    ",
}

var six = placeholder{
	"     .ue~~%u.  ",
	"   .d88   z88i ",
	"  x888E  *8888 ",
	" :8888E   ^\"\"  ",
	" 98888E.=tWc.  ",
	" 98888N  '888N ",
	" 98888E   8888E",
	" '8888E   8888E",
	"  ?888E   8888\"",
	"   \"88&   888\" ",
	"     \"\"==*\"\"   ",
}

var seven = placeholder{
	" dL ud8Nu  :8c ",
	" 8Fd888888L %8 ",
	" 4N88888888cuR ",
	" 4F   ^\"\"%\"\"d  ",
	" d       .z8   ",
	" ^     z888    ",
	"     d8888'    ",
	"    888888     ",
	"   :888888     ",
	"    888888     ",
	"    '%**%      ",
}

var eight = placeholder{
	"   u+=~~~+u.   ",
	" z8F      `8N. ",
	"d88L       98E ",
	"98888bu.. .@*  ",
	"\"88888888NNu.  ",
	" \"*8888888888i ",
	" .zf\"\"*8888888L",
	"d8F      ^%888E",
	"88>        `88~",
	"'%N.       d*\" ",
	"   ^\"=====\"`   ",
}

var nine = placeholder{
	"  .xn!~%x.     ",
	"  x888   888.  ",
	" X8888   8888: ",
	" 88888   X8888 ",
	" 88888   88888>",
	" `8888  :88888X",
	"   `\"**~ 88888>",
	"  .xx.   88888 ",
	" '8888>  8888~ ",
	"  888\"  :88%   ",
	"   ^\"===\"\"     ",
}

var colon = placeholder{
	"               ",
	"        .      ",
	"       d8c     ",
	"     ^*888%    ",
	"       \"8      ",
	"               ",
	"        .      ",
	"       d8c     ",
	"     ^*888%    ",
	"       \"8      ",
	"               ",
}

var space = placeholder{
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
}

var digits = []placeholder{
	zero,
	one,
	two,
	three,
	four,
	five,
	six,
	seven,
	eight,
	nine,
}

func number2ascii(n int) (a, b placeholder) {
	if n < 0 || n > 59 {
		panic("number2ascii can only be used with numbers from 0 to 59")
	}

	leftDigit := n / 10
	rightDigit := n % 10

	return digits[leftDigit], digits[rightDigit]
}

func printClock(t time.Time) {

	h1, h2 := number2ascii(t.Hour())
	m1, m2 := number2ascii(t.Minute())
	s1, s2 := number2ascii(t.Second())
	shouldRenderColon := t.Second()%2 == 0

	separator := space
	if shouldRenderColon {
		separator = colon
	}

	for r := 0; r < rows; r++ {
		fmt.Printf("%s %s %s %s %s %s %s %s\n",
			h1[r], h2[r],
			separator[r],
			m1[r], m2[r],
			separator[r],
			s1[r], s2[r])
	}
}

func main() {
	for {
		printClock(time.Now())
		time.Sleep(time.Second)
		fmt.Printf("\033[11A")
	}
}
