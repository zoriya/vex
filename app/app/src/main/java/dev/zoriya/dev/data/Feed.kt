package dev.zoriya.dev.data

data class Feed(
	val id: String,
	val title: String,
	val url: String,
	val faviconUrl: String,
	val tags: List<Tag>
)
