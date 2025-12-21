<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import { fade, fly } from "svelte/transition";

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
	let showOverlay = true;
	let showThumb = true;

	$: currentImage = images[currentIndex];

	function handleImageLoad() {
		loading = false;
		// Delay hiding thumb/overlay so full image fade-in completes first
		setTimeout(() => {
			showThumb = false;
			showOverlay = false;
		}, 250);
	}
	$: hasPrev = currentIndex > 0;
	$: hasNext = currentIndex < images.length - 1;

	// Reset loading state when image changes
	$: currentIndex, ((loading = true), (showOverlay = true), (showThumb = true));

	// Prevent body scroll when modal is open
	$: if (typeof document !== "undefined") {
		document.body.style.overflow = open ? "hidden" : "";
	}

	function handleClose() {
		dispatch("close");
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

	// Carousel state
	let touchStartX = 0;
	let trackOffset = 0;
	let isDragging = false;
	let isAnimating = false;

	$: prevImage = currentIndex > 0 ? images[currentIndex - 1] : null;
	$: nextImage = currentIndex < images.length - 1 ? images[currentIndex + 1] : null;
	$: slideWidth = typeof window !== "undefined" ? window.innerWidth : 400;

	function handleTouchStart(e: TouchEvent) {
		if (isAnimating) return;
		touchStartX = e.changedTouches[0].screenX;
		isDragging = true;
	}

	function handleTouchMove(e: TouchEvent) {
		if (!isDragging || isAnimating) return;
		const currentX = e.changedTouches[0].screenX;
		let diff = currentX - touchStartX;

		// Add resistance at boundaries
		if ((diff > 0 && !hasPrev) || (diff < 0 && !hasNext)) {
			diff = diff * 0.3;
		}
		trackOffset = diff;
	}

	function handleTouchEnd() {
		if (!isDragging || isAnimating) return;
		isDragging = false;
		const threshold = slideWidth * 0.2;

		if (trackOffset < -threshold && hasNext) {
			goToNext(true);
		} else if (trackOffset > threshold && hasPrev) {
			goToPrev(true);
		} else {
			// Snap back
			isAnimating = true;
			trackOffset = 0;
			setTimeout(() => {
				isAnimating = false;
			}, 250);
		}
	}

	function goToPrev(fromSwipe = false) {
		if (!hasPrev || isAnimating) return;
		isAnimating = true;

		const animate = () => {
			trackOffset = slideWidth;
			setTimeout(() => {
				currentIndex--;
				trackOffset = 0;
				isAnimating = false;
			}, 250);
		};

		// For button clicks (trackOffset is 0), wait a frame for transition class
		// For swipes, animate immediately from current position
		if (fromSwipe) {
			animate();
		} else {
			requestAnimationFrame(animate);
		}
	}

	function goToNext(fromSwipe = false) {
		if (!hasNext || isAnimating) return;
		isAnimating = true;

		const animate = () => {
			trackOffset = -slideWidth;
			setTimeout(() => {
				currentIndex++;
				trackOffset = 0;
				isAnimating = false;
			}, 250);
		};

		// For button clicks (trackOffset is 0), wait a frame for transition class
		// For swipes, animate immediately from current position
		if (fromSwipe) {
			animate();
		} else {
			requestAnimationFrame(animate);
		}
	}

	function prev() {
		goToPrev();
	}

	function next() {
		goToNext();
	}
</script>

<svelte:window on:keydown={handleKeydown} />

{#if open && currentImage}
	<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
	<div
		class="modal-backdrop"
		transition:fade={{ duration: 200 }}
		on:click={handleBackdropClick}
	>
		<button class="close-button" on:click|stopPropagation={handleClose} aria-label="Close modal">
			<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
				<path d="M6 18L18 6M6 6l12 12" />
			</svg>
		</button>

		<div class="counter desktop">{currentIndex + 1} / {images.length}</div>

		<button class="nav-button prev" class:hidden={!hasPrev} on:click={prev} aria-label="Previous image">
			<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
				<path d="M15 18l-6-6 6-6" />
			</svg>
		</button>

		<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
		<div
			class="carousel-container"
			on:touchstart={handleTouchStart}
			on:touchmove={handleTouchMove}
			on:touchend={handleTouchEnd}
			on:click={handleBackdropClick}
		>
			<div
				class="carousel-track"
				class:dragging={isDragging}
				class:animating={isAnimating}
				style:transform="translateX(calc(-100vw + {trackOffset}px))"
			>
				<!-- Previous slide -->
				<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
				<div class="carousel-slide" on:click={handleClose}>
					{#if prevImage}
						<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
						<div
							class="image-container"
							on:click|stopPropagation
							style:--img-width="{prevImage.width}px"
							style:--img-height="{prevImage.height}px"
							style:aspect-ratio={prevImage.width && prevImage.height
								? `${prevImage.width} / ${prevImage.height}`
								: undefined}
						>
							<img
								src={prevImage.thumb || prevImage.src}
								alt={prevImage.alt}
								width={prevImage.width}
								height={prevImage.height}
							/>
							<div class="loading-overlay">
								<div class="loading-dots">
									<div class="dot"></div>
									<div class="dot"></div>
									<div class="dot"></div>
								</div>
							</div>
						</div>
					{/if}
				</div>

				<!-- Current slide -->
				<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
				<div class="carousel-slide current" on:click={handleClose}>
					{#key currentIndex}
						<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
						<div
							class="image-container"
							on:click|stopPropagation
							style:--img-width="{currentImage.width}px"
							style:--img-height="{currentImage.height}px"
							style:aspect-ratio={currentImage.width && currentImage.height
								? `${currentImage.width} / ${currentImage.height}`
								: undefined}
						>
							<img
								class="full"
								class:loaded={!loading}
								src={currentImage.src}
								alt={currentImage.alt}
								width={currentImage.width}
								height={currentImage.height}
								on:load={handleImageLoad}
							/>
							{#if showThumb && currentImage.thumb}
								<img class="thumb" src={currentImage.thumb} alt={currentImage.alt} transition:fade={{ duration: 200 }} />
							{/if}
							{#if showOverlay}
								<div class="loading-overlay" transition:fade>
									<div class="loading-dots">
										<div class="dot"></div>
										<div class="dot"></div>
										<div class="dot"></div>
									</div>
								</div>
							{/if}
						</div>
					{/key}
				</div>

				<!-- Next slide -->
				<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
				<div class="carousel-slide" on:click={handleClose}>
					{#if nextImage}
						<!-- svelte-ignore a11y_click_events_have_key_events a11y_no_static_element_interactions -->
						<div
							class="image-container"
							on:click|stopPropagation
							style:--img-width="{nextImage.width}px"
							style:--img-height="{nextImage.height}px"
							style:aspect-ratio={nextImage.width && nextImage.height
								? `${nextImage.width} / ${nextImage.height}`
								: undefined}
						>
							<img
								src={nextImage.thumb || nextImage.src}
								alt={nextImage.alt}
								width={nextImage.width}
								height={nextImage.height}
							/>
							<div class="loading-overlay">
								<div class="loading-dots">
									<div class="dot"></div>
									<div class="dot"></div>
									<div class="dot"></div>
								</div>
							</div>
						</div>
					{/if}
				</div>
			</div>
		</div>

		<button class="nav-button next" class:hidden={!hasNext} on:click={next} aria-label="Next image">
			<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
				<path d="M9 18l6-6-6-6" />
			</svg>
		</button>

		<div class="nav-row">
			<button class="nav-inline prev" class:invisible={!hasPrev} on:click={prev} aria-label="Previous image">
				<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
					<path d="M15 18l-6-6 6-6" />
				</svg>
			</button>
			<div class="counter">{currentIndex + 1} / {images.length}</div>
			<button class="nav-inline next" class:invisible={!hasNext} on:click={next} aria-label="Next image">
				<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
					<path d="M9 18l6-6-6-6" />
				</svg>
			</button>
		</div>
	</div>
{/if}

<style>
	.modal-backdrop {
		position: fixed;
		inset: 0;
		background: rgba(0, 0, 0, 0.9);
		z-index: 100;
		overflow: hidden;
	}

	.carousel-container {
		position: absolute;
		inset: 0;
		overflow: hidden;
	}

	.carousel-track {
		display: flex;
		width: 300vw;
		height: 100%;
	}

	.carousel-track.dragging {
		/* No transition while dragging */
	}

	.carousel-track.animating {
		transition: transform 0.25s ease-out;
	}

	.carousel-slide {
		width: 100vw;
		height: 100%;
		flex-shrink: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 2rem;
	}

	.image-container {
		position: relative;
		max-width: min(var(--img-width, 100%), 90vw);
		max-height: 80vh;
	}

	.image-container img {
		max-width: 100%;
		max-height: 80vh;
		object-fit: contain;
		border-radius: 2px;
		box-shadow: 0 8px 32px rgba(0, 0, 0, 0.5);
	}

	.image-container img.full {
		opacity: 0;
		transition: opacity 0.2s ease;
	}

	.image-container img.full.loaded {
		opacity: 1;
	}

	.image-container img.thumb {
		position: absolute;
		inset: 0;
		width: 100%;
		height: 100%;
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
		position: fixed;
		top: 1rem;
		right: 1rem;
		background: transparent;
		border: none;
		color: white;
		cursor: pointer;
		padding: 8px;
		opacity: 0.7;
		transition: opacity 0.2s ease;
		z-index: 20;
	}

	.counter.desktop {
		position: fixed;
		bottom: 1rem;
		left: 50%;
		transform: translateX(-50%);
		color: rgba(255, 255, 255, 0.7);
		font-size: 14px;
		font-variant-numeric: tabular-nums;
		z-index: 10;
	}

	.close-button:hover {
		opacity: 1;
	}

	.close-button svg {
		width: 24px;
		height: 24px;
	}

	.hidden {
		display: none !important;
	}

	.invisible {
		visibility: hidden;
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
		transition: opacity 0.2s ease;
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 10;
	}

	.nav-button:hover {
		opacity: 1;
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

	.nav-row {
		display: none;
		align-items: center;
		justify-content: center;
		gap: 1rem;
		z-index: 10;
	}

	.nav-inline {
		background: transparent;
		border: none;
		color: white;
		cursor: pointer;
		opacity: 0.7;
		padding: 0.75rem;
		transition: opacity 0.2s ease;
	}

	.nav-inline:hover {
		opacity: 1;
	}

	.nav-inline svg {
		width: 24px;
		height: 24px;
	}

	.counter {
		color: rgba(255, 255, 255, 0.7);
		font-size: 14px;
		font-variant-numeric: tabular-nums;
	}

	@media (max-width: 640px) {
		.counter.desktop {
			display: none;
		}

		.carousel-slide {
			padding: 1rem;
			padding-bottom: 5rem;
		}

		.image-container {
			max-height: 60vh;
		}

		.image-container img {
			max-height: 60vh;
		}

		.close-button {
			top: 0.5rem;
			right: 0.5rem;
		}

		.close-button svg {
			width: 20px;
			height: 20px;
		}

		.nav-button {
			display: none;
		}

		.nav-row {
			display: flex;
			position: fixed;
			bottom: 0;
			left: 0;
			right: 0;
			justify-content: center;
			padding: 1rem 0;
			background: rgba(0, 0, 0, 0.8);
		}

		.nav-inline {
			padding: 0.75rem 1.5rem;
		}

		.nav-inline svg {
			width: 28px;
			height: 28px;
		}

		.counter {
			font-size: 16px;
			min-width: 4rem;
			text-align: center;
		}
	}
</style>
