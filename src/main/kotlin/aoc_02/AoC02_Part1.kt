package aoc_02

import java.io.File

fun main() {
    val infile = File("src/main/kotlin/aoc_02/in.txt")


    // Trying out mapping instead of iterating more traditionally to ty it out
    infile.useLines { lines ->
        val safeCount = lines.map { line ->
            val levels = line.split(" ").map { it.toInt() }

            val diffs = levels.zipWithNext { a, b -> a - b }

            if (diffs.all { diff -> diff in 1..3 } ||
                diffs.all { diff -> diff in -3..-1 }) {
                return@map 1
            } else {
                return@map 0
            }
        }.sum()

        println("The total safe reports are $safeCount")
    }


}