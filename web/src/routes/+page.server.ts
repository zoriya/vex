import type { Post, Feed } from "$lib/types";
import { getPosts } from "$lib/server/api";
import { get } from "svelte/store";

const feeds: Feed[] = [
	{
		id: "1",
		name: "The first feed",
		link: "https://example.com/feed",
		faviconUrl: "https://kit.svelte.dev/favicon.png",
		tags: ["tag1", "tag2"],
		submitterId: "1",
		addedDate: new Date(),
	},
	{
		id: "2",
		name: "Phoronix",
		link: "https://www.phoronix.com",
		faviconUrl: "https://www.phoronix.com/favicon.ico",
		tags: ["linux", "kernel", "gnu", "gnu/linux", "gnu+linux", "gnu linux"],
		submitterId: "1",
		addedDate: new Date(),
	},
	{
		id: "3",
		name: "LWN.net",
		link: "https://lwn.net",
		faviconUrl: "https://lwn.net/favicon.ico",
		tags: ["linux", "kernel", "gnu", "gnu/linux", "gnu+linux", "gnu linux"],
		submitterId: "1",
		addedDate: new Date(),
	},
];

const posts: Post[] = [
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
		authors: ["John Doe"],
		feed: feeds[0],
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
		feed: feeds[1],
	}
];

export async function load() {
	return { 
		posts: await getPosts(),
		feeds
	 };
}

