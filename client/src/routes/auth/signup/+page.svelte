<!--create signup form-->
<script>
	import Button from '../../../components/button.svelte';
	import Input from '../../../components/input.svelte';
	import Logo from '../../../components/logo.svelte';

	let email = '';
	let password = '';
</script>

<section class="flex flex-col items-center justify-center">
	<Logo margin="mt-10" width={50} height={50} className="fill-gray-800" />
	<h2 class="font-medium text-xl text-gray-800 text-center mt-5">Signup</h2>
	<form
		class="border border-gray-300 bg-white rounded sm:p-5 p-5 mt-5 w-full md:max-w-[18rem] max-w-xs"
	>
		<Input bind:value={email} label="Email:" />
		<Input type="password" bind:value={password} label="Password:" />
		<Button>Submit</Button>
	</form>

	<div
		class="border border-gray-300 rounded sm:p-5 p-5 mt-5 w-full md:max-w-[18rem] max-w-xs flex items-center justify-center"
	>
		<span class="text-sm"
			>Already have an account? <a
				href="/auth/signin"
				class="text-blue-500 underline cursor-pointer">Signin</a
			></span
		>
	</div>
</section>
