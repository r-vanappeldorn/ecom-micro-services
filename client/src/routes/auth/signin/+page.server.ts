import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	const session = cookies.get('session');

	if (session) {
		throw redirect(302, '/');
	}

	return {};
}) satisfies PageServerLoad;
