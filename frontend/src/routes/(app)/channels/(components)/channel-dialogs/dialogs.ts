import { CHANNEL, type ChannelCategory } from '$utils/interfaces/channels.interface';
import AddressDialog from './address-dialog.svelte';
import EmailDialog from './email-dialog.svelte';
import PhoneDialog from './phone-dialog.svelte';

export const channelDialogs: {
	category: ChannelCategory;
	dialogs: {
		component: any;
		type: CHANNEL;
	}[];
}[] = [
	{
		category: 'General',
		dialogs: [
			{ component: EmailDialog, type: CHANNEL.EMAIL },
			{ component: PhoneDialog, type: CHANNEL.PHONE },
			{ component: AddressDialog, type: CHANNEL.ADDRESS }
		]
	}
];
