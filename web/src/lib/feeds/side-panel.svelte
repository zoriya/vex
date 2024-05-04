<script lang="ts">
    import Card from './card.svelte';
    import type { Feed } from '$lib/types';
    export let feeds: Feed[] = [];
</script>


<section>
    <h1>All Feeds</h1>
    <h1>Feeds</h1>
    <ul>
        {#each feeds as feed}
            <li>
                <Card {feed} />
            </li>
        {/each}
    </ul>
</section>

<style>
    ul {
        list-style: none;
        padding: 0;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

</style>