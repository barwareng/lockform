import { client } from '$lib/api/Client';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	try {
		return { members: await client.members.getAll(), loadingMembers: false };
	} catch (error) {
		toastError(error);
		return { loadingMembers: false };
	}
}) satisfies PageLoad;
