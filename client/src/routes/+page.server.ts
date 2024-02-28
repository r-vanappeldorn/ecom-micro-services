import { getCurrentUser } from '../helpers/get-current-user';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	const user = await getCurrentUser(cookies);

	return {
		user
	};
}) satisfies PageServerLoad;
