<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Label } from '$lib/components/ui/label';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { Check, ChevronsUpDownIcon, PhoneIcon } from 'lucide-svelte';
	import * as Select from '$lib/components/ui/select';
	import type { IChannel } from '$utils/interfaces/channels.interface';
	import { toastError, toastSuccess } from '$utils/toasts';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import { TelInput, normalizedCountries } from 'svelte-tel-input';

	import * as Command from '$lib/components/ui/command';
	import type {
		DetailedValue,
		E164Number,
		CountryCode,
		TelInputOptions
	} from 'svelte-tel-input/types';
	import { cn } from '$lib/utils';
	import { closeAndRefocusTrigger } from '$utils';
	let channel: Partial<IChannel> = {};
	let addingChannel = false;
	let open = false;

	// E164 formatted value, usually you should store and use this.
	let phoneNumber: E164Number | null;
	$: phoneNumber = null;

	// Selected country
	let country: CountryCode | null = 'US';

	// Validity
	let valid: boolean;

	// Phone number details
	let detailedValue: DetailedValue | null = null;

	let options: TelInputOptions = {
		invalidateOnCountryChange: true,
		format: 'international'
	};
	let showCountryCodePopOver = false;
	const addChannel = async () => {
		try {
			addingChannel = true;
			channel.value = phoneNumber?.replaceAll(' ', '');
			// TODO API call
			addingChannel = false;
			open = false;
			toastSuccess('Channel successfully added.');
		} catch (error) {
			addingChannel = false;
			toastError(error);
		}
	};
</script>

<BaseDialog title="Phone" icon={PhoneIcon} bind:open>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Number</Label>
			<!-- <Input type="email" bind:value={channel.value} placeholder="mail@acme.com" /> -->
			<div class="flex">
				<Popover.Root bind:open={showCountryCodePopOver} let:ids>
					<Popover.Trigger asChild let:builder>
						<Button
							builders={[builder]}
							variant="outline"
							role="combobox"
							aria-expanded={open}
							class=" w-[200px] justify-between rounded-r-none "
						>
							{country}
							<ChevronsUpDownIcon class="ml-2 h-4 w-4 shrink-0 opacity-50" />
						</Button>
					</Popover.Trigger>
					<Popover.Content class="max-h-80 w-[200px] overflow-y-scroll p-0">
						<Command.Root>
							<Command.Input placeholder="Search country code..." />
							<Command.Empty>No country code found.</Command.Empty>
							<Command.Group>
								{#each normalizedCountries as currentCountry (currentCountry.id)}
									<Command.Item
										value={currentCountry.iso2}
										onSelect={(currentValue) => {
											country = currentValue;
											showCountryCodePopOver = closeAndRefocusTrigger(ids.trigger);
										}}
									>
										<Check
											class={cn(
												'mr-2 h-4 w-4',
												country !== currentCountry.iso2 && 'text-transparent'
											)}
										/>
										{currentCountry.iso2} (+{currentCountry.dialCode})
									</Command.Item>
								{/each}
							</Command.Group>
						</Command.Root>
					</Popover.Content>
				</Popover.Root>
				<TelInput
					{options}
					bind:country
					bind:valid
					bind:value={phoneNumber}
					bind:detailedValue
					class={cn(
						'w-full rounded-r-lg border border-l-0 bg-transparent px-4 py-1 outline-none',
						!valid ? 'border-destructive' : ''
					)}
				/>
			</div>
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input type="text" bind:value={channel.label} placeholder="e.g Work" />
		</div>
		<div class="flex items-center justify-end gap-x-4">
			<Button variant="secondary" on:click={() => (open = false)}>Cancel</Button>
			<Button disabled={addingChannel} on:click={addChannel}>
				<ButtonLoadingSpinner bind:state={addingChannel} />
				Save
			</Button>
		</div>
	</div>
</BaseDialog>
