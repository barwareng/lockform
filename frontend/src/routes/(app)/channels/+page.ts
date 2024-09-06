import { client } from '$lib/api/Client';
import type { IChannel } from '$utils/interfaces/channels.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	try {
		return {
			loading: false,
			channels: (await client.channels.getAll()) as Partial<IChannel[]>
		};
	} catch (error) {
		toastError(error);
		return {
			loading: false
		};
	}
}) satisfies PageLoad;
