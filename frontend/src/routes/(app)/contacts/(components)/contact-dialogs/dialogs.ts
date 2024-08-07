import { CONTACT } from '$utils/interfaces/contacts.interface';
import AddressDialog from './address-dialog.svelte';
import EmailDialog from './email-dialog.svelte';
import PhoneDialog from './phone-dialog.svelte';

export const contactDialogs: {
	component: any;
	type: CONTACT;
}[] = [
	{ component: EmailDialog, type: CONTACT.EMAIL },
	{ component: PhoneDialog, type: CONTACT.PHONE },
	{ component: AddressDialog, type: CONTACT.ADDRESS }
];
