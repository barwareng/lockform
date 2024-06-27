import { client } from '$lib/api/Client';
import type { IContact } from '$utils/interfaces/contacts.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	let laodingContacts = true;
	let contacts: Partial<IContact>[] = [];
	try {
		contacts = await client.contacts.getAll();
		laodingContacts = false;
		return { laodingContacts, contacts };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
