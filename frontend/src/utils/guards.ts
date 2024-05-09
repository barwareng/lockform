import { browser } from '$app/environment';
import type { ROLE_VALUES } from './interfaces/roles.interface';

export const requireRoles = (roles: ROLE_VALUES[]): boolean => {
	const userRoles: ROLE_VALUES[] = getRolesLocal();
	return userRoles.length > 0 && roles.some((role) => userRoles?.includes(role));
};

export const setRolesLocal = (roles: ROLE_VALUES[]) => {
	browser && localStorage.setItem('roles', JSON.stringify(roles));
};
export const getRolesLocal = (): ROLE_VALUES[] => {
	return browser && localStorage.getItem('roles') && JSON.parse(localStorage.getItem('roles')!);
};
