#!/usr/bin/env Rscript
stats <- read.csv("stats", header=TRUE)

meanTime <- mean(stats$time)

paste("Mean Average execution time:", meanTime, "ms", sep=" ")

plot (stats)
title("Execution time")


