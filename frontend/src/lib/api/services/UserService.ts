import { BaseService } from '$lib/api/services/utils/BaseService';

export class UserService extends BaseService {
	update(body) {
		return this.client.send('/api/users', {
			method: 'PUT',
			body
		});
	}
	getProfile() {
		return this.client.send('/api/users', {
			method: 'GET'
		});
	}
}
