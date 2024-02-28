<script lang="ts">
	import type { HTMLInputTypeAttribute } from 'svelte/elements';

	export let value = '';
	export let label = '';
	export let id = '';
	export let type: HTMLInputTypeAttribute = 'text';

	const handleInput = (e: Event) => {
		value = (e.target as HTMLInputElement).value;
	};
</script>

<div class="flex flex-col mb-2">
	<label class="text-gray-700 text-sm font-medium mb-1" for={id}>{label}</label>
	<input
		{type}
		class=" outline-none rounded text-sm border focus:border-2 border-gray-300 focus:border-blue-500 max-w-xs w-full bg-white text-gray-500 py-2 px-3"
		{id}
		on:input={handleInput}
	/>
</div>
