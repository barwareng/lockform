<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import * as Pagination from '$lib/components/ui/pagination';
	export let perPage = 2;

	export let count = 10;
	// Write a function that takes a page number (as page) adds a query parameter to the current URL and navigates to the new URL
	const goToPage = async (pageNumber: number) => {
		if (pageNumber < 1) return;
		$page.url.searchParams.set('page', pageNumber.toString());
		await goto($page.url, { invalidateAll: true });
	};
</script>

<Pagination.Root bind:count bind:perPage let:pages let:currentPage>
	<Pagination.Content>
		<Pagination.Item>
			<Pagination.PrevButton on:click={() => goToPage(currentPage - 1)} />
		</Pagination.Item>
		{#each pages as page (page.key)}
			{#if page.type === 'ellipsis'}
				<Pagination.Item>
					<Pagination.Ellipsis />
				</Pagination.Item>
			{:else}
				<Pagination.Item isVisible={currentPage == page.value}>
					<Pagination.Link
						on:click={() => {
							console.log(page.value, 'page.value');
							goToPage(page.value);
						}}
						{page}
						isActive={currentPage == page.value}
					>
						{page.value}
					</Pagination.Link>
				</Pagination.Item>
			{/if}
		{/each}
		<Pagination.Item>
			<Pagination.NextButton on:click={() => goToPage(currentPage + 1)} />
		</Pagination.Item>
	</Pagination.Content>
</Pagination.Root>
