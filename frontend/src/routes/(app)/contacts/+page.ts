import { client } from '$lib/api/Client';
import type { IContactList } from '$utils/interfaces/contacts.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	try {
		return {
			loading: false,
			contacts: (await client.contacts.getAll()) as Partial<IContactList[]>
		};
	} catch (error) {
		toastError(error);
		return { loading: false };
	}
}) satisfies PageLoad;
