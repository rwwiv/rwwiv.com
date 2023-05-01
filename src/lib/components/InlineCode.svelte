<script lang="ts">
	import hljs from "highlight.js/lib/common";
	// @ts-ignore
	import hljs_svelte from "highlightjs-svelte";

	hljs_svelte(hljs);

	export let code: string;
	export let language: string;
</script>

<span>
	<code class="bg-neutral-100 dark:bg-neutral-800 py-[3px] px-1 rounded whitespace-nowrap text-sm">
		{@html hljs.highlight(code, { language }).value}
	</code>
</span>
