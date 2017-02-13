package main

import (
    "personal/polynomial/term"
    "fmt"
    )

func main() {
    t1 := term.NewTerm(-0.5, 0)
    t2 := term.NewTerm(-3, 6)
    t3 := term.NewTerm(3, 6)
    t4 := term.NewTerm(3, -2.1)
    t5 := term.NewTerm(3, -6)
    t6 := term.NewTerm(0, 6)
    t7 := term.NewTerm(4, 1)
    t8 := term.NewTerm(-3, -6)
    t9 := term.NewTerm(1, 1)

    fmt.Println(t1)
    fmt.Println(t2)
    fmt.Println(t3)
    fmt.Println(t4)
    fmt.Println(t5)
    fmt.Println(t6)
    fmt.Println(t7)
    fmt.Println(t8)
    fmt.Println(t9)
}
