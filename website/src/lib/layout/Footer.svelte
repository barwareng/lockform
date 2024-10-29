<script>
	import { page } from '$app/stores';
	import { PUBLIC_PLATFORM_BASE_URL } from '$env/static/public';
	import LockformLogo from '$lib/icons/LockformLogo.svelte';
	import DiscordSvg from '$lib/imgs/discord.svg';
	import TwitterSvg from '$lib/imgs/x.svg';
	import { Target } from 'lucide-svelte';

	const footerNavs = [
		// {
		// 	label: 'Product',
		// 	items: [
		// 		{
		// 			href: '/',
		// 			name: 'Email Collection'
		// 		},
		// 		{
		// 			href: '/pricing',
		// 			name: 'Pricing'
		// 		},
		// 		{
		// 			href: '/faq',
		// 			name: 'FAQ'
		// 		}
		// 	]
		// },

		{
			label: 'Community',
			items: [
				{
					href: 'https://discord.com/channels/930471500480413746/1293654482801922130',
					target: '_blank',
					name: 'Discord'
				}
				// {
				// 	href: '/',
				// 	name: 'Twitter'
				// },
				// {
				// 	href: 'mailto:hello@chatcollect.com',
				// 	name: 'Email'
				// }
			]
		},
		{
			label: 'Legal',
			items: [
				{
					href: `${$page.url.origin}/pdf/LockformToS.pdf`,
					target: '_blank',
					download: 'LockformToS.pdf',
					name: 'Terms'
				},

				{
					href: `${$page.url.origin}/pdf/LockformPrivacy.pdf`,
					target: '_blank',
					download: 'LockformPrivacy.pdf',
					name: 'Privacy'
				}
			]
		}
	];

	const footerSocials = [
		{
			href: 'https://discord.com/channels/930471500480413746/1293654482801922130',
			target: '_blank',
			name: 'Discord',
			icon: DiscordSvg
		}
		// {
		// 	href: '',
		// 	name: 'Twitter',
		// 	icon: TwitterSvg
		// }
	];
</script>

<footer>
	<div class="mx-auto w-full max-w-screen-xl xl:pb-2">
		<div class="gap-4 p-4 px-8 py-16 sm:pb-16 md:flex md:justify-between">
			<div class="mb-12 flex flex-col gap-4">
				<a class="flex items-center gap-x-2 text-lg font-black" href="/">
					<LockformLogo class="h-6 w-6" />
					lockform
				</a>
				<p class="max-w-xs">Protection for your team</p>
			</div>
			<div class="grid grid-cols-1 gap-8 sm:grid-cols-3 sm:gap-10">
				{#each footerNavs as nav}
					<div>
						<h2
							class="mb-6 text-sm font-medium uppercase tracking-tighter text-gray-900 dark:text-white"
						>
							{nav.label}
						</h2>
						<ul class="grid gap-2">
							{#each nav.items as item}
								<li>
									<a
										href={item.href}
										target={item.target}
										download={item.download}
										class="cursor-pointer text-sm font-[450] text-gray-400 duration-200 hover:text-gray-200"
									>
										{item.name}
									</a>
								</li>
							{/each}
						</ul>
					</div>
				{/each}
			</div>
		</div>

		<div
			class="flex flex-col gap-2 rounded-md border-neutral-700/20 px-8 py-4 sm:flex sm:flex-row sm:items-center sm:justify-between"
		>
			<div class="flex items-center space-x-5 sm:mt-0 sm:justify-center">
				{#each footerSocials as social}
					<a
						href={social.href}
						target={social.target}
						class="fill-gray-500 text-gray-500 hover:fill-gray-900 hover:text-gray-900 dark:hover:fill-gray-600 dark:hover:text-gray-600"
					>
						<img src={social.icon} class="size-4" alt={social.name} />
						<span class="sr-only">{social.name}</span>
					</a>
				{/each}
			</div>
			<span class="text-sm text-gray-500 sm:text-center dark:text-gray-400">
				Copyright ©
				{' '}
				{new Date().getFullYear()}
				{' '}
				<a href="/" class="cursor-pointer">Lockform</a>
				. All Rights Reserved.
			</span>
		</div>
	</div>
	<!-- Site Banner -->
</footer>
