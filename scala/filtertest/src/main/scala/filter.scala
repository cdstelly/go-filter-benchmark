import com.github.nscala_time.time.Imports._
import scala.collection.mutable.ArrayBuffer

object filter {
  val numFiles = 10000000
  //Execute function passed in; time it.
  def time[R](block: => R): R = {
    val t0 = System.nanoTime()
    val result = block    // call-by-name
    val t1 = System.nanoTime()
    val elapsed = (t1-t0).toFloat / 1000000
    val fps = numFiles / elapsed
    println(f"Elapsed time: $elapsed%2.5f ms, $fps%2.5f files/ms")
    result
  }

  //Simulate file metadata
  class metafile(var Name: String = generateFilename(), var CreateTime: DateTime = generateDateTime(), var EditTime: DateTime = generateDateTime()) {
    override def toString: String = s"Filename: $Name, Create: $CreateTime, EditTime: $EditTime"
  }

  def generateDateTime(): DateTime = {
    val r = scala.util.Random
    val startTimestamp = new DateTime(math.abs(r.nextLong()))
    return startTimestamp
  }

  def generateFilename(): String = {
    var fn = scala.util.Random.alphanumeric.take(15).mkString

    val A = Array("pdf", "exe", "png", "ini", "txt", "docx", "xls", "bmp", "wav", "mp3")
    val random_index = scala.util.Random.nextInt(A.length)
    val result = A(random_index)

    return fn + "." + result
  }

  def main(args: Array[String]): Unit = {
    var fileList = ArrayBuffer[metafile]()

    for(_ <- 1 to numFiles) {
      val x = new metafile()
      //println(x)
      fileList += x
    }
    val filterresult = time { fileList.filter(_.Name.endsWith("pdf"))}
    val mapresult = time { fileList.map(_.Name.concat(".bk" ))}
  }
}
