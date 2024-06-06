import { BaseService } from '$lib/api/services/utils/BaseService';
import type { ITrustedContact } from '$utils/interfaces/trusted-contacts.interface';

export class TrustedContactService extends BaseService {
	create(body: Partial<ITrustedContact>) {
		return this.client.send('/api/trusted-contacts', {
			method: 'POST',
			body
		});
	}
	update(body: Partial<ITrustedContact>) {
		return this.client.send('/api/trusted-contacts', {
			method: 'PUT',
			body
		});
	}
	getAll(): Promise<ITrustedContact[]> {
		return this.client.send('/api/trusted-contacts', {
			method: 'GET'
		});
	}
	delete(query: { id: number }) {
		return this.client.send('/api/trusted-contacts', {
			method: 'DELETE',
			query
		});
	}
}
