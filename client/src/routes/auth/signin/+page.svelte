<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '../../../components/button.svelte';
	import Input from '../../../components/input.svelte';
	import Logo from '../../../components/logo.svelte';
	import { sendRequest } from '../../../helpers/request';
	import type { FormError } from '../../../types/form-error';

	let email = '';
	let password = '';

	let emailErrorMessage = '';
	let passwordErrorMessage = '';

	const handleSubmit = async (e: Event) => {
		e.preventDefault();

		const errors = await sendRequest('https://ticketing.io/api/users/signin', 'post', {
			email,
			password
		});

		if (errors.length === 0) {
			goto('/')
			return;
		}

		// Check if field key exists in array of errors
		const fieldKeyExists = errors.filter((err) => err.hasOwnProperty('field')).length > 0;

		// If field key does not exist, then set error message to first error in array.
		if (!fieldKeyExists) {
			emailErrorMessage = errors[0].message;
			passwordErrorMessage = errors[0].message;
		}

		// Set the error message for the fields
		errors.forEach((error) => {
			setFieldError(error);
		});
	};

	const setFieldError = ({ message, field }: FormError) => {
		switch (field) {
			case 'email':
				emailErrorMessage = message;
				return;
			case 'password':
				passwordErrorMessage = message;
				return;
		}
	};
</script>

<section class="flex flex-col items-center justify-center">
	<Logo margin="mt-10" width={50} height={50} className="fill-gray-800" />
	<h2 class="mt-5 text-center text-xl font-medium text-gray-800">Signin</h2>
	<form
		on:submit={handleSubmit}
		class="mt-5 w-full max-w-xs rounded border border-gray-300 bg-white p-5 sm:p-5 md:max-w-[18rem]"
	>
		<Input bind:errorMessage={emailErrorMessage} bind:value={email} label="Email:" />
		<Input
			bind:errorMessage={passwordErrorMessage}
			type="password"
			bind:value={password}
			label="Password:"
		/>
		<Button>Submit</Button>
	</form>

	<div
		class="mt-5 flex w-full max-w-xs items-center justify-center rounded border border-gray-300 p-5 sm:p-5 md:max-w-[18rem]"
	>
		<span class="text-sm"
			>New to Ticketing? <a href="/auth/signup" class="cursor-pointer text-blue-500 underline"
				>Create a account</a
			></span
		>
	</div>
</section>
