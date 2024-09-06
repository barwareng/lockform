import { VITE_SUPERTOKENS_COOKIE_DOMAIN } from '$lib/env';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request }) => {
	// Define the cookie value and options
	const { teamId } = await request.json();
	let expires: Date;
	if (!teamId) {
		expires = new Date(0); // Expire the cookie
	} else {
		expires = new Date(Date.now() + 60 * 60 * 24 * 7 * 1000); // 7 days from now
	}

	const cookie = `teamId=${teamId}; Path=/; HttpOnly; Secure; SameSite=Lax; Expires=${expires.toUTCString()}; Domain=${VITE_SUPERTOKENS_COOKIE_DOMAIN}`;
	const res = new Response(`team cookie is ${teamId ? 'set' : 'deleted'}`, {
		status: 200
	});
	res.headers.append('set-cookie', cookie);
	return res;
};
