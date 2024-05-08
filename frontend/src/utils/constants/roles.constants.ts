import { ROLE_VALUES } from '$utils/interfaces/roles.interface';

export const ROLES: { name: string; value: ROLE_VALUES; description: string }[] = [
	{
		name: 'Viewer',
		value: ROLE_VALUES.VIEWER,
		description: 'Can view and comment.'
	},
	{
		name: 'Developer',
		value: ROLE_VALUES.DEVELOPER,
		description: 'Can view, comment and edit.'
	},
	{
		name: 'Billing',
		value: ROLE_VALUES.BILLING,
		description: 'Can view, comment and manage billing.'
	},
	{
		name: 'Admin',
		value: ROLE_VALUES.ADMIN,
		description: 'Access to all resources.'
	},
	{
		name: 'Owner',
		value: ROLE_VALUES.OWNER,
		description: 'Admin-level access to all resources.'
	}
];
