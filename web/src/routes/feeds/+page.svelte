<script lang="ts">
	import Card from "$lib/feeds/card.svelte";
	export let data;
</script>

<section>
	<h1>All Feeds</h1>
	<h1>Feeds</h1>
	<ul>
		{#each data.feeds as feed}
			<Card {feed} />
		{/each}
	</ul>
</section>

<style>
	ul {
		list-style: none;
		padding: 0;
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
		gap: 10px;
		align-items: stretch;
		justify-items: stretch;
	}
</style>
