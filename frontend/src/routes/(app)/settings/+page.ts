import { client } from '$lib/api/Client';
import type { IMember } from '$utils/interfaces/users.interfaces';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	let loadingMembers = true;
	let members: IMember[] = [];
	try {
		members = await client.members.getAll();
		loadingMembers = false;
		return { loadingMembers, members };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
