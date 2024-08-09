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
	isTrusted: boolean;
	reasonForUntrusting: string;
}

export interface ISaveContactBody {
	contact: Partial<IContact>;
	teamContact: Partial<ITeamContact>;
}
export interface IContactList extends IContact, ITeamContact {
	addedBy: {
		id: string;
		name: string;
		email: string;
	};
}
