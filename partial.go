package partial

type fnArg0[A any] func() A

func FnArg0[A any](fn fnArg0[A]) fnArg0[A] {
	return fn
}

type fnArg1[A, B any] func(A) B

func FnArg1[A, B any](fn fnArg1[A, B]) fnArg1[A, B] {
	return fn
}

func (fn fnArg1[A, B]) P1(a A) fnArg0[B] {
	return func() B {
		return fn(a)
	}
}

type fnArg2[A, B, C any] func(A, B) C

func FnArg2[A, B, C any](fn fnArg2[A, B, C]) fnArg2[A, B, C] {
	return fn
}

func (fn fnArg2[A, B, C]) P1(a A) fnArg1[B, C] {
	return func(b B) C {
		return fn(a, b)
	}
}

func (fn fnArg2[A, B, C]) P2(a A, b B) fnArg0[C] {
	return func() C {
		return fn(a, b)
	}
}

type fnArg3[A, B, C, D any] func(A, B, C) D

func FnArg3[A, B, C, D any](fn fnArg3[A, B, C, D]) fnArg3[A, B, C, D] {
	return fn
}

func (fn fnArg3[A, B, C, D]) P1(a A) fnArg2[B, C, D] {
	return func(b B, c C) D {
		return fn(a, b, c)
	}
}

func (fn fnArg3[A, B, C, D]) P2(a A, b B) fnArg1[C, D] {
	return func(c C) D {
		return fn(a, b, c)
	}
}

func (fn fnArg3[A, B, C, D]) P3(a A, b B, c C) fnArg0[D] {
	return func() D {
		return fn(a, b, c)
	}
}

type fnArg4[A, B, C, D, E any] func(A, B, C, D) E

func FnArg4[A, B, C, D, E any](fn fnArg4[A, B, C, D, E]) fnArg4[A, B, C, D, E] {
	return fn
}

func (fn fnArg4[A, B, C, D, E]) P1(a A) fnArg3[B, C, D, E] {
	return func(b B, c C, d D) E {
		return fn(a, b, c, d)
	}
}

func (fn fnArg4[A, B, C, D, E]) P2(a A, b B) fnArg2[C, D, E] {
	return func(c C, d D) E {
		return fn(a, b, c, d)
	}
}

func (fn fnArg4[A, B, C, D, E]) P3(a A, b B, c C) fnArg1[D, E] {
	return func(d D) E {
		return fn(a, b, c, d)
	}
}

func (fn fnArg4[A, B, C, D, E]) P4(a A, b B, c C, d D) fnArg0[E] {
	return func() E {
		return fn(a, b, c, d)
	}
}

type fnArg5[A, B, C, D, E, F any] func(A, B, C, D, E) F

func FnArg5[A, B, C, D, E, F any](fn fnArg5[A, B, C, D, E, F]) fnArg5[A, B, C, D, E, F] {
	return fn
}

func (fn fnArg5[A, B, C, D, E, F]) P1(a A) fnArg4[B, C, D, E, F] {
	return func(b B, c C, d D, e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn fnArg5[A, B, C, D, E, F]) P2(a A, b B) fnArg3[C, D, E, F] {
	return func(c C, d D, e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn fnArg5[A, B, C, D, E, F]) P3(a A, b B, c C) fnArg2[D, E, F] {
	return func(d D, e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn fnArg5[A, B, C, D, E, F]) P4(a A, b B, c C, d D) fnArg1[E, F] {
	return func(e E) F {
		return fn(a, b, c, d, e)
	}
}

func (fn fnArg5[A, B, C, D, E, F]) P5(a A, b B, c C, d D, e E) fnArg0[F] {
	return func() F {
		return fn(a, b, c, d, e)
	}
}
