import { BaseService } from '$lib/api/services/utils/BaseService';
import type { IMember } from '$utils/interfaces/users.interfaces';

export class MemberService extends BaseService {
	create(query: { email: string; role: string }) {
		return this.client.send('/api/members', {
			method: 'POST',
			query
		});
	}
	update(query: { userId: string; email: string; role: string }) {
		return this.client.send('/api/members', {
			method: 'PUT',
			query
		});
	}
	getAll(): Promise<IMember[]> {
		return this.client.send('/api/members', {
			method: 'GET'
		});
	}
	delete(query: { userId: string }) {
		return this.client.send('/api/members', {
			method: 'DELETE',
			query
		});
	}
}
