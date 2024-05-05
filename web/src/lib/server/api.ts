import { env } from '$env/dynamic/private';
import type { Post } from '$lib/types';

export async function login(email: string, password: string) {
	const r = await fetch(env.API_URL + '/login', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ email, password })
	});
	const j = await r.json();
	return j.token as string;
}

export async function register(email: string, username: string, password: string) {
	const r = await fetch(env.API_URL + '/register', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ email, password, name: username })
	});
	const j = await r.json();
	return j.token as string;
}

export async function getPosts(token?: string) {
	const opts = {
		headers: {
			Authorization: `Bearer ${token}`
		}
	}
	const r = await fetch(env.API_URL + '/entries', token ? opts : undefined);
	const j = await r.json();
	return j.posts as Post[];
}