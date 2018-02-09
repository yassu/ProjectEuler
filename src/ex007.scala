object Ex007 {
  // raise stackoverflow by recursive
  def primes(cnt: Int): List[Long] = {
    if (cnt == 1) return List(2)
    val beforePrimes = primes(cnt - 1)
    val lastPrime = beforePrimes.last
    val nextPrime = (lastPrime + 1 to 2 * lastPrime).
      filter(n => beforePrimes.forall(n % _ != 0)).head
    println(nextPrime)
    return beforePrimes :+ nextPrime
  }
  def execute(count: Int): Long =
    primes(count).last
  def main(args: Array[String]) {
    println(execute(7000)) // => 25164150
  }
}

