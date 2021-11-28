colrz
==

[![GoDoc](https://godoc.org/github.com/unnamedxaer/colrz?status.svg)](https://godoc.org/github.com/unnamedxaer/colrz)
[![Build Status](https://travis-ci.org/unnamedxaer/colrz.svg)](https://travis-ci.org/unnamedxaer/colrz)
[![GoReportCard](https://goreportcard.com/badge/unnamedxaer/colrz)](https://goreportcard.com/report/unnamedxaer/colrz)

It's a set of funcs and constants to provide basic colors to your terminal app.

# How to use

**Get it**

```bash
go get github.com/unnamedxaer/colrz
```

**Use it**
```go
package main

import (
	"fmt"

	c "github.com/unnamedxaer/colrz"
)

func main() {
	// with method
	fmt.Print(c.Yellow("\nThe tree is yellow."))

	// with const
	fmt.Printf("\n%sThe tree is red.%s", c.FgRed, c.FgReset)

	// background color
	fmt.Printf("\n%s", c.BlueBg(c.Yellow("The tree is green with blue leafs.")))

	fmt.Printf("\n%s", c.WhiteBg(c.Green("The tree is green and ")+c.Yellow("yellow with white leafs.")))

	fmt.Printf("\n%s", c.Yellow(c.MagentaBg("The tree is yellow ")+c.CyanBg(" with colorful leafs.")))

	fmt.Printf("\n%s%sThe tree is %s yellow with %scolorful leafs.%s%s",
		c.BgMagenta, c.FgYellow, c.BgCyan, c.FgBack, c.BgReset, c.FgReset)
}
```

![output from above code snippet](/../../../../unnamedxaer/assets/blob/main/images/colrz/example.png)

* Use colored text
```go
fmt.Print(c.Red("The tree is red."))
```
![red text](/../../../../unnamedxaer/assets/blob/main/images/colrz/red_text.png)

* Use colored text on colored background
```go
fmt.Printf("\n%s", c.BlueBg(c.Red("The tree is red with blue leafs.")))
```

![red text on blue background color](/../../../../unnamedxaer/assets/blob/main/images/colrz/red_text_blue_bg.png)

## Persistent coloring

If you set color via constant and do not reset it, 
then the color will persist for the next prints.

```go
	fmt.Printf("\n%sThe tree is red.", c.FgRed)
	// ... some code
	fmt.Printf("\ndone at %v", time.Now().Local())
	// ... some code maybe more prints
	// restore color to normal
	fmt.Print(c.FgReset)
	fmt.Print("\nThe winter is coming, the leaves have fallen.")
```

![persistent text color](/../../../../unnamedxaer/assets/blob/main/images/colrz/persistent_color.png)

## All colors

![all colors example](/../../../../unnamedxaer/assets/blob/main/images/colrz/all_colors.png)

## Colors in different terminals

Colors may and will vary slightly depending on the terminals.

![cmder colors example](/../../../../unnamedxaer/assets/blob/main/images/colrz/example_cmder.png)

## Warning
Not all terminals support this features eg. windows cmd / ps < 7.0

![example of terminal not supporting colors](/../../../../unnamedxaer/assets/blob/main/images/colrz/unsupported_terminal.png)

# License

MIT. See the LICENSE file for more details.