<script lang="ts">
	export let data;
	import {
		BottomNav,
		BottomNavItem,
		Skeleton,
		ImagePlaceholder,
		Avatar,
		P,
	} from "flowbite-svelte";
	import List from "$lib/posts/list.svelte";
	import { CalendarWeekSolid } from "flowbite-svelte-icons";
	import TimelineItem from "$lib/posts/timeline-item.svelte";
	import { Timeline, TimelineItem as TI } from "flowbite-svelte";
	import {
		HomeSolid,
		WalletSolid,
		AdjustmentsVerticalOutline,
		UserCircleSolid,
	} from "flowbite-svelte-icons";

	$: displayPosts = (data?.posts || []).map((post) => {
		let dateStr = post.feed.name + " - ";
		dateStr += new Date(post.date).toLocaleDateString("en-US", {
			year: "numeric",
			month: "long",
			day: "numeric",
		});

		if (post.authors && post.authors.length > 0) {
			dateStr += ` by ${post.authors.join(", ")}`;
		}
		return {
			...post,
			dateStr,
		};
	});
</script>

<main>
	<Timeline order="vertical" class="max-w-3xl">
		{#each displayPosts as post}
			<TimelineItem title={post.title} date={post.dateStr} post={post}>
				<svelte:fragment slot="icon">
					<span
						class="flex absolute -start-3 justify-center items-center w-6 h-6 bg-primary-200 rounded-full ring-8 ring-white dark:ring-gray-900 dark:bg-primary-900"
					>
						<Avatar
							src={post.feed.faviconUrl}
							alt={post.feed.name}
							class="w-4 h-4 -start-2 -top-2"
						/>
					</span>
				</svelte:fragment>
				<p
					class="mb-4 text-base font-normal text-gray-500 dark:text-gray-400 line-clamp-2"
				>
					{#if post.content}
						{post.content}
					{/if}
				</p>
			</TimelineItem>
		{/each}
	</Timeline>
</main>

<style>
	main {
		display: flex;
		width: 100vw;
		justify-content: center;
	}
</style>
