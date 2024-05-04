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
	<section>
        <List posts={data.posts} />
    </section>
	<footer>
		<BottomNav position="absolute" classInner="grid-cols-2">
			<BottomNavItem btnName="Posts">
				<HomeSolid
					class="w-6 h-6 mb-1 text-gray-500 dark:text-gray-400 group-hover:text-primary-600 dark:group-hover:text-primary-500"
				/>
			</BottomNavItem>
			<BottomNavItem btnName="Profile">
				<UserCircleSolid
					class="w-6 h-6 mb-1 text-gray-500 dark:text-gray-400 group-hover:text-primary-600 dark:group-hover:text-primary-500"
				/>
			</BottomNavItem>
		</BottomNav>
	</footer>
</main>

<style>
	main {
		display: flex;
	}
	.sidebar {
		display: flex;
		flex-direction: column-reverse;
		width: 400px;
		height: 100vh;
	}
</style>
