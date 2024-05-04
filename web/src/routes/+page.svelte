<script>
	export let data;
	import List from "$lib/posts/list.svelte";
</script>

<main>
	<h1>Welcome to SvelteKit</h1>
	<p>
		Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
	</p>

	<List posts={data.data} />
</main>

<style>
	main {
		margin-left: 10%;
	}
</style>
