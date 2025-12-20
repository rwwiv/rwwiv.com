<script lang="ts">
	export let value: number[] = [50];
	export let min = 0;
	export let max = 100;
	export let step = 1;
	export let name: string[] = [];

	$: progress = ((value[0] - min) / (max - min)) * 100;
</script>

<div class="slider">
	<input
		type="range"
		{min}
		{max}
		{step}
		name={name[0] ?? undefined}
		bind:value={value[0]}
		style="--progress: {progress}%"
	/>
</div>

<style>
	.slider {
		width: 100%;
		padding: 8px 0;
	}

	input[type="range"] {
		-webkit-appearance: none;
		appearance: none;
		width: 100%;
		height: 6px;
		border-radius: 3px;
		background: linear-gradient(
			to right,
			var(--progress-bg, #6e6e6e) 0%,
			var(--progress-bg, #6e6e6e) var(--progress),
			var(--track-bg, #cdcdcd) var(--progress),
			var(--track-bg, #cdcdcd) 100%
		);
		outline: none;
		cursor: pointer;
	}

	input[type="range"]::-webkit-slider-thumb {
		-webkit-appearance: none;
		appearance: none;
		width: 18px;
		height: 18px;
		border-radius: 50%;
		background: var(--thumb-bg, #dadada);
		cursor: pointer;
		border: none;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
	}

	input[type="range"]::-moz-range-thumb {
		width: 18px;
		height: 18px;
		border-radius: 50%;
		background: var(--thumb-bg, #dadada);
		cursor: pointer;
		border: none;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
	}

	input[type="range"]:focus::-webkit-slider-thumb {
		box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.3);
	}

	input[type="range"]:focus::-moz-range-thumb {
		box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.3);
	}
</style>
