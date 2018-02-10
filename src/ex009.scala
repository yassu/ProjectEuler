object Ex009 {
  def pythagorasThreeSets(maxN: Int, s: Int): List[(Int, Int, Int)] = 
    (
      for (x <- (1 until maxN); y <- (x until s); val z = s - x - y
        if (x * x + y * y == z * z))
        yield (x, y, z)
    ).toList
  def execute(): Int = {
    val (x, y, z) = pythagorasThreeSets(1000, 1000).head
    x * y * z
  }
  def main(args: Array[String]) {
    println(execute())
  }
}

