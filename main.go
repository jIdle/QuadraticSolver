package main

import (
    "fmt"
    "bufio"
    "os"
    "math/cmplx"
    "strconv"
    "log"
)

var pl = fmt.Println
var pf = fmt.Printf
var Flt = strconv.ParseFloat
var sqrt = cmplx.Sqrt
var pow = cmplx.Pow
var reader = bufio.NewReader(os.Stdin)

func quadratic(a, b, c complex128) (complex128, complex128) {
    root1 := (-b + sqrt(pow(b,2) - 4*a*c))/(2*a)
    root2 := (-b - sqrt(pow(b,2) - 4*a*c))/(2*a)
    return root1, root2
}

func main() {
    pl("Enter a: ")
    aStr, _ := reader.ReadString('\n')
    aStr = aStr[:len(aStr)-1]
    pl("Enter b: ")
    bStr, _ := reader.ReadString('\n')
    bStr = bStr[:len(bStr)-1]
    pl("Enter c: ")
    cStr, _ := reader.ReadString('\n')
    cStr = cStr[:len(cStr)-1]

    aFl, err1 := Flt(aStr, 64)
    if err1 != nil {
        log.Fatal(err1)
    }
    bFl, err2 := Flt(bStr, 64)
    if err2 != nil {
        log.Fatal(err2)
    }
    cFl, err3 := Flt(cStr, 64)
    if err3 != nil {
        log.Fatal(err3)
    }

    aCom := complex(aFl, 0)
    bCom := complex(bFl, 0)
    cCom := complex(cFl, 0)
    root1, root2 := quadratic(aCom, bCom, cCom)

    pf(
        "\nThe roots of this quadratic are:\n\troot1: %v\n\troot2: %v\n",
        root1,
        root2,
    )
}
