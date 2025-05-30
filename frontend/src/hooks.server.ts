import { client } from '$lib/api/Client';
import { VITE_API_BASE_URL } from '$lib/env';
import { isPublicRoute, onboardingAllowedRoutes } from '$utils/routing';
import { redirect, type Handle } from '@sveltejs/kit';
import * as jose from 'jose';

export const handle = (async ({ event, resolve }) => {
	client.setFetch(event.fetch);
	if (isPublicRoute(event.url.pathname)) {
		const response = await resolve(event);
		return response;
	}

	const jwt = event.cookies.get('sAccessToken');
	// If there's no jwt token at all, either the user signed out (or never logged in)
	// OR they've never been to the app in the first place.
	// Since the latter is far more likely right now, we just redirect them to signup rather than signin.
	// This avoids the infinite redirect issue in (https://supertokens.com/docs/thirdparty/common-customizations/sessions/ssr#why-do-we-trigger-the-refresh-session-flow-instead-of-redirecting-the-user-to-the-login-page-directly)
	// because we have separate handling for a present but expired/invalid jwt token below
	if (!jwt) {
		event.cookies.delete('teamId', { path: '/' }); // Delete team id, so the next logged in user does not access it.
		if (!isPublicRoute(event.url.pathname)) {
			throw redirect(302, '/signin');
		} else {
			const response = await resolve(event);
			return response;
		}
	}
	const JWKS = jose.createRemoteJWKSet(new URL(`${VITE_API_BASE_URL}/auth/jwt/jwks.json`));
	// @ts-ignore
	const { payload } = await jose.jwtVerify(jwt, JWKS).catch(async (err) => {
		if (!isPublicRoute(event.url.pathname)) {
			const redirectBack =
				event.url.href.replace(event.url.origin, '') !== '/'
					? `?redirectBack=${event.url.href.replace(event.url.origin, '')}`
					: '';
			throw redirect(302, `/refresh-session${redirectBack}`);
		}
	});
	if (payload && typeof payload === 'object') {
		// Prevent access until email verification is complete
		const isEmailVerified = (payload as any)['st-ev'].v;
		if (!isEmailVerified) {
			throw redirect(302, '/verify-email/request-verification');
		}

		// Handle onboarding
		const isOnboarded = !!(payload as any)['isOnboarded'];
		if (!isOnboarded && !onboardingAllowedRoutes.has(event.url.pathname)) {
			throw redirect(302, '/settings/profile');
		}

		if (
			!payload.teams?.length &&
			event.url.pathname !== '/team' &&
			!onboardingAllowedRoutes.has(event.url.pathname)
		) {
			throw redirect(302, '/team');
		}
	}
	const response = await resolve(event);
	return response;
}) satisfies Handle;

/** @type {import('@sveltejs/kit').HandleFetch} */
export async function handleFetch({ event, request, fetch }) {
	if (request.url.startsWith(VITE_API_BASE_URL)) {
		const cookie = event.request.headers.get('cookie');
		if (cookie) {
			request.headers.set('cookie', cookie);
		}
	}
	return fetch(request);
}
