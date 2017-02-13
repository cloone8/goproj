package polynomial

import (
    "personal/polynomial/term"
    "sort"
)

type Polynomial []Term

func (this *Polynomial) Len() int {
    return len(this)
}

func (this *Polynomial) Less(i, j int) bool {
    if this[i].power > this[j].power {
        return true
    } else if this[i].power == this[j].power &&
              this[i].coefficient >= this[j].power {
        return true
    } else {
        return false
    }
}

func (this *Polynomial) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func NewPolynomial(numbers ...float64) *Polynomial {
    p := new(Polynomial)

    for i := 0; i < len(numbers)/2; i++ {
        p = append(p, term.NewTerm(numbers[i * 2], numbers[(i * 2) + 1]))
    }
    p.simplify()
    return p
}

func (this *Polynomial) simplify() {
    sort.Sort(this)
    pos := 1

    for pos < len(this) {
        if this[pos - 1].power == this[pos].power {
            this[pos -1] = this[pos - 1].Add(this[pos])
            this = append(this[:pos], this[pos + 1:]...)
        } else {
            pos++
        }
    }

    for i := 0; i < len(this); i++ {
        if this[i].coefficient == 0 {
            this = append(this[:i], this[i + 1:]...)
        }
    }
}

func (this *Polynomial) negative() *Polynomial {
    p := new(Polynomial)

    for i := 0; i < len(this); i++ {
        p = append(p, this[i].negative())
    }

    return NewPolynomial(p)
}

func (this *Polynomial) Add(other *Polynomial) *Polynomial {
    p := append(this, other)
    return NewPolynomial(p)
}

func (this *Polynomial) Subtract(other *Polynomial) *Polynomial {
    p := append(this, other.negative())
    return NewPolynomial(p)
}

func (this *Polynomial) Multiply(other *Polynomial) *Polynomial {
    p := new(Polynomial)

    for i := 0; i < len(this); i++ {
        for j := 0; j < len(other); j++ {
            p = append(p, this[i].Multiply(other[i]))
        }
    }

    return NewPolynomial(p)
}

func (this *Polynomial) Divide(other *Polynomial) *Polynomial {
    p := new(Polynomial)

    for i := 0; i < len(this); i++ {
        for j := 0; j < len(other); j++ {
            p = append(p, this[i].Divide(other[i]))
        }
    }

    return NewPolynomial(p)
}

func (this *Polynomial) Differentiate() *Polynomial {
    p := new(Polynomial)
}
