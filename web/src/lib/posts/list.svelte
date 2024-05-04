<script lang="ts">
	import type { Post } from "$lib/types";
	import Card from "./card.svelte";
	export let posts: Post[] = [];
</script>

<div>
	<h1>Posts</h1>

	<ul>
		{#each posts as post}
			<li>
				<Card {post} />
			</li>
			<hr>
		{/each}
	</ul>
</div>

<style>
	ul {
		list-style: none;
		padding: 0;
		display: flex;
		flex-direction: column;
	}
</style>
