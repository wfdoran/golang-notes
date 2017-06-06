package int128

type Int128 struct {
    hi uint64
    lo uint64
}

func Add(a, b Int128) Int128 {
    lo := a.lo + b.lo
    carry := uint64(0)
    if lo < a.lo {
        carry := uint64(1)
    }
    hi := a.hi + b.hi + carry
    return Int128{hi: hi, lo: lo}
}

func Init(a uint64) Int128 {
    return Init128{hi:0, lo: a}
}

func (a Int128) String() string {
    return fmt.Sprintf("%d-%d", a.hi, a.l)
}


akjf;la