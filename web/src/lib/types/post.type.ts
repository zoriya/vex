import type { Feed } from "./feed.type";

export type Post = {
	id: string;
	title: string;
	content: string;
	link: string;
	date: Date;

	authors?: string[];
	isRead: boolean;
	isBookmarked: boolean;
	isIgnored: boolean;
	isReadLater: boolean;
	feed: Feed;
}