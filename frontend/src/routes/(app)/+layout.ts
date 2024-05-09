import { browser } from '$app/environment';
import { getTeamCookie, setTeamCookie } from '$utils';
import { setRolesLocal } from '$utils/guards';
import type { ROLE_VALUES } from '$utils/interfaces/roles.interface';
import type { ITeam } from '$utils/interfaces/teams.interface';
import type { LayoutLoad } from './$types';
import Session from 'supertokens-web-js/recipe/session';
export const load = (async () => {
	let teams: Partial<ITeam>[] = [];
	let userId = '';
	if (browser && (await Session.doesSessionExist())) {
		const accessTokenPayload = await Session.getAccessTokenPayloadSecurely();
		const roleClaims: string[] = accessTokenPayload?.['st-role']?.v;
		userId = accessTokenPayload.userId;
		teams = accessTokenPayload.teams;
		if (!getTeamCookie() && teams.length) {
			setTeamCookie(teams[0].id!);
		}
		const teamId = getTeamCookie();
		// get roles for this current team and strip the team ID
		const roles = roleClaims
			?.filter((c) => c.includes(teamId))
			?.map((claim) => claim.replace(`${teamId}_`, '')) as ROLE_VALUES[];
		setRolesLocal(roles);
	}
	return { teams, userId };
}) satisfies LayoutLoad;
