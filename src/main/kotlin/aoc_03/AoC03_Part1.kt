package aoc_03

import java.io.File

fun main() {
    val infile = File("src/main/kotlin/aoc_03/in.txt")

    val input = infile.readText()

    val regex = Regex("""mul\((\d{1,3}),(\d{1,3})\)""")

    val sum = regex.findAll(input).map { match ->
        val (a, b) = match.groupValues.drop(1).map { it.toInt() }
        a * b
    }.sum()


    println("Sum of multiples: $sum")
}
