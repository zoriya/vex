package dev.zoriya.dev.utils

import java.time.Duration
import java.util.Date
import kotlin.math.abs


fun getElapsedTime(startDate: Date, endDate: Date): Duration {
	return Duration.between(endDate.toInstant(),startDate.toInstant())
}
