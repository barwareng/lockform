import { BaseService } from '$lib/api/services/utils/BaseService';
import type { IChannel } from '$utils/interfaces/channels.interface';

export class ChannelService extends BaseService {
	create(body: Partial<IChannel>) {
		return this.client.send('/api/channels', {
			method: 'POST',
			body
		});
	}
	update(body: Partial<IChannel>) {
		return this.client.send('/api/channels', {
			method: 'PUT',
			body
		});
	}
	getAll(): Promise<IChannel[]> {
		return this.client.send('/api/channels', {
			method: 'GET'
		});
	}
	delete(query: { channelId: string }) {
		return this.client.send('/api/channels', {
			method: 'DELETE',
			query
		});
	}
}
