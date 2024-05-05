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
		Span,
	} from "flowbite-svelte";
	import { EnvelopeOutline, LockOutline } from "flowbite-svelte-icons";
</script>

<div class="text-center mb-10">
	<Heading
		tag="h1"
		class="mb-4"
		customSize="text-5xl font-extrabold  md:text-5xl lg:text-6xl"
	>
		Welcome to <Span gradient>Vex</Span>
	</Heading>
	<P class="mb-4 w-full text-center">The best place to share your thoughts with the world</P>
</div>
<div class="main">
	<form>
		<Heading tag="h3" class="mb-4">Login</Heading>
		<div>
			<Label for="email" class="block mb-2">Email</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<EnvelopeOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input id="email" type="email" placeholder="name@flowbite.com"></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label for="password" class="block mb-2">Password</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<LockOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input id="password" type="password"></Input>
			</ButtonGroup>
		</div>
		<div>
			<Button type="submit" color="primary" class="w-full">Connect</Button>
		</div>
	</form>
	<form>
		<Heading tag="h3" class="mb-4">Welcome</Heading>
		<div>
			<Label for="email" class="block mb-2">Your Email</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<EnvelopeOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input id="email" type="email" placeholder="name@flowbite.com"></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label for="password" class="block mb-2">Your Password</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<LockOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input id="password" type="password"></Input>
			</ButtonGroup>
		</div>
		<div>
			<Label for="password-repeat" class="block mb-2">Repeat Password</Label>
			<ButtonGroup class="w-full">
				<InputAddon>
					<LockOutline class="w-4 h-4 text-gray-500 dark:text-gray-400" />
				</InputAddon>
				<Input id="password-repeat" type="password"></Input>
			</ButtonGroup>
		</div>
		<div>
			<Button type="submit" color="primary" class="w-full">Register</Button>
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
