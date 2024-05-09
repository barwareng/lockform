import { client } from '$lib/api/Client';
import type { IUser } from '$utils/interfaces/users.interfaces';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	let loadingProfile = true;
	let profile: Partial<IUser> = {};
	try {
		loadingProfile = true;
		profile = await client.users.getProfile();
		loadingProfile = false;
		return { loadingProfile, profile };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
