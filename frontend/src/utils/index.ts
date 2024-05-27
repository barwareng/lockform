import { mediaQuery } from 'svelte-legos';
import Cookies from 'js-cookie';
import { tick } from 'svelte';

export const setTeamCookie = async (teamId: string) => {
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
