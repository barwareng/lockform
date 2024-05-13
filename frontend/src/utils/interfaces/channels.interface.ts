export enum CHANNEL {
	EMAIL = 'email',
	PHONE = 'phone',
	ADDRESS = 'address'
}
export interface IChannel {
	id: number;
	value: string;
	label: string;
	type: CHANNEL;
	url?: string;
}
