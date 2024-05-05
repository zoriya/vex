<script>
	import {
		Label,
		Input,
		InputAddon,
		ButtonGroup,
		Button,
		Heading,
		Secondary,
		P,
		Helper,
		Span,
		Spinner,
	} from "flowbite-svelte";
	import {
		EnvelopeOutline,
		LockOutline,
		UserOutline,
	} from "flowbite-svelte-icons";
	import { enhance } from "$app/forms";
	export let form;
	let registering = false;
	let connecting = false;
</script>

<div class="text-center mb-10">
	<Heading
		tag="h1"
		class="mb-4"
		customSize="text-5xl font-extrabold  md:text-5xl lg:text-6xl"
	>
		Welcome to <Span gradient>Vex</Span>
	</Heading>
	<P class="mb-4 w-full text-center">The best place to gather your feeds</P>
</div>
<div class="main">
	<form
		method="POST"
		action="?/connect"
		use:enhance={() => {
			connecting = true;
			return async ({ update }) => {
				await update();
				connecting = false;
			};
		}}
	>
		<Heading tag="h3" class="mb-4">Login</Heading>
		<div>
			<Label for="email" class="block mb-2">Email</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<EnvelopeOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input
					id="email"
					type="email"
					name="email"
					placeholder="name@flowbite.com"
					required
				></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label for="password" class="block mb-2">Password</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<LockOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input id="password" type="password" name="password" required></Input>
			</ButtonGroup>
		</div>
		<div>
			<Button
				type="submit"
				color="primary"
				disabled={connecting}
				class="w-full"
			>
				{#if connecting}
					<Spinner class="me-3" size="4" color="white" />
				{/if}
				Connect</Button
			>
		</div>
	</form>
	<form
		method="POST"
		action="?/register"
		use:enhance={() => {
			registering = true;
			return async ({ update }) => {
				await update();
				registering = false;
			};
		}}
	>
		<Heading tag="h3" class="mb-4">Welcome</Heading>
		<div>
			<Label for="username" class="block mb-2">Your Username</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<UserOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input
					id="username"
					type="text"
					name="username"
					placeholder="Jean luc"
					required
				></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label for="email" class="block mb-2">Your Email</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<EnvelopeOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input
					id="email"
					type="email"
					name="email"
					placeholder="name@flowbite.com"
					value={form?.data?.email}
					required
				></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label
				for="password"
				class="block mb-2"
				color={form?.errors?.password ? "red" : undefined}>Your Password</Label
			>
			<ButtonGroup class="w-full">
				<InputAddon>
					<LockOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input
					color={form?.errors?.password ? "red" : undefined}
					id="password"
					name="password"
					type="password"
					required
				></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label
				for="password-repeat"
				color={form?.errors?.password ? "red" : undefined}
				class="block mb-2">Repeat Password</Label
			>
			<ButtonGroup class="w-full">
				<InputAddon>
					<LockOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input
					id="password-repeat"
					name="password-repeat"
					type="password"
					color={form?.errors?.password ? "red" : undefined}
					required
				></Input>
			</ButtonGroup>
			{#if form?.errors?.password}
				<Helper class="mt-2" color="red"
					><span class="font-medium">Invalid!:</span> Passwords do not match</Helper
				>
			{/if}
		</div>
		<div>
			<Button
				type="submit"
				color="primary"
				disabled={registering}
				class="w-full"
			>
				{#if registering}
					<Spinner class="me-3" size="4" color="white" />
				{/if}
				Register</Button
			>
		</div>
	</form>
</div>

<style>
	.main {
		display: flex;
		flex-wrap: wrap;
		gap: 1rem;
		justify-content: center;
	}
	form {
		flex-grow: 1;
		max-width: 400px;
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}
</style>
