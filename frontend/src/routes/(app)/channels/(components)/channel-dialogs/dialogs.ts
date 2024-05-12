import AddressDialog from './address-dialog.svelte';
import EmailDialog from './email-dialog.svelte';
import PhoneDialog from './phone-dialog.svelte';

export const channelDialogs = [
	{
		type: 'General',
		dialogs: [EmailDialog, PhoneDialog, AddressDialog]
	}
];
