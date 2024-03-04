import { getCurrentUser } from '../helpers/get-current-user';
import type { PageServerLoad } from './$types';
import type { Ticket } from '../types/tickets';

export const load = (async ({ cookies }) => {
	const user = await getCurrentUser(cookies);
	let tickets: Ticket[] = [];

	try {
		const res = await fetch('http://tickets-srv:8080/api/tickets', {
			method: 'get',
			headers: {
				'Content-Type': 'application/json',
				Host: 'ticketing.io'
			}
		});
		tickets = ((await res.json()) as { tickets: Ticket[] }).tickets;
	} catch (e) {
		console.log(e);
	}

	return {
		user,
		tickets
	};
}) satisfies PageServerLoad;
