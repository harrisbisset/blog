<div class="grid h-lvh grid-cols-3 gap-4">
	<div id="chess-box-1"></div>
	<div id="chess-box-2"></div>
	<div id="chess-box-3"></div>
	<div id="checkered-blue-1"></div>
	<div></div>
	<div></div>
	<div id="patchwork-1"></div>
</div>

<style>
	#chess-box-1 {
		background-image:
			linear-gradient(45deg, #808080 25%, transparent 25%),
			linear-gradient(-45deg, #808080 25%, transparent 25%),
			linear-gradient(45deg, transparent 75%, #808080 75%),
			linear-gradient(-45deg, transparent 75%, #808080 75%);
		background-size: 20px 20px;
		background-position:
			0 0,
			0 10px,
			10px -10px,
			-10px 0px;
	}

	#chess-box-2 {
		background-image:
			linear-gradient(45deg, #808080 25%, transparent 25%),
			linear-gradient(-45deg, #808080 25%, transparent 25%),
			linear-gradient(45deg, transparent 75%, #808080 75%);
		background-size: 20px 20px;
		background-position:
			0 0,
			0 10px,
			10px -10px,
			-10px 0px;
	}

	#chess-box-3 {
		background-image:
			linear-gradient(45deg, #808080 25%, transparent 25%),
			linear-gradient(-45deg, #808080 25%, transparent 25%),
			linear-gradient(45deg, transparent 75%, #808080 75%),
			linear-gradient(-45deg, transparent 75%, #808080 75%);
		background-size: 20px 20px;
		background-position:
			0 0,
			0 10px,
			10px -10px;
	}

	#checkered-blue-1 {
		background-image: linear-gradient(90deg, #0542c4 5%, transparent 0%);
		background-size: 30px 30px;
		background-position: 0 0;
	}

	#patchwork-1 {
		background-image:
			linear-gradient(90deg, #0542c4 50%, transparent 0%),
			linear-gradient(90deg, #f6aa1d 100%, transparent 0%);
		background-size: 100%;
		background-position: left;
	}
</style>
