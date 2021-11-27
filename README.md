colrz
==

[![GoDoc](https://godoc.org/github.com/unnamedxaer/colrz?status.svg)](https://godoc.org/github.com/unnamedxaer/colrz)
[![Build Status](https://travis-ci.org/unnamedxaer/colrz.svg)](https://travis-ci.org/unnamedxaer/colrz)
[![GoReportCard](https://goreportcard.com/badge/unnamedxaer/colrz)](https://goreportcard.com/report/unnamedxaer/colrz)
<!-- [![Gitter](https://img.shields.io/badge/chat-on_gitter-46bc99.svg?logo=data:image%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGhlaWdodD0iMTQiIHdpZHRoPSIxNCI%2BPGcgZmlsbD0iI2ZmZiI%2BPHJlY3QgeD0iMCIgeT0iMyIgd2lkdGg9IjEiIGhlaWdodD0iNSIvPjxyZWN0IHg9IjIiIHk9IjQiIHdpZHRoPSIxIiBoZWlnaHQ9IjciLz48cmVjdCB4PSI0IiB5PSI0IiB3aWR0aD0iMSIgaGVpZ2h0PSI3Ii8%2BPHJlY3QgeD0iNiIgeT0iNCIgd2lkdGg9IjEiIGhlaWdodD0iNCIvPjwvZz48L3N2Zz4%3D&logoWidth=10)](https://gitter.im/unnamedxaer/colrz?utm_source=share-link&utm_medium=link&utm_campaign=share-link) -->

It's a set of funcs and constants to provide basic colors to your terminal app.

# How to use

Get it

```bash
go get github.com/unnamedxaer/colrz
```

Use it
```go
package main

import(
	"fmt"

	c "github.com/unnamedxaer/colrz"
)

func main(){
	// with method
	fmt.Print(c.Red("\nThe tree is red."))

	// with const 
	fmt.Printf("\n%sThe tree is reds.%s", c.FgRed, c.FgReset)

	// background color
	fmt.Printf("\n%s", c.BlueBg(c.Red("The tree is red with blue leafs.")))
}
```

```go
fmt.Print(c.Red("The tree is red."))
```

![red text](https://example.com/img1.png)

```go
fmt.Printf("\n%s", c.BlueBg(c.Red("The tree is red with blue leafs.")))
```

![red text with blue background color](https://example.com/img1.png)

# License

This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.