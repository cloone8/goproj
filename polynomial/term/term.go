package term

import (
    "strconv"
)

type Term struct {
    coefficient float64
    power float64
}

func NewTerm(c float64, p float64) *Term {
    t := new(Term)
    t.coefficient = c
    t.power = p

    return t
}

func (this *Term) GetValues() (c float64, p float64) {
    return this.coefficient, this.power
}

func (this *Term) Negative() *Term {
    return NewTerm(-this.coefficient, this.power)
}

func (this *Term) Add(other *Term) *Term {
    if this.power == other.power{
        return NewTerm(this.coefficient + other.coefficient, this.power)
    } else {
        return this
    }
}

func (this *Term) Subtract(other *Term) *Term {
    if this.power == other.power{
        return NewTerm(this.coefficient - other.coefficient, this.power)
    } else {
        return this
    }
}

func (this *Term) Multiply(other *Term) *Term {
    return NewTerm(this.coefficient * other.coefficient,
                   this.power + other.power)
}

func (this *Term) Divide(other *Term) *Term {
    return NewTerm(this.coefficient / other.coefficient,
                   this.power - other.power)
}

func (this *Term) Equals(other *Term) bool {
    return (this.coefficient == other.coefficient) &&
           (this.power == other.power)
}

func (this *Term) String() string {
    toReturn := ""
    c := this.coefficient
    p := this.power

    if c == 0 {
        return toReturn
    }

    if c < 0 && c != -1 {
        toReturn = "- " + strconv.FormatFloat(-c, 'g', -1, 64)
    } else if c == 1 && p == 0 {
        return toReturn + strconv.FormatFloat(c, 'g', -1, 64)
    } else if c == -1 &&  p == 0 {
        return toReturn + "- " + strconv.FormatFloat(-c, 'g', -1, 64)
    } else if c == -1 {
        toReturn = "- "
    } else if c != 1 {
        toReturn += strconv.FormatFloat(c, 'g', -1, 64)
    }


    if p == 0{
        return toReturn
    } else if p == 1 {
        toReturn += "x"
    } else {
        toReturn += "x^" + strconv.FormatFloat(p, 'g', -1, 64)
    }

    return toReturn
}
