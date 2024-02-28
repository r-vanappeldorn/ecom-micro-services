<script lang="ts">
	import type { User } from '../types/user';
	import Logo from './logo.svelte';

	export let user: User | null = null;
</script>

<nav class="flex justify-center border-b border-b-gray-300">
	<div class="border-gra flex w-full max-w-[80rem] flex-row items-center justify-between py-3">
		<Logo width={30} height={30} className="fill-gray-800" />
		<div class="flex flex-row items-center">
			<a class=" cursor-pointer" href="/">Home</a>
			{#if user}
				<a class="ml-5 cursor-pointer" href="/auth/signout">Signout</a>
			{:else}
				<a class="ml-5 cursor-pointer" href="/auth/signin">Signin</a>
				<a class="ml-5 cursor-pointer" href="/auth/signup">Signup</a>
			{/if}
		</div>
	</div>
</nav>
