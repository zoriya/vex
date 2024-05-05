package dev.zoriya.dev.component

import android.text.format.DateUtils
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Card
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import dev.zoriya.dev.data.Feed
import dev.zoriya.dev.data.Post
import dev.zoriya.dev.data.Tag
import dev.zoriya.dev.ui.theme.VexTheme
import dev.zoriya.dev.utils.getElapsedTime
import java.lang.String.format
import java.text.DateFormat
import java.util.Date
import java.util.Locale


@Composable
fun MainFeed(posts: List<Post>, modifier: Modifier) {
	LazyColumn(
		modifier = modifier.fillMaxSize()
	) {
		items(posts) { post ->
			Box(
				modifier = Modifier
					.fillMaxWidth().padding(5.dp)
					.clip(MaterialTheme.shapes.small)
					.background(
						MaterialTheme.colorScheme.secondaryContainer
					)

			) {
				Column(
					modifier = modifier.fillMaxWidth().padding(10.dp)
				) {
					Text(
						text = post.author + " | " +
							getElapsedTime(Date(), post.date).seconds + "s ago"
					)
					Text(
						text = post.title,
						style = TextStyle(fontSize = 18.sp),
						fontWeight = FontWeight.Bold
					)
					Text(
						text = post.content, style = TextStyle(fontSize = 16.sp)
					)
					Row(
						verticalAlignment = Alignment.CenterVertically,
						horizontalArrangement = Arrangement.Start,
						modifier = modifier
					) {
						for (tag in post.feed.tags) {
							Box(
								modifier = Modifier
									.clip(MaterialTheme.shapes.small)
									.background(
										MaterialTheme.colorScheme.secondaryContainer
									)
							) {
								Text(tag.name, color = MaterialTheme.colorScheme.secondary, modifier = Modifier.padding(1.dp))
								Spacer(modifier = Modifier.padding(1.dp))
							}
						}
					}
				}
				Spacer(modifier = Modifier)
			}
		}
	}
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
@Preview
fun CreateView() {
	VexTheme {
		MainFeed(
			posts = List(10) { i ->
				Post(
					i.toString(),
					"titre",
					"content",
					"link",
					Date(),
					"author",
					isRead = false,
					isBookmarked = false,
					isIgnored = false,
					isReadLater = false,
					Feed(
						i.toString(), "title", "", "", listOf(Tag("Tageul"), Tag("Tageul1"))
					)
				)
			}, modifier = Modifier.fillMaxWidth()
		)
	}
}
