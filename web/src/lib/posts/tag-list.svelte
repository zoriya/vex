<script lang="ts">
	export let tags: string[] = [];

    const nbNonExpandedTags = 5;
    let isExpanded = false;
    $: tagsToDisplay = isExpanded ? tags : tags.slice(0, nbNonExpandedTags);
</script>

<div>
	{#each tagsToDisplay as tag}
		<span class="tag">{tag}</span>
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
	}

    button {
        /* border-radius: 5px; */
        padding: 0.25rem 0.5rem;
        cursor: pointer;
    }
</style>
