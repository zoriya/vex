import { fail } from '@sveltejs/kit';
import { login, register } from '$lib/server/api.js';


export const actions = {
	connect: async ({ cookies, request }) => {
		const data = await request.formData();
		const email = data.get('email');
		const password = data.get('password');
		const token = await login(email as any, password as any);
		return {
			token
		};
	},
	register: async ({ cookies, request }) => {
		const data = await request.formData();
		const email = data.get('email');
		const username = data.get('username');
		const password = data.get('password');
		const passwordRepeat = data.get('password-repeat');
		console.log("register", email, username, password, passwordRepeat);
		if (password !== passwordRepeat) {
			return fail(400, {
				data: {
					email
				},
				errors: {
					password: 'Passwords do not match'
				}
			})
		}
		const token = await register(email as any, username as any, password as any);
		return {
			token
		};
	},
}