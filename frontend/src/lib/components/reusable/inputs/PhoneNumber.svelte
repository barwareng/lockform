<script lang="ts">
	import { TelInput, normalizedCountries } from 'svelte-tel-input';
	import * as Popover from '$lib/components/ui/popover';
	import { Check, ChevronsUpDownIcon } from 'lucide-svelte';

	import { Button } from '$lib/components/ui/button';
	import * as Command from '$lib/components/ui/command';
	import type {
		DetailedValue,
		E164Number,
		CountryCode,
		TelInputOptions
	} from 'svelte-tel-input/types';
	import { cn } from '$lib/utils';
	import { closeAndRefocusTrigger } from '$utils';
	export let phoneNumber: E164Number | undefined;

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
</script>

<div class="flex">
	<Popover.Root bind:open={showCountryCodePopOver} let:ids>
		<Popover.Trigger asChild let:builder>
			<Button
				builders={[builder]}
				variant="outline"
				role="combobox"
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
								class={cn('mr-2 h-4 w-4', country !== currentCountry.iso2 && 'text-transparent')}
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
