import { BaseService } from '$lib/api/services/utils/BaseService';

export class UserService extends BaseService {
	create(body: any) {
		return this.client.send('/api/users', {
			method: 'POST',
			body
		});
	}
}
