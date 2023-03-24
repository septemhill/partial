package partial

type FnArg1[A, B any] func(A) B

type FnArg2[A, B, C any] func(A, B) C

func (fn FnArg2[A, B, C]) P1(a A) FnArg1[B, C] {
	return func(b B) C {
		return fn(a, b)
	}
}

type FnArg3[A, B, C, D any] func(A, B, C) D

func (fn FnArg3[A, B, C, D]) P1(a A) FnArg2[B, C, D] {
	return func(b B, c C) D {
		return fn(a, b, c)
	}
}

func (fn FnArg3[A, B, C, D]) P2(a A, b B) FnArg1[C, D] {
	return func(c C) D {
		return fn(a, b, c)
	}
}

type FnArg4[A, B, C, D, E any] func(A, B, C, D) E

func (fn FnArg4[A, B, C, D, E]) P1(a A) FnArg3[B, C, D, E] {
	return func(b B, c C, d D) E {
		return fn(a, b, c, d)
	}
}

func (fn FnArg4[A, B, C, D, E]) P2(a A, b B) FnArg2[C, D, E] {
	return func(c C, d D) E {
		return fn(a, b, c, d)
	}
}

func (fn FnArg4[A, B, C, D, E]) P3(a A, b B, c C) FnArg1[D, E] {
	return func(d D) E {
		return fn(a, b, c, d)
	}
}

type FnArg5[A, B, C, D, E, F any] func(A, B, C, D, E) F

func (fn FnArg5[A, B, C, D, E, F]) P1(a A) FnArg4[B, C, D, E, F] {
	return func(b B, c C, d D, e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn FnArg5[A, B, C, D, E, F]) P2(a A, b B) FnArg3[C, D, E, F] {
	return func(c C, d D, e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn FnArg5[A, B, C, D, E, F]) P3(a A, b B, c C) FnArg2[D, E, F] {
	return func(d D, e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn FnArg5[A, B, C, D, E, F]) P4(a A, b B, c C, d D) FnArg1[E, F] {
	return func(e E) F {
		return fn(a, b, c, d, e)
	}
}
