import { BaseService } from '$lib/api/services/utils/BaseService';
import type { Iverification } from '$utils/interfaces/teams.interface';

export class VerificationService extends BaseService {
	verify(query: { searchPhrase: string }): Promise<Iverification> {
		return this.client.send('/public/verify', {
			method: 'GET',
			query
		});
	}
}
