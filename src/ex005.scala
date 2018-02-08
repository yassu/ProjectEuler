import math.{pow, log, floor}

object Ex005 {
  def isPrime(p: Int): Boolean =
    p > 1 &&
    ! (2 until p).exists(p % _ == 0)
  def getPrimes(maxNumber: Int): List[Int] =
    (2 to maxNumber).filter(isPrime).toList
  def execute(maxNumber: Int): Int =
    getPrimes(maxNumber).map(p => pow(p, floor(log(maxNumber)/log(p))).toInt).product
  def main(args: Array[String]) {
    println(execute(20)) // => 232792560
  }
}

