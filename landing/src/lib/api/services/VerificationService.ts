import { BaseService } from '$lib/api/services/utils/BaseService';

export class VerificationService extends BaseService {
	verify(body: { searchPhrase: string }) {
		return this.client.send('/public/verify', {
			method: 'POST',
			body
		});
	}
}
