import type { ROLE_VALUES } from './roles.interface';

export interface IUser {
	id: string;
	email: string;
	phoneNumber: string;
	firstName: string;
	middleName: string;
	lastName: string;
}

export interface IMember extends IUser {
	role: ROLE_VALUES;
}
