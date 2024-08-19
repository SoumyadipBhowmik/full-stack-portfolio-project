<script lang="ts">
    import { onMount } from 'svelte';

    let posts: string | any[] = [];

    onMount(async () => {
        try {
            const response = await fetch('http://127.0.0.1:5010/posts');
            if (response.ok) {
                const data = await response.json();
                console.log('Fetched data:', data);
                posts = data;
            } else {
                console.error('Failed to fetch:', response.status, response.statusText);
            }
        } catch (error) {
            console.error('Error fetching posts:', error);
        }
    });
</script>

<main>
    <h1>Hi from the API</h1>
    {#if posts.length > 0}
        {#each posts as post}
            <div>
                <img src={post.image_url} alt={post.description} />
                <p>{post.description}</p>
            </div>
        {/each}
    {:else}
        <p>No posts found.</p>
    {/if}
</main>
