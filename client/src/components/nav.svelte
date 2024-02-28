<script lang="ts">
	import type { User } from '../types/user';
	import Logo from './logo.svelte';

	export let user: User | null = null;

	const handleSignout = async () => {
		try {
			const res = await fetch('https://ticketing.io/api/users/signout', {
				method: 'post',
				credentials: 'include'
			});

			window.location.href = '/';
		} catch (error) {
			console.error(error);
		}
	};
</script>

<nav class="flex justify-center border-b border-b-gray-300">
	<div class="flex w-full max-w-[80rem] flex-row items-center justify-between py-3 mx-2">
		<Logo width={30} height={30} className="fill-gray-800" />
		<div class="flex flex-row items-center">
			<a class=" cursor-pointer" href="/">Home</a>
			{#if user}
				<button class="ml-5 cursor-pointer" type="button" on:click={handleSignout} aria-label="Signout">Signout</button>
			{:else}
				<a class="ml-5 cursor-pointer" href="/auth/signin">Signin</a>
				<a class="ml-5 cursor-pointer" href="/auth/signup">Signup</a>
			{/if}
		</div>
	</div>
</nav>
