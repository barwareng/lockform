import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request }) => {
	// Define the cookie value and options
	const { teamId } = await request.json();
	const expires = new Date(Date.now() + 60 * 60 * 24 * 7 * 1000); // 7 days from now
	// Set the cookie header
	const headers = new Headers();
	headers.append(
		'Set-Cookie',
		`teamId=${teamId}; Path=/; HttpOnly; Secure; SameSite=Lax; Expires=${expires.toUTCString()}`
	);

	return new Response('Cookie is set', {
		status: 200,
		headers
	});
};
