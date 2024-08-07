import { BaseService } from '$lib/api/services/utils/BaseService';
import type {
	IContact,
	IContactList,
	ISaveContactBody
} from '$utils/interfaces/contacts.interface';

export class ContactService extends BaseService {
	create(body: Partial<ISaveContactBody>) {
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
	getAll(): Promise<IContactList[]> {
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
