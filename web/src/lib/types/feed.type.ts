import type { User } from "./user.type";

export type Feed = {
	id: string;
	name: string;
	link: string;
	faviconUrl: string;
	tags: string[];
	submitterId: string;
	addedDate: Date;
	submitter?: User;
}