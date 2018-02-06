object Ex002 {
  def fib(n: Int): Int =
    if (n <= 1) n else fib(n - 1) + fib(n - 2)
  def compute(maxNum: Int, maxIndex: Int): Int =
    (0 to maxIndex).map(fib).filter(f => f % 2 == 0).takeWhile(f => f <= maxNum).sum
  def main(args: Array[String]) {
    val n = 40
    println(fib(n) > 4000000) // => true
    println(compute(4000000, n)) // => 4613732
  }
}

