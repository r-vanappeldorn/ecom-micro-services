<script lang="ts">
	import Nav from '../components/nav.svelte';
	import type { User } from '../types/user';
	import type { PageData } from './$types';
	export let data: PageData;

	let user: User | null = data.user as User | null;
</script>

<Nav {user} />

{#if data.user}
	<h1>Signed in</h1>
	<pre>
        {JSON.stringify(data.user, null, 2)}
    </pre>
{/if}
