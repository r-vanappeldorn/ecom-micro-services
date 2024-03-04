<script lang="ts">
	import Nav from '../components/nav.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<Nav user={data.user} />

<section
	class="gird-cols-1 mx-auto mt-5 grid max-w-[80rem] gap-2 px-2 md:grid-cols-2 md:gap-2 lg:grid-cols-3 lg:gap-3"
>
	{#each data.tickets as { id, title, price }}
		<article class="rounded border border-gray-300 bg-white px-4 py-2 text-gray-500" {id}>
			<h2 class="text-lg font-medium text-gray-700">{title}</h2>
			<span>${price}</span>
		</article>
	{/each}
</section>
