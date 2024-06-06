import { client } from '$lib/api/Client';
import type { ITrustedContact } from '$utils/interfaces/trusted-contacts.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async () => {
	let loadingTrustedContacts = true;
	let trustedContacts: Partial<ITrustedContact>[] = [];
	try {
		trustedContacts = await client.trustedContacts.getAll();
		loadingTrustedContacts = false;
		return { loadingTrustedContacts, trustedContacts };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
