object Ex006 {
  def square(n: Int): Int = n * n
  def execute(n: Int): Int =
    square((1 to n).sum) - (1 to n).map(square).sum
  def main(args: Array[String]) {
    println(execute(100)) // => 25164150
  }
}

