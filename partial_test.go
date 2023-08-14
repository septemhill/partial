package partial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnArg1(t *testing.T) {
	sum := func(a int) int {
		return a + 100
	}

	t.Run("FnArg1 P1 Case", func(t *testing.T) {
		r1 := FnArg1(sum).P1(1)()
		assert.Equal(t, sum(1), r1)
	})
}

func TestFnArg2(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}

	t.Run("FnArg2 Wrapper Case", func(t *testing.T) {
		r1 := FnArg2[int, int, int](sum)(1, 2)
		assert.Equal(t, sum(1, 2), r1)
	})

	t.Run("FnArg2 P1 Case", func(t *testing.T) {
		r1 := FnArg2(sum).P1(1)(2)
		assert.Equal(t, sum(1, 2), r1)
	})

	t.Run("FnArg2 P2 Case", func(t *testing.T) {
		r1 := FnArg2(sum).P2(1, 2)()
		assert.Equal(t, sum(1, 2), r1)
	})
}

func TestFnArg3(t *testing.T) {
	sum := func(a, b, c int) int {
		return a + b + c
	}

	t.Run("FnArg3 Wrapper Case", func(t *testing.T) {
		r1 := FnArg3[int, int, int, int](sum)(1, 2, 3)
		assert.Equal(t, sum(1, 2, 3), r1)
	})

	t.Run("FnArg3 P1 Case", func(t *testing.T) {
		r1 := FnArg3(sum).P1(1)(2, 3)
		assert.Equal(t, sum(1, 2, 3), r1)

		r2 := FnArg3(sum).P1(1).P1(2)(3)
		assert.Equal(t, sum(1, 2, 3), r2)
	})

	t.Run("FnArg3 P2 Case", func(t *testing.T) {
		r1 := FnArg3(sum).P2(1, 2)(3)
		assert.Equal(t, sum(1, 2, 3), r1)
	})

	t.Run("FnArg3 P3 Case", func(t *testing.T) {
		r1 := FnArg3(sum).P3(1, 2, 3)()
		assert.Equal(t, sum(1, 2, 3), r1)
	})
}

func TestFnArg4(t *testing.T) {
	sum := func(a, b, c, d int) int {
		return a + b + c + d
	}

	t.Run("FnArg4 Wrapper Case", func(t *testing.T) {
		r1 := FnArg4[int, int, int, int, int](sum)(1, 2, 3, 4)
		assert.Equal(t, sum(1, 2, 3, 4), r1)
	})

	t.Run("FnArg4 P1 Case", func(t *testing.T) {
		r1 := FnArg4(sum).P1(1)(2, 3, 4)
		assert.Equal(t, sum(1, 2, 3, 4), r1)

		r2 := FnArg4(sum).P1(1).P1(2)(3, 4)
		assert.Equal(t, sum(1, 2, 3, 4), r2)

		r3 := FnArg4(sum).P1(1).P1(2).P1(3)(4)
		assert.Equal(t, sum(1, 2, 3, 4), r3)

		r4 := FnArg4(sum).P1(1).P2(2, 3)(4)
		assert.Equal(t, sum(1, 2, 3, 4), r4)
	})

	t.Run("FnArg4 P2 Case", func(t *testing.T) {
		r1 := FnArg4(sum).P2(1, 2)(3, 4)
		assert.Equal(t, sum(1, 2, 3, 4), r1)

		r2 := FnArg4(sum).P2(1, 2).P1(3)(4)
		assert.Equal(t, sum(1, 2, 3, 4), r2)
	})

	t.Run("FnArg4 P3 Case", func(t *testing.T) {
		r1 := FnArg4(sum).P3(1, 2, 3)(4)
		assert.Equal(t, sum(1, 2, 3, 4), r1)
	})

	t.Run("FnArg4 P4 Case", func(t *testing.T) {
		r1 := FnArg4(sum).P4(1, 2, 3, 4)()
		assert.Equal(t, sum(1, 2, 3, 4), r1)
	})
}

func TestFnArg5(t *testing.T) {
	sum := func(a, b, c, d, e int) int {
		return a + b + c + d + e
	}

	t.Run("FnArg5 Wrapper Case", func(t *testing.T) {
		r1 := FnArg5[int, int, int, int, int, int](sum)(1, 2, 3, 4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r1)
	})
	t.Run("FnArg5 P1 Case", func(t *testing.T) {
		r1 := FnArg5(sum).P1(1)(2, 3, 4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r1)

		r2 := FnArg5(sum).P1(1).P1(2)(3, 4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r2)

		r3 := FnArg5(sum).P1(1).P1(2).P1(3)(4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r3)

		r4 := FnArg5(sum).P1(1).P1(2).P1(3).P1(4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r4)

		r5 := FnArg5(sum).P1(1).P2(2, 3)(4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r5)

		r6 := FnArg5(sum).P1(1).P2(2, 3).P1(4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r6)

		r7 := FnArg5(sum).P1(1).P3(2, 3, 4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r7)

		r8 := FnArg5(sum).P1(1).P1(2).P2(3, 4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r8)
	})

	t.Run("FnArg5 P2 Case", func(t *testing.T) {
		r1 := FnArg5(sum).P2(1, 2)(3, 4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r1)

		r2 := FnArg5(sum).P2(1, 2).P2(3, 4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r2)

		r3 := FnArg5(sum).P2(1, 2).P1(3).P1(4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r3)

		r4 := FnArg5(sum).P2(1, 2).P1(3)(4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r4)
	})

	t.Run("FnArg5 P3 Case", func(t *testing.T) {
		r1 := FnArg5(sum).P3(1, 2, 3)(4, 5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r1)

		r2 := FnArg5(sum).P3(1, 2, 3).P1(4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r2)
	})

	t.Run("FnArg5 P4 Case", func(t *testing.T) {
		r1 := FnArg5(sum).P4(1, 2, 3, 4)(5)
		assert.Equal(t, sum(1, 2, 3, 4, 5), r1)
	})

	t.Run("FnArg5 P5 Case", func(t *testing.T) {
		r1 := FnArg5(sum).P5(1, 2, 3, 4, 5)()
		assert.Equal(t, sum(1, 2, 3, 4, 5), r1)
	})
}
