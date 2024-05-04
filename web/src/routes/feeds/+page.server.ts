import type { Feed } from "$lib/types";


const feeds: Feed[] = [
	{
		id: "1",
		title: "The first feed",
		url: "https://example.com/feed",
		faviconUrl: "https://kit.svelte.dev/favicon.png",
		tags: ["tag1", "tag2"],
		submitterId: "1",
		addedDate: new Date(),
	},
	{
		id: "2",
		title: "Phoronix",
		url: "https://www.phoronix.com",
		faviconUrl: "https://www.phoronix.com/favicon.ico",
		tags: ["linux", "kernel", "gnu", "gnu/linux", "gnu+linux", "gnu linux"],
		submitterId: "1",
		addedDate: new Date(),
	},
	{
		id: "3",
		title: "LWN.net",
		url: "https://lwn.net",
		faviconUrl: "https://lwn.net/favicon.ico",
		tags: ["linux", "kernel", "gnu", "gnu/linux", "gnu+linux", "gnu linux"],
		submitterId: "1",
		addedDate: new Date(),
	},
	{
		id: "4",
		title: "Reddit",
		url: "https://www.reddit.com",
		faviconUrl: "https://www.reddit.com/favicon.ico",
		tags: ["social", "news", "discussion"],
		submitterId: "1",
		addedDate: new Date(),
		submitter: {
			id: "1",
			name: "John Doe",
			email: "john.doe@gmail.com"
		}
	},
	{
		id: "5",
		title: "Hacker News",
		url: "https://news.ycombinator.com",
		faviconUrl: "https://news.ycombinator.com/favicon.ico",
		tags: ["news", "discussion"],
		submitterId: "1",
		addedDate: new Date(),
	}
];

export function load() {
	return {
		feeds
	};
}

