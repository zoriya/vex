import type { Post } from "$lib/types";

const data: Post[] = [
	{
		id: "1",
		title: "The first post",
		content: "This is the",
		link: "https://example.com",
		date: new Date(),
		isRead: false,
		isBookmarked: false,
		isIgnored: false,
		isReadLater: false,
		author: "John Doe",
		feed: {
			id: "1",
			title: "The first feed",
			url: "https://example.com/feed",
			faviconUrl: "https://kit.svelte.dev/favicon.png",
			tags: ["tag1", "tag2"],
		},
	},
	{
		id: "2",
		title: "Wine 9.8 Fixes Nearly 20 Year Old Bug For Installing Microsoft Office 97",
		content: "Wine 9.8 is out today as the newest bi-weekly development release of this open-source software for enjoying Windows games/applications on Linux / Chrome OS, macOS, and other platforms...",
		link: "https://example.com",
		date: new Date(),
		isRead: false,
		isBookmarked: true,
		isIgnored: false,
		isReadLater: false,
		feed: {
			id: "2",
			title: "Phoronix",
			url: "https://www.phoronix.com",
			faviconUrl: "https://www.phoronix.com/favicon.ico",
			// put 50 tags in the array linked with linux world
			tags: [
				"linux", "kernel", "gnu", "gnu/linux", "gnu+linux", "gnu linux",
				"wine", "winehq", "wine-staging", "wine-devel", "winehq-devel", "winehq-staging",
				"mesa", "mesa3d", "mesa 3d", "mesa-3d", "mesa3d-devel", "mesa3d-staging",
				"NVK"
			],
		},
	}
];

export function load() {
	return { data };
}

