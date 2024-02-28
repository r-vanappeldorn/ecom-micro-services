import type { Cookies } from '@sveltejs/kit';
import type { User } from '../types/user';

export const getCurrentUser = async (cookies: Cookies): Promise<User | null> => {
	try {
		const session = cookies.get('session');
		if (!session) return null;

		const res = await fetch('http://auth-srv:3000/api/users/currentuser', {
			method: 'get',
			headers: {
				'Content-Type': 'application/json',
				cookie: `session=${session};`
			}
		});

		const data = await res.json();

		return data.currentUser;
	} catch (err) {
		console.log(err);
		return null;
	}
};
