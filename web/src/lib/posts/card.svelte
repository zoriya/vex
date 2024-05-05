<script lang="ts">
	import type { Post } from "$lib/types";
	import Time, { dayjs } from "svelte-time";
	import QuillPenLine from "virtual:icons/mingcute/quill-pen-line";
	export let post: Post;
</script>

<a
	href={post.link}
	target="_blank"
	on:click={() => {
		console.log("clicked");
	}}
>
	<article>
		<div class="origin">
			<img src={post.feed.faviconUrl} alt={post.feed.name} />
			<span>{post.feed.name}</span> |
			{#if post.authors?.length}
				{#each post.authors as author}
					<div class="author">
						<QuillPenLine />
						<span>{author}</span>
					</div>
					|
				{/each}
			{/if}
			<Time
				timestamp={post.date}
				relative={true}
				title={dayjs(post.date).locale("en").toString()}
			/>
		</div>
		<content>
			<h2>{post.title}</h2>
			<p>{post.content}</p>
		</content>
		<!-- <footer>
		<TagList tags={post.feed.tags} />
	</footer> -->
	</article>
</a>

<style>
	.origin {
		display: flex;
		align-items: center;
		border-top-left-radius: 5px;
		border-top-right-radius: 5px;
		gap: 6px;
	}

	img {
		width: 20px;
		height: 20px;
		margin-right: 0.5rem;
	}

	h2 {
		margin: 0.5rem;
		font-size: 1.3rem;
		font-weight: 600;
	}

	p {
		margin: 0.5rem;
	}

	.author {
		display: flex;
		align-items: center;
		gap: 3px;
	}
</style>
