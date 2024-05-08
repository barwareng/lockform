<script lang="ts">
	import { identicon } from '@dicebear/collection';
	import { createAvatar } from '@dicebear/core';
	import * as Avatar from '$lib/components/ui/avatar';
	export let seed = '';
	export let src: string;
	const avatar = createAvatar(identicon, {
		seed
	});
</script>

<div class={$$props.class}>
	<Avatar.Root>
		<Avatar.Image {src} alt={seed} />
		<Avatar.Fallback>
			{@html avatar}
		</Avatar.Fallback>
	</Avatar.Root>
</div>
