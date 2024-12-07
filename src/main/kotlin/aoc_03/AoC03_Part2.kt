package aoc_03

import java.io.File

fun main() {
    val infile = File("src/main/kotlin/aoc_03/in.txt")

    val input = infile.readText()

    val regex = Regex("""mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))""")

    var sum = 0
    var multiplicationEnabled = true
    regex.findAll(input).forEach { match ->
        when (match.value) {
            "do()" -> multiplicationEnabled = true
            "don't()" -> multiplicationEnabled = false
            else -> if (multiplicationEnabled) {
                sum += match.groupValues.slice(1..2).map { it.toInt() }.reduce { a, b -> a * b }
            }
        }
    }

    println("Sum of multiples: $sum")
}