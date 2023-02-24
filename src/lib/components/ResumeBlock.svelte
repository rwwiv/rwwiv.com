<script lang="ts">
	export let title: string;
</script>

<div
	class="flex flex-col md:flex-row content-center md:items-start text-left md:justify-left space-y-2 md:space-y-0 md:space-x-5 w-10/12 lg:w-1/2 3xl:w-1/3 mt-4 mb-8"
>
	<div class="w-full md:w-1/6 mb-4 md:mb-0 text-4xl md:text-xl font-bold">
		{title}
	</div>
	<div class="w-full md:w-5/6">
		<slot />
	</div>
</div>
