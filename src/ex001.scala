object Ex001 {
  def execute(untilNumber: Int, divNumber1: Int, divNumber2: Int) =
    (1 until untilNumber).filter(n => n % divNumber1 == 0 || n % divNumber2 == 0).sum
  def main(args: Array[String]) {
    println(
      execute(1000, 3, 5) // -> 233168
    )
  }
}
