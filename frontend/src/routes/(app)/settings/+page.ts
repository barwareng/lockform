import { client } from '$lib/api/Client';
import type { PageLoad } from './$types';

export const load = (async () => {
	let loadingMembers = true;
	const members = await client.members.getAll();
	loadingMembers = false;
	return { loadingMembers, members };
}) satisfies PageLoad;
