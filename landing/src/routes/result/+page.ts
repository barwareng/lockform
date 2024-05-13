import { client } from '$lib/api/Client';
import type { Iverification } from '$utils/interfaces/teams.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './[searchPhrase]/$types';
import { parsePhoneNumber, isPossiblePhoneNumber } from 'libphonenumber-js';
export const load = (async ({ url }) => {
	let verifying = false;
	let verification: Iverification;
	let searchPhrase = url.searchParams.get('search')!;
	try {
		verifying = true;
		if (isPossiblePhoneNumber(searchPhrase)) {
			searchPhrase = parsePhoneNumber(searchPhrase)?.formatInternational().replaceAll(' ', '');
		}
		verification = await client.verifications.verify({ searchPhrase });

		verifying = false;
		return { verifying, verification };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
