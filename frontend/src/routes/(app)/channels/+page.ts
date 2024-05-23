import { client } from '$lib/api/Client';
import type { IChannel } from '$utils/interfaces/channels.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	let loadingChannels = true;
	let channels: Partial<IChannel>[] = [];
	try {
		channels = await client.channels.getAll();
		loadingChannels = false;
		return { loadingChannels, channels };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
