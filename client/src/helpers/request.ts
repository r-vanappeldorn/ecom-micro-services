import type { FormError } from '../types/form-error';

export const sendRequest = async (url: string, method: string, body?: object) => {
	const res = await fetch(url, {
		method,
		body: JSON.stringify(body),
		headers: {
			'Content-Type': 'application/json'
		}
	});
	const data = await res.json();

	if (!data.errors) {
		return [];
	}

	return data.errors as FormError[];
};
