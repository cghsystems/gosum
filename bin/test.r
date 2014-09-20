#!/usr/bin/env Rscript
stats <- read.csv("stats", header=TRUE)

meanTime <- mean(stats$time)
medianTime <- median(stats$time)
minTime <- min(stats$time)
maxTime <- max(stats$time)

paste("Minimum execution time:", minTime, "ms", sep=" ")
paste("Maximum execution time:", maxTime, "ms", sep=" ")
paste("Mean Average execution time:", meanTime, "ms", sep=" ")
paste("Median execution time:", medianTime, "ms", sep=" ")

#Generate graph
plot (stats)
title("Execution time")


