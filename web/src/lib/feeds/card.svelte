<script lang="ts">
	import type { Feed } from "$lib/types";
	import TagList from "$lib/posts/tag-list.svelte";
	import User from "virtual:icons/mingcute/user-5-line";
	export let feed: Feed;

	let viewTags = false;
	let hover = false;
</script>

<article>
	<div>
		<img src={feed.faviconUrl} alt={feed.name} />
		<h3>{feed.name}</h3>
	</div>
	<div>
		<span>{feed.link}</span>
	</div>
	{#if feed.submitter}
		<div>
			Added by
			<User />
			<span title={feed.submitter.email}>{feed.submitter.name}</span>
		</div>
	{/if}
	{#if viewTags}
		<TagList tags={feed.tags} />
	{:else}
		<button on:click={() => (viewTags = true)}>View tags</button>
	{/if}
</article>

<style>
	div {
		display: flex;
		align-items: center;
		border-top-left-radius: 5px;
		border-top-right-radius: 5px;
		gap: 6px;
	}
	article {
		border-radius: 10px;
		padding: 0.7rem;
	}

	img {
		width: 40px;
		height: 40px;
		margin-right: 0.5rem;
	}
</style>
