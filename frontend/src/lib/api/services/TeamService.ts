import { BaseService } from '$lib/api/services/utils/BaseService';
import type { ITeam } from '$utils/interfaces/teams.interface';

export class TeamService extends BaseService {
	create(body: Partial<ITeam>): Promise<ITeam> {
		return this.client.send('/api/teams', {
			method: 'POST',
			body
		});
	}
}
