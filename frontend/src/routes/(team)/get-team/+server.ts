import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ cookies }) => {
	// Define the cookie value and options
	const teamId = cookies.get('teamId');

	return new Response(teamId);
};
