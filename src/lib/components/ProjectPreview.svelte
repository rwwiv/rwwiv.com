<script lang="ts">
	export let src: string;
	export let title: string;
	export let subtitle: string;
	export let sourceUrl: string | undefined = undefined;
	export let siteUrl: string | undefined = undefined;
</script>

<div class="flex flex-col items-center space-y-6">
	<div class="relative max-w-[350px] md:max-w-[400px]">
		<img
			class="duration-300 rounded border border-neutral-200 dark:border-none"
			{src}
			alt="TeachMeTo"
		/>
		<div
			class="absolute top-0 left-0 h-full w-full bg-black bg-opacity-0 dark:bg-opacity-20 hover:dark:bg-opacity-0 transition-all duration-300"
		/>
	</div>
	<div>
		<div class="space-y-1">
			<h2 class="text-2xl text-green-700 dark:text-green-200">{title}</h2>
			<h3 class="text-lg">{subtitle}</h3>
		</div>
		<sub class="text-sm">
			{#if sourceUrl}
				<a
					href={sourceUrl}
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Source
				</a>
			{/if}
			{#if siteUrl}
				{#if sourceUrl}
					&bull;
				{/if}
				<a
					href={siteUrl}
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Site
				</a>
			{/if}
		</sub>
	</div>
</div>
