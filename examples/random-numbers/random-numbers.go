// `math/rand/v2` 包提供[伪随机数](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)生成能力。

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// 例如 `rand.IntN` 会返回满足 `0 <= n < 100` 的随机整数。
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` 返回满足 `0.0 <= f < 1.0` 的 `float64`。
	fmt.Println(rand.Float64())

	// 可以据此生成其他范围内的随机浮点数，例如 `5.0 <= f' < 10.0`。
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 如果需要可预期的种子，可以创建新的 `rand.Source` 并传入 `New` 构造函数。
	// `NewPCG` 会创建一个新的[PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator) 源，
	// 需要两个 `uint64` 作为种子。
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
