export enum CONTACT {
	EMAIL = 'email',
	PHONE = 'phone',
	ADDRESS = 'address'
}
export type ContactCategory = 'General' | 'Social' | 'Payment' | 'Messaging' | 'Business';
export interface IContact {
	id: number;
	value: string;
	label: string;
	domain: string;
	url: string;
	type: CONTACT;
	category: ContactCategory;
}

export interface TeamContact {
	id: number;
	contactId: number;
	teamId: string;
	addedById: string;
	isTrusted: string;
	reasonForUntrusting: string;
	addedBy?: {
		id: string;
		name: string;
	};
}
