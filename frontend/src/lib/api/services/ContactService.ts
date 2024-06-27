import { BaseService } from '$lib/api/services/utils/BaseService';
import type { IContact } from '$utils/interfaces/contacts.interface';

export class ContactService extends BaseService {
	create(body: Partial<IContact>) {
		return this.client.send('/api/contacts', {
			method: 'POST',
			body
		});
	}
	update(body: Partial<IContact>) {
		return this.client.send('/api/contacts', {
			method: 'PUT',
			body
		});
	}
	getAll(): Promise<IContact[]> {
		return this.client.send('/api/contacts', {
			method: 'GET'
		});
	}
	delete(query: { id: number }) {
		return this.client.send('/api/contacts', {
			method: 'DELETE',
			query
		});
	}
}
