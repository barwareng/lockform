<script lang="ts">
	import type { PageData } from './[searchPhrase]/$types';
	import * as Card from '$lib/components/ui/card';
	import {
		BadgeCheck,
		BadgeCheckIcon,
		MailIcon,
		MapPinIcon,
		PhoneIcon,
		SearchSlashIcon,
		ShieldCheckIcon,
		ShieldOffIcon,
		ShieldPlus,
		ShieldXIcon
	} from 'lucide-svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { parsePhoneNumber, isPossiblePhoneNumber } from 'libphonenumber-js';
	import { page } from '$app/stores';
	import Button from '$lib/components/ui/button/button.svelte';
	export let data: PageData;
	$: ({ verification } = data);
	$: console.log(verification);
	const getIcon = (type: CHANNEL) => {
		let icon: any;
		switch (type) {
			case CHANNEL.EMAIL:
				icon = MailIcon;
				break;
			case CHANNEL.PHONE:
				icon = PhoneIcon;
				break;
			default:
				icon = MapPinIcon;
				break;
		}
		return icon;
	};
	let matchingChannel: IChannel;
	$: searchPhrase = $page.url.searchParams.get('search')!;
	const getMatchingChannel = () => {
		if (isPossiblePhoneNumber(searchPhrase)) {
			matchingChannel = verification?.channels?.find(
				(channel) =>
					parsePhoneNumber(channel.value)?.formatInternational() ==
					parsePhoneNumber(searchPhrase)?.formatInternational()
			)!;
		} else {
			matchingChannel = verification?.channels?.find((channel) => channel.value == searchPhrase)!;
		}
	};
	$: verification && getMatchingChannel();
</script>

<main class="flex w-full items-center justify-center py-32">
	{#if verification}
		<Card.Root class="w-full max-w-lg overflow-hidden">
			<Card.Content class="bg-primary pt-6">
				<div class="text-background mb-2 flex select-none items-center justify-end text-[10px]">
					<ShieldCheckIcon class="mr-px h-3 w-3" />
					<p>verified by <span class="text-xs font-black">lockform</span></p>
				</div>
				<div class="text-background flex items-center gap-6">
					<div class="aspect-square w-32 overflow-hidden rounded">
						<img
							src="https://picsum.photos/seed/{verification.name}/200/300"
							alt={verification.name}
							class="h-full w-full rounded-md object-cover"
						/>
					</div>
					<div class="space-y-3">
						<div>
							<h2 class="text-xl font-semibold">{verification?.name || ''}</h2>
							{#if matchingChannel}
								<div class="mt-2 flex items-center gap-2">
									<div
										class=" bg-background text-primary flex h-8 w-8 items-center justify-center rounded-full"
									>
										<svelte:component this={getIcon(matchingChannel.type)} class="h-4 w-4" />
									</div>
									{#if matchingChannel.type == CHANNEL.PHONE}
										<p>{parsePhoneNumber(matchingChannel.value)?.formatInternational()}</p>
									{:else}
										<p>{matchingChannel.value}</p>
									{/if}
									<BadgeCheckIcon class="h-3 w-3" />
								</div>
							{/if}
						</div>
					</div>
				</div>
			</Card.Content>
			{#if verification?.channels?.length}
				<Card.Footer class="bg-muted flex flex-col items-start gap-2 pt-6">
					{#each verification.channels as channel, i}
						{@const icon = getIcon(channel.type)}
						<div class="flex items-center gap-2">
							<div
								class="bg-primary text-background flex h-8 w-8 items-center justify-center rounded-full"
							>
								<svelte:component this={icon} class="h-4 w-4" />
							</div>
							{#if channel.type == CHANNEL.PHONE}
								<p>{parsePhoneNumber(channel.value)?.formatInternational()}</p>
							{:else}
								<p>{channel.value}</p>
							{/if}
						</div>
						{#if i < verification.channels.length - 1}
							<Separator class="my-1" />
						{/if}
					{/each}
				</Card.Footer>
			{/if}
		</Card.Root>
	{:else}
		<Card.Root class="w-full max-w-lg overflow-hidden">
			<Card.Content class="bg-primary pt-6">
				<div class="text-background mb-2 flex select-none items-center justify-end text-[10px]">
					<span class="text-xs font-black">lockform</span>
				</div>
				<div
					class="text-background flex w-full flex-col items-center justify-center pt-6 text-center"
				>
					<ShieldOffIcon class="h-8 w-8" />
					<h3 class="mt-2 text-sm font-semibold">{searchPhrase}</h3>
					<p class="text-muted-foreground mt-1 max-w-md text-sm">
						We could not verify this business contact. <br /> Please proceed transactions with caution.
					</p>
					<div class="mt-6">
						<slot />
					</div>
				</div>
			</Card.Content>
			<Card.Footer class="bg-muted flex items-center gap-2 pt-6">
				<Button size="sm">
					<ShieldPlus class="mr-2 h-4 w-4" />
					Suggest verification
				</Button>
				<Button variant="destructive" size="sm">
					<ShieldXIcon class="mr-2 h-4 w-4" />
					Report fraud
				</Button>
			</Card.Footer>
		</Card.Root>
	{/if}
</main>
