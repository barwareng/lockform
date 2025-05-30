import { browser } from '$app/environment';
import { getTeamCookie, setTeamCookie } from '$utils';
import { setRolesLocal } from '$utils/guards';
import type { ROLE_VALUES } from '$utils/interfaces/roles.interface';
import type { ITeam } from '$utils/interfaces/teams.interface';
import { onboardingAllowedRoutes } from '$utils/routing';
import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';
import Session from 'supertokens-web-js/recipe/session';
export const load = (async ({ url }) => {
	let teams: Partial<ITeam>[] = [];
	let userId = '';
	let isOnboarded = false;
	let selectedTeam: Partial<ITeam> = {};
	if (browser && (await Session.doesSessionExist())) {
		const accessTokenPayload = await Session.getAccessTokenPayloadSecurely();
		const roleClaims: string[] = accessTokenPayload?.['st-role']?.v;
		userId = accessTokenPayload.userId;
		teams = accessTokenPayload.teams;
		isOnboarded = accessTokenPayload.isOnboarded;
		const teamId = await getTeamCookie();
		if (teams?.length) {
			if (!teamId) {
				await setTeamCookie(teams[0].id!);
				selectedTeam = teams[0];
			} else {
				teams.forEach((team) => {
					if (teamId == team.id) {
						selectedTeam = team;
					}
				});
			}
		}
		// get roles for this current team and strip the team ID
		const roles = roleClaims
			?.filter((c) => c.includes(teamId))
			?.map((claim) => claim.replace(`${teamId}_`, '')) as ROLE_VALUES[];
		setRolesLocal(roles);
		if (!isOnboarded && !onboardingAllowedRoutes.has(url.pathname))
			throw redirect(302, '/settings/profile');
	}
	return { teams, userId, selectedTeam };
}) satisfies LayoutLoad;
