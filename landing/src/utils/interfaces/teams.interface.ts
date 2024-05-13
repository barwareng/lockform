import type { IChannel } from './channels.interface';

export interface ITeam {
	id: string;
	name: string;
	description: string;
	createdAt: Date;
	updatedAt: Date;
	deletedAt: Date;
}

export interface Iverification extends ITeam {
	channels: IChannel[];
}
