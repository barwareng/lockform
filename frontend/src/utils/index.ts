import { mediaQuery } from 'svelte-legos';
import Cookies from 'js-cookie';
import { tick } from 'svelte';

export const setTeamCookie = async (teamId: string) => {
	// Cookies.set('teamId', teamId, {
	// 	sameSite: 'Lax',
	// 	secure: true,
	// 	expires: 60 * 60 * 24 * 7,
	// });
	await fetch('/', {
		method: 'POST',
		body: JSON.stringify({ teamId })
	});
};
export const getTeamCookie = (): string => {
	// const teamId = getCookie('teamId')!;
	const teamId = Cookies.get('teamId')!;
	console.log('Getting team cookie', teamId);
	return teamId;
};
export const deleteTeamCookie = () => {
	Cookies.remove('teamId')!;
};
export const isMobile = mediaQuery('(max-width: 768px)');
export const closeAndRefocusTrigger = (triggerId: string): boolean => {
	tick().then(() => document.getElementById(triggerId)?.focus());
	return false;
};

const getCookie = (name) => {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${name}=`);
	if (parts.length === 2) return parts.pop().split(';').shift();
};
