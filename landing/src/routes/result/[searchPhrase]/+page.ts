import { client } from '$lib/api/Client';
import type { Iverification } from '$utils/interfaces/teams.interface';
import { toastError } from '$utils/toasts';
import type { PageLoad } from './$types';

export const load = (async ({ params }) => {
	let verifying = false;
	let verification: Iverification;
	const { searchPhrase } = params;
	try {
		verifying = true;
		verification = await client.verifications.verify({ searchPhrase });

		verifying = false;
		return { verifying, verification };
	} catch (error) {
		toastError(error);
	}
}) satisfies PageLoad;
