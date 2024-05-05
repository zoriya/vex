<script lang="ts">
	export let data;
	import {
		BottomNav,
		BottomNavItem,
		Skeleton,
		ImagePlaceholder,
	} from "flowbite-svelte";
	import List from "$lib/posts/list.svelte";
	import {
		HomeSolid,
		WalletSolid,
		AdjustmentsVerticalOutline,
		UserCircleSolid,
	} from "flowbite-svelte-icons";
</script>

<main>
	<List posts={data.posts} />
</main>

<style>
	main {
		display: flex;
		justify-content: center;
	}
</style>
