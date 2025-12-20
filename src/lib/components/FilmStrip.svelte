<script lang="ts">
	import { createEventDispatcher } from "svelte";

	export let images: Array<{
		src: string;
		alt: string;
		fullSrc?: string;
	}> = [];
	export let tilt: number = -2; // degrees of tilt
	export let hoverEffect: "reveal" | "darken" = "reveal"; // hover behavior
	export let maxImages: number | undefined = undefined; // for consistent sprocket count
	export let startIndex: number = 0; // global starting index for frame numbers

	const dispatch = createEventDispatcher<{ select: { index: number } }>();

	function handleClick(index: number) {
		dispatch("select", { index });
	}

	// Sprocket count based on max images for consistent appearance across strips
	$: sprocketCount = (maxImages ?? images.length) * 5;
</script>

<div class="film-strip-wrapper" style="--tilt: {tilt}deg">
	<div class="film-strip" class:hover-darken={hoverEffect === "darken"}>
		<!-- Top sprocket holes -->
		<div class="sprockets">
			{#each Array(sprocketCount) as _}
				<div class="sprocket-hole"></div>
			{/each}
		</div>

		<!-- Film frames -->
		<div class="frames">
			{#each images as image, i}
				<button class="frame" on:click={() => handleClick(i)}>
					<div class="frame-content">
						<img src={image.src} alt={image.alt} />
					</div>
					<span class="frame-number">{String(startIndex + i + 1).padStart(2, "0")}</span>
				</button>
			{/each}
		</div>

		<!-- Bottom sprocket holes -->
		<div class="sprockets">
			{#each Array(sprocketCount) as _}
				<div class="sprocket-hole"></div>
			{/each}
		</div>
	</div>
</div>

<style lang="postcss">
	.film-strip-wrapper {
		--shadow-color: rgba(0, 0, 0, 0.3);
		--shadow-color-hover: rgba(0, 0, 0, 0.4);
		--hole-bg: theme(colors.neutral.200);

		transform: rotate(var(--tilt));
		transition:
			transform 0.3s ease,
			filter 0.3s ease;
		filter: drop-shadow(0 4px 8px var(--shadow-color));
	}

	:global(.dark) .film-strip-wrapper {
		--shadow-color: rgba(255, 255, 255, 0.06);
		--shadow-color-hover: rgba(255, 255, 255, 0.1);
		--hole-bg: theme(colors.neutral.800);
	}

	.film-strip-wrapper:hover {
		transform: rotate(var(--tilt)) translateY(-2px);
		filter: drop-shadow(0 6px 10px var(--shadow-color-hover));
	}

	.film-strip {
		--film-bg: rgba(196, 92, 26, 0.95);
		--film-bg-dark: rgba(158, 74, 21, 0.95);
		--sprocket-width: 10px;
		--sprocket-height: 6px;
		--frame-gap: 8px;

		display: flex;
		flex-direction: column;
		background: linear-gradient(180deg, var(--film-bg-dark) 0%, var(--film-bg) 15%, var(--film-bg) 85%, var(--film-bg-dark) 100%);
		border-radius: 2px;
		padding: 0 4px;
	}

	.sprockets {
		display: flex;
		flex-direction: row;
		justify-content: space-evenly;
		padding: 4px 0;
	}

	.sprocket-hole {
		width: var(--sprocket-width);
		height: var(--sprocket-height);
		background: var(--hole-bg);
		border-radius: 1px;
	}

	.frames {
		display: flex;
		gap: var(--frame-gap);
		padding: 4px 8px;
		padding-bottom: 12px; /* Space for frame numbers */
	}

	.frame {
		position: relative;
		background: transparent;
		border: none;
		padding: 0;
		cursor: pointer;
		transition: transform 0.2s ease;
	}

	.frame:hover {
		transform: scale(1.02);
	}

	.frame:focus {
		outline: 2px solid rgba(255, 255, 255, 0.5);
		outline-offset: 2px;
	}

	.frame-content {
		width: 120px;
		height: 80px;
		overflow: hidden;
		border: 1px solid rgba(0, 0, 0, 0.3);
	}

	.frame-content img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		/* Color negative film effect: invert + orange mask tint */
		filter: invert(0.85) sepia(0.6) saturate(1.4) hue-rotate(345deg) brightness(0.9);
		transition: filter 0.3s ease;
	}

	/* Default hover: reveal true colors */
	.frame:hover .frame-content img {
		filter: invert(0) sepia(0) hue-rotate(0deg) saturate(1) brightness(1);
	}

	/* Optional: darken on hover instead of reveal */
	.hover-darken .frame:hover .frame-content img {
		filter: invert(0.85) sepia(0.6) saturate(1.4) hue-rotate(345deg) brightness(0.7);
	}

	.frame-number {
		position: absolute;
		bottom: -12px;
		left: 50%;
		transform: translateX(-50%);
		font-family: "Courier New", monospace;
		font-size: 8px;
		color: rgba(0, 0, 0, 0.7);
		letter-spacing: 1px;
		font-weight: bold;
	}

	/* Responsive adjustments */
	@media (max-width: 640px) {
		.film-strip {
			--sprocket-width: 8px;
			--sprocket-height: 5px;
			--frame-gap: 6px;
		}

		.frames {
			padding-bottom: 9px;
		}

		.frame-content {
			width: 80px;
			height: 53px;
		}

		.frame-number {
			font-size: 6px;
			bottom: -8px;
		}
	}
</style>
