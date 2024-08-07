import { client } from '$lib/api/Client';
import type { IContactList } from '$utils/interfaces/contacts.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	let loadingContacts = true;
	let contacts: Partial<IContactList>[] = [];
	try {
		contacts = await client.contacts.getAll();
		loadingContacts = false;
		return { loadingContacts, contacts };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
