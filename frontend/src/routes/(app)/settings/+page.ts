import { client } from '$lib/api/Client';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	try {
		let loadingMembers = true;
		const members = await client.members.getAll();
		loadingMembers = false;
		return { loadingMembers, members };
	} catch (error) {
		toastError(error);
		// return { loadingMembers: false };
	}
}) satisfies PageLoad;
