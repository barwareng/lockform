import { mediaQuery } from 'svelte-legos';
import Cookies from 'js-cookie';
import { tick } from 'svelte';

export const setTeamCookie = (teamID: string) => {
	Cookies.set('teamId', teamID, {
		sameSite: 'Lax',
		secure: true,
		expires: 60 * 60 * 24 * 7,
		httpOnly: true
		// httpOnly: process.env.NODE_ENV === 'production'
	});
};
export const getTeamCookie = (): string => {
	return Cookies.get('teamId')!;
};
export const deleteTeamCookie = () => {
	Cookies.remove('teamId')!;
};
export const isMobile = mediaQuery('(max-width: 768px)');
export const closeAndRefocusTrigger = (triggerId: string): boolean => {
	tick().then(() => document.getElementById(triggerId)?.focus());
	return false;
};
