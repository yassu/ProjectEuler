import math.sqrt

object Ex003 {
  def isFactorWithInt(n: Long, q: Long): (Boolean, Long) =
    if (n % q == 0) (true, isFactorWithInt(n / q, q)._2) else (false, n)
  def factors(n: Long): List[Long] = {
    def _factors(n: Long, r: Long): List[Long] = {
      val (b, res) = isFactorWithInt(n, r)
      if (res == 1) r:: Nil
      else if (b) r :: _factors(res, r + 1)
      else _factors(res, r + 1)
    }
    _factors(n, 2)
  }
  def execute(n: Long): Long = {
    factors(n).max
  }
  def main(args: Array[String]) {
    println(execute(600851475143l)) // => 6857
  }
}

