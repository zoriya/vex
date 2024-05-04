package dev.zoriya.dev.data

import java.io.Serializable
import java.util.Date

data class Post(
	val id: String,
	val title: String,
	val content: String,
	val link: String,
	val date: Date,
	val author: String,
	val isRead: Boolean,
	val isBookmarked: Boolean,
	val isIgnored: Boolean,
	val isReadLater: Boolean,
	val feed: Feed
) : Serializable
