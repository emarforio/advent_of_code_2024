package aoc_01

import java.io.File
import kotlin.math.absoluteValue

fun main() {
    val infile = File("src/main/kotlin/aoc_01/in.txt")

    val leftNumbers = mutableListOf<Int>()
    val rightNumbers = mutableListOf<Int>()

    infile.forEachLine { line ->
        val (left, right) = line.split("   ").map { it.toInt() }
        leftNumbers.addLast(left)
        rightNumbers.addLast(right)
    }

    leftNumbers.sort()
    rightNumbers.sort()

    val sum = leftNumbers.zip(rightNumbers) { left, right -> (left - right).absoluteValue }.sum()


    println("The total distance is $sum")
}