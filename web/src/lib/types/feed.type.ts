import type { User } from "./user.type";

export type Feed = {
	id: string;
	title: string;
	url: string;
	faviconUrl: string;
	tags: string[];
	submitterId: string;
	addedDate: Date;
	submitter?: User;
}