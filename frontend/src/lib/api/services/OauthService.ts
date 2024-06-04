import { BaseService } from '$lib/api/services/utils/BaseService';
import type { IOauthAuthorizationRequest } from '$utils/interfaces/oauth.interface';

export class OauthService extends BaseService {
	authorize(query: Partial<IOauthAuthorizationRequest>) {
		return this.client.send('/api/oauth/authorize', {
			method: 'POST',
			query
		});
	}
}
