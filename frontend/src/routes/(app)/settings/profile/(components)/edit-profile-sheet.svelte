<script lang="ts">
	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { PencilIcon } from 'lucide-svelte';
	import { isMobileDevice } from '$utils';
	import type { IUser } from '$utils/interfaces/users.interfaces';
	import Label from '$lib/components/ui/label/label.svelte';
	import FileUpload from '$lib/components/reusable/files/FileUpload.svelte';

	export let profile: Partial<IUser>;
	let open = true;
</script>

<Sheet.Root bind:open>
	<Sheet.Trigger asChild let:builder>
		<Button builders={[builder]} class="text-xs">
			<PencilIcon class="mr-2 h-4 w-4" />
			Edit Profile
		</Button>
	</Sheet.Trigger>
	<Sheet.Content side={isMobileDevice() ? 'bottom' : 'right'}>
		<form class="flex max-h-screen min-h-full flex-col justify-between">
			<div class="space-y-3 overflow-y-scroll px-1">
				<Sheet.Header class="bg-background sticky top-0 mb-6">
					<Sheet.Title>Edit profile</Sheet.Title>
					<Sheet.Description class="text-xs">
						Change how your profile will appear on the site.
					</Sheet.Description>
				</Sheet.Header>
				<FileUpload />
				<div class="flex-1 space-y-1">
					<Label>First name</Label>
					<Input bind:value={profile.firstName} placeholder="Luke" />
				</div>
				<div class="flex-1 space-y-1">
					<Label>Last name</Label>
					<Input bind:value={profile.lastName} placeholder="Skywalker" />
				</div>
				<div class="flex-1 space-y-1">
					<Label>Email</Label>
					<Input bind:value={profile.email} placeholder="luke@starwars.com" />
				</div>
				<div class="flex-1 space-y-1">
					<Label>Phone number</Label>
					<Input bind:value={profile.phoneNumber} placeholder="(702)-234-5566" />
				</div>
			</div>
			<Sheet.Footer class="sticky bottom-0">
				<Sheet.Close asChild let:builder class="flex">
					<Button builders={[builder]} variant="secondary">Cancel</Button>
					<Button>Save</Button>
				</Sheet.Close>
			</Sheet.Footer>
		</form>
	</Sheet.Content>
</Sheet.Root>
