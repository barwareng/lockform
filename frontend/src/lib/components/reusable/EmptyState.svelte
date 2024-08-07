<script lang="ts">
	export let title = '';
	export let subtitle = '';
</script>

<div class="flex h-96 w-full flex-col items-center justify-center text-center {$$props.class}">
	<slot name="leading-icon" />
	<h3 class="mt-2 text-sm font-semibold">{title}</h3>
	<p class="mt-1 max-w-md text-sm text-muted-foreground">{subtitle}</p>
	<div class="mt-6">
		<slot />
	</div>
</div>
