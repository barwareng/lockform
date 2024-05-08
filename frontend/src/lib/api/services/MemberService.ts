import { BaseService } from '$lib/api/services/utils/BaseService';
import type { IUser } from '$utils/interfaces/users.interfaces';

export class MemberService extends BaseService {
	create(query: { email: string; role: string }) {
		return this.client.send('/api/members', {
			method: 'POST',
			query
		});
	}
}
