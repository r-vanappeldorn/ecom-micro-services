<script lang="ts">
	import type { HTMLInputTypeAttribute } from 'svelte/elements';

	export let value = '';
	export let label = '';
	export let id = '';
	export let type: HTMLInputTypeAttribute = 'text';
	export let errorMessage = '';

	const handleInput = (e: Event) => {
		errorMessage = '';
		value = (e.target as HTMLInputElement).value;
	};

	$: errorTheme =
		errorMessage !== '' ? 'text-red-500 border-red-500' : 'text-gray-500 border-gray-300';
</script>

<div class="mb-2 flex flex-col">
	<label class="mb-1 text-sm font-medium text-gray-700" for={id}>{label}</label>
	<input
		{type}
		class="w-full max-w-xs rounded border bg-white px-3 py-2 text-sm outline-none focus:border-2 focus:border-blue-500 {errorTheme}"
		{id}
		on:input={handleInput}
	/>
	{#if errorMessage !== ''}
		<span class="mt-1 text-xs text-red-500">{errorMessage}</span>
	{/if}
</div>
