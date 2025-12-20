<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import { fade } from "svelte/transition";

	export let images: Array<{
		src: string;
		alt: string;
		thumb?: string;
		width?: number;
		height?: number;
	}> = [];
	export let currentIndex = 0;
	export let open = false;

	const dispatch = createEventDispatcher<{ close: void }>();

	let loading = true;

	$: currentImage = images[currentIndex];
	$: hasPrev = currentIndex > 0;
	$: hasNext = currentIndex < images.length - 1;

	// Reset loading state when image changes
	$: currentIndex, (loading = true);

	// Prevent body scroll when modal is open
	$: if (typeof document !== "undefined") {
		document.body.style.overflow = open ? "hidden" : "";
	}

	function handleClose() {
		dispatch("close");
	}

	function prev() {
		if (hasPrev) currentIndex--;
	}

	function next() {
		if (hasNext) currentIndex++;
	}

	function handleKeydown(e: KeyboardEvent) {
		if (!open) return;
		if (e.key === "Escape") {
			handleClose();
		} else if (e.key === "ArrowLeft") {
			prev();
		} else if (e.key === "ArrowRight") {
			next();
		}
	}

	function handleBackdropClick(e: MouseEvent) {
		if (e.target === e.currentTarget) {
			handleClose();
		}
	}
</script>

<svelte:window on:keydown={handleKeydown} />

{#if open && currentImage}
	<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
	<div class="modal-backdrop" transition:fade={{ duration: 200 }} on:click={handleBackdropClick}>
		<div class="modal-content">
			<button class="close-button" on:click={handleClose} aria-label="Close modal">
				<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
					<path d="M6 18L18 6M6 6l12 12" />
				</svg>
			</button>

			{#if hasPrev}
				<button class="nav-button prev" on:click={prev} aria-label="Previous image">
					<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
						<path d="M15 18l-6-6 6-6" />
					</svg>
				</button>
			{/if}

			<div
				class="image-wrapper"
				style:--img-width="{currentImage.width}px"
				style:--img-height="{currentImage.height}px"
				style:aspect-ratio={currentImage.width && currentImage.height
					? `${currentImage.width} / ${currentImage.height}`
					: undefined}
			>
				{#key currentIndex}
					<img
						src={currentImage.src}
						alt={currentImage.alt}
						width={currentImage.width}
						height={currentImage.height}
						on:load={() => (loading = false)}
					/>
					{#if loading && currentImage.thumb}
						<img class="thumb" src={currentImage.thumb} alt={currentImage.alt} />
					{/if}
					{#if loading}
						<div class="loading-overlay" transition:fade>
							<div class="loading-dots">
								<div class="dot"></div>
								<div class="dot"></div>
								<div class="dot"></div>
							</div>
						</div>
					{/if}
				{/key}
			</div>

			{#if hasNext}
				<button class="nav-button next" on:click={next} aria-label="Next image">
					<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
						<path d="M9 18l6-6-6-6" />
					</svg>
				</button>
			{/if}

			<div class="counter">{currentIndex + 1} / {images.length}</div>
		</div>
	</div>
{/if}

<style lang="postcss">
	.modal-backdrop {
		position: fixed;
		inset: 0;
		background: rgba(0, 0, 0, 0.9);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 100;
		padding: 2rem;
		overflow: hidden;
	}

	.modal-content {
		position: relative;
		max-width: min(90vw, calc(100vw - 2rem));
		max-height: 90vh;
	}

	.image-wrapper {
		position: relative;
		max-width: min(var(--img-width, 100%), 90vw);
		max-height: 80vh;
	}

	.modal-content img {
		max-width: 100%;
		max-height: 80vh;
		object-fit: contain;
		border-radius: 2px;
		box-shadow: 0 8px 32px rgba(0, 0, 0, 0.5);
	}

	.modal-content img.thumb {
		position: absolute;
		inset: 0;
		width: 100%;
		height: 100%;
		object-fit: contain;
	}

	.loading-overlay {
		position: absolute;
		inset: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		background: rgba(0, 0, 0, 0.6);
		border-radius: 2px;
	}

	.loading-dots {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.dot {
		width: 0.75rem;
		height: 0.75rem;
		background: white;
		border-radius: 50%;
		animation: pulse 2s infinite;
	}

	.dot:nth-child(2) {
		animation-delay: 200ms;
	}

	.dot:nth-child(3) {
		animation-delay: 400ms;
	}

	@keyframes pulse {
		0%,
		100% {
			opacity: 1;
		}
		50% {
			opacity: 0.3;
		}
	}

	.close-button {
		position: absolute;
		top: -40px;
		right: 0;
		background: transparent;
		border: none;
		color: white;
		cursor: pointer;
		padding: 8px;
		opacity: 0.7;
		transition: opacity 0.2s ease;
	}

	.close-button:hover {
		opacity: 1;
	}

	.close-button svg {
		width: 24px;
		height: 24px;
	}

	.nav-button {
		position: fixed;
		top: 0;
		bottom: 0;
		width: 80px;
		background: transparent;
		border: none;
		color: white;
		cursor: pointer;
		opacity: 0.5;
		transition: opacity 0.2s ease, background 0.2s ease;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.nav-button:hover {
		opacity: 1;
		background: rgba(0, 0, 0, 0.3);
	}

	.nav-button svg {
		width: 32px;
		height: 32px;
	}

	.nav-button.prev {
		left: 0;
	}

	.nav-button.next {
		right: 0;
	}

	.counter {
		position: absolute;
		bottom: -32px;
		left: 50%;
		transform: translateX(-50%);
		color: rgba(255, 255, 255, 0.7);
		font-size: 14px;
		font-variant-numeric: tabular-nums;
	}

	@media (max-width: 640px) {
		.modal-backdrop {
			padding: 1rem;
		}

		.modal-content {
			max-width: calc(100vw - 2rem);
		}

		.modal-content img {
			max-height: 70vh;
		}

		.close-button {
			top: -32px;
		}

		.close-button svg {
			width: 20px;
			height: 20px;
		}

		.nav-button {
			width: 50px;
		}

		.nav-button svg {
			width: 24px;
			height: 24px;
		}

		.counter {
			font-size: 12px;
			bottom: -28px;
		}
	}
</style>
