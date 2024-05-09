import { browser } from '$app/environment';
import Cookies from 'js-cookie';
import { tick } from 'svelte';

export const setTeamCookie = (teamID: string) => {
	Cookies.set('teamId', teamID, { sameSite: 'Lax', secure: true });
};
export const getTeamCookie = (): string => {
	return Cookies.get('teamId')!;
};
export const deleteTeamCookie = () => {
	Cookies.remove('teamId')!;
};

export const isMobileDevice = (): boolean => browser && window.innerWidth <= 640;

export const closeAndRefocusTrigger = (triggerId: string): boolean => {
	tick().then(() => document.getElementById(triggerId)?.focus());
	return false;
};
