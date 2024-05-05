<script>
	import "../app.pcss";

	import { BottomNav, BottomNavItem } from "flowbite-svelte";
	import { HomeSolid, UserCircleSolid } from "flowbite-svelte-icons";
</script>

<slot></slot>

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
