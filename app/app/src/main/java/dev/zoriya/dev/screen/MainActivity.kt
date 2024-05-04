package dev.zoriya.dev.screen

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults.topAppBarColors
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import dev.zoriya.dev.component.MainFeed
import dev.zoriya.dev.data.Feed
import dev.zoriya.dev.data.Post
import dev.zoriya.dev.data.Tag
import dev.zoriya.dev.ui.theme.VexTheme
import java.util.Date

class MainActivity : ComponentActivity() {
	override fun onCreate(savedInstanceState: Bundle?) {
		super.onCreate(savedInstanceState)
		enableEdgeToEdge()
		setContent {
			CreateView()
		}
	}

	@OptIn(ExperimentalMaterial3Api::class)
	@Composable
	fun CreateView(){
		VexTheme {
			Scaffold(
				modifier = Modifier.fillMaxSize(),
				topBar = {
					TopAppBar(
						title = {
							Column {
								Text ("Vex")
								Text ("Rex")
							}
						},

						colors = topAppBarColors(
							containerColor = MaterialTheme.colorScheme.primaryContainer,
							titleContentColor = MaterialTheme.colorScheme.primary,
						)
					)
				},
			) { innerPadding ->
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
					}, modifier = Modifier.padding(innerPadding)
				)

			}
		}
	}

	@Composable
	@Preview
	fun ScrollableColumnOfCardsPreview(){
		CreateView()
	}

}


