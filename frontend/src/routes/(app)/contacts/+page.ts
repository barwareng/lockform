import { client } from '$lib/api/Client';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async ({ url }) => {
	try {
		// Parse url searchparams into object
		const { page, pageSize } = Object.fromEntries(url.searchParams);
		return {
			loading: false,
			contacts: await client.contacts.getAll({
				page: Number(page),
				pageSize: Number(pageSize)
			})
		};
	} catch (error) {
		toastError(error);
		return { loading: false };
	}
}) satisfies PageLoad;
