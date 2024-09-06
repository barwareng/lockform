import { mediaQuery } from 'svelte-legos';
import Cookies from 'js-cookie';
import { tick } from 'svelte';

// Pass a string to set the team cookie, pass nothing to delete it.
export const setTeamCookie = async (teamId?: string) => {
	await fetch('/set-team', {
		method: 'POST',
		body: JSON.stringify({ teamId })
	});
};
export const getTeamCookie = async (): Promise<string> => {
	const res = await fetch('/get-team');
	return await res.text();
};
export const deleteTeamCookie = () => {
	Cookies.remove('teamId')!;
};
export const isMobile = mediaQuery('(max-width: 768px)');
export const closeAndRefocusTrigger = (triggerId: string): boolean => {
	tick().then(() => document.getElementById(triggerId)?.focus());
	return false;
};
// TODO modify this function to fill in only keys required by the
export const parseSearchParams = <T>(searchParams: URLSearchParams): Partial<T> => {
	const params: Partial<T> = {};
	searchParams.forEach((value, key) => {
		(params as any)[key] = value;
	});
	return params;
};
