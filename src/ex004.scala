object Ex004 {
  def isParotum(s: String): Boolean = s == s.reverse
  def maxParotum(xs: List[Int], ys: List[Int]): Int =
    (for (x <- xs; y <- ys) yield x * y).filter(n => isParotum(n.toString)).max
  def execute(): Int = {
    maxParotum((100 until 1000).toList, (100 until 1000).toList)
  }
  def main(args: Array[String]) {
    println(execute()) // => 906609
  }
}

