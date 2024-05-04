<script lang="ts">
	import type { Post } from "$lib/types";
	import Time, { dayjs } from "svelte-time";
	import TagList from "./tag-list.svelte";
	export let post: Post;
</script>

<article>
	<div class="origin">
		<img src={post.feed.faviconUrl} alt={post.feed.title} />
		<span>{post.feed.title}</span> |
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
	<footer>
		<TagList tags={post.feed.tags} />
	</footer>
</article>

<style>
	article {
		border-radius: 5px;
		/* box-shadow: 0px 0px 12px 9px rgba(0, 0, 0, 0.28); */
		border: 1px solid #e0e0e0;
		padding: 0.5rem;
	}

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
	}

	p {
		margin: 0.5rem;
	}
</style>
