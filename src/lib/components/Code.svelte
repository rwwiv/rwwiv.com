<script lang="ts">
	import hljs from "highlight.js/lib/common";
	// @ts-ignore
	import hljs_svelte from "highlightjs-svelte";

	hljs_svelte(hljs);

	export let code: string;
	export let language: string;
</script>

<div
	class="text-sm bg-neutral-100 dark:bg-neutral-800 py-3 px-4 rounded border border-neutral-200 dark:border-neutral-700 whitespace-pre overflow-x-scroll"
>
	<code>
		{@html hljs.highlight(code, { language }).value}
	</code>
</div>
