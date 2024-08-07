<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { LoaderCircleIcon, Upload } from 'lucide-svelte';
	import imageCompression, { type Options } from 'browser-image-compression';
	import { FileDrop } from 'svelte-droplet';
	import { toastError } from '$utils/toasts';
	import { VITE_API_BASE_URL } from '$lib/env';
	const dispatch = createEventDispatcher();
	export let accept: string[] | null = ['images/*'];
	export let path: string;
	export let name: string;
	export let uploading = false;
	const handleFiles = async (files: File[]) => {
		uploading = true;
		const options: Options = {
			maxSizeMB: 1.5,
			maxWidthOrHeight: 400,
			useWebWorker: true,
			maxIteration: 6
		};
		try {
			const file = await imageCompression(files[0], options);
			await handleUpload(file);
		} catch (error) {
			toastError(error);
		}
	};
	const handleUpload = async (uploadFile: File) => {
		const formData = new FormData();
		formData.append(name, uploadFile);
		const xhr = new XMLHttpRequest();
		xhr.onload = () => {
			if (xhr.status != 200) {
				toastError(JSON.parse(xhr.response)!.error);
			} else {
				dispatch('file-uploaded', { url: JSON.parse(xhr.response)!.data });
			}
			uploading = false;
		};
		xhr.onerror = (e) => {
			uploading = false;
			toastError('Failed to upload your file');
		};
		xhr.open('POST', `${VITE_API_BASE_URL}/admin/${path}`, true);
		xhr.send(formData);
	};
</script>

<FileDrop {handleFiles}>
	<div class="flex w-fit gap-4 rounded border border-dashed hover:bg-muted">
		<div class="flex aspect-square min-w-20 flex-col items-center justify-center">
			{#if uploading}
				<LoaderCircleIcon class="h-4 w-4 animate-spin" />
			{:else}
				<Upload class="h-4 w-4" />
			{/if}
		</div>
	</div>
</FileDrop>
