import { BaseService } from '$lib/api/services/utils/BaseService';

export class VerificationService extends BaseService {
	verify(query: { searchPhrase: string }) {
		return this.client.send('/public/verify', {
			method: 'GET',
			query
		});
	}
}
