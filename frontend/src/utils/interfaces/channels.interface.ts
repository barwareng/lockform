export enum CHANNEL {
	EMAIL = 'email',
	PHONE = 'phone',
	ADDRESS = 'address'
}
export type ChannelCategory = 'General' | 'Social' | 'Payment' | 'Messaging' | 'Business';
export interface IChannel {
	id: number;
	value: string;
	label: string;
	category: ChannelCategory;
	type: CHANNEL;
	url?: string;
	isPublic: boolean;
}
