package aoc_01

import java.io.File

fun main() {
    val infile = File("src/main/kotlin/aoc_01/in.txt")

    val leftNumbers = mutableListOf<Int>()
    val rightNumbers = mutableMapOf<Int, Int>()

    infile.forEachLine { line ->
        val (left, right) = line.split("   ").map { it.toInt() }
        leftNumbers.addLast(left)
        rightNumbers.merge(right, 1) { oldSum, newSum -> oldSum + newSum }
    }

    val sum = leftNumbers.sumOf { number -> number * rightNumbers.getOrDefault(number, 0) }

    println("The total distance is $sum")
}