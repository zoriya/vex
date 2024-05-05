<script lang="ts">
	import type { Post } from "$lib/types";
	import Card from "./card.svelte";
	export let posts: Post[] = [];
</script>

<ul>
	{#each posts as post, index}
		<li>
			<Card {post} />
		</li>
		{#if index !== posts.length - 1}
			<hr />
		{/if}
	{/each}
</ul>

<style>
	ul {
		list-style: none;
		padding: 0;
		display: flex;
		gap: 1rem;
		flex-direction: column;
	}
</style>
