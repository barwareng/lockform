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

export interface ITeamContact {
	id: number;
	contactId: number;
	teamId: string;
	addedById: string;
	isTrusted: boolean;
	reasonForUntrusting: string;
	addedBy?: {
		id: string;
		name: string;
	};
}

export interface ISaveContactBody {
	contact: Partial<IContact>;
	teamContact: Partial<ITeamContact>;
}
