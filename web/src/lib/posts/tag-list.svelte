<script lang="ts">
	export let tags: string[] = [];

    const nbNonExpandedTags = 5;
    let isExpanded = false;
    $: tagsToDisplay = isExpanded ? tags : tags.slice(0, nbNonExpandedTags);
</script>

<div>
	{#each tagsToDisplay as tag}
		<div class="tag">{tag}</div>
	{/each}
    {#if tags.length > nbNonExpandedTags}
        <button on:click={() => isExpanded = !isExpanded}>
            {isExpanded ? "Show less" : "Show more"}
        </button>
    {/if}
</div>

<style>
	div {
		display: flex;
		gap: 0.5rem;
		flex-wrap: wrap;
	}
	.tag {
		background-color: #4d4c4c;
		padding: 0.25rem 0.5rem;
		border-radius: 5px;
        font-size: small;
        color: white;
	}

    button {
        padding: 0;
        background-color: transparent;
        border: none;
        cursor: pointer;
    }
</style>
