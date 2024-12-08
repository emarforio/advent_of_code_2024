package aoc_02

import java.io.File

fun main() {
    val infile = File("src/main/kotlin/aoc_02/in.txt")


    // Trying out mapping instead of iterating more traditionally to ty it out
    infile.useLines { lines ->
        val safeCount = lines.map { line ->
            val levels = line.split(" ").map { it.toInt() }

            // Try all combinations of removing a value from the list
            for (i in 0..levels.lastIndex) {
                val oneShorterLevels = mutableListOf(*levels.toTypedArray())
                oneShorterLevels.removeAt(i)

                if (checkIfSafe(oneShorterLevels)) {
                    return@map 1
                }
            }

            return@map 0


        }.sum()

        println("The total safe reports are $safeCount")
    }
}

/**
 * Naive solution iterating over all combinations of one removed. Ignoring nicer solutions comparing the diffs
 */
fun checkIfSafe(levels: List<Int>): Boolean {
    val diffs = levels.zipWithNext { a, b -> a - b }

    return diffs.all { diff -> diff in 1..3 } ||
            diffs.all { diff -> diff in -3..-1 }
}