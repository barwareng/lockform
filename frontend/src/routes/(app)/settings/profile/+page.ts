import { client } from '$lib/api/Client';
import type { IUser } from '$utils/interfaces/users.interfaces';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	try {
		return {
			loadingProfile: false,
			profile: (await client.users.getProfile()) as Partial<IUser>
		};
	} catch (error) {
		toastError(error);
		return { loadingProfile: false };
	}
}) satisfies PageLoad;
