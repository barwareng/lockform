import { BaseService } from '$lib/api/services/utils/BaseService';

export class UserService extends BaseService {
	create(body: any) {
		return this.client.send('/api/users', {
			method: 'POST',
			body
		});
	}
	update(query, body) {
		return this.client.send('/api/users', {
			method: 'PUT',
			body,
			query
		});
	}
}
