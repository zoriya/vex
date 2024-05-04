// /home/cbihan/Documents/front_stuff/vex/web/src/lib/stores/post.store.js
import type { Post } from "$lib/types";
import { writable } from 'svelte/store';

// Initial state
const initialState: Post[] = [];

// Create the writable store
export const postStore = writable<Post[]>(initialState);

// Actions
export const addPost = (post: Post) => {
    postStore.update((posts) => [...posts, post]);
};

export const removePost = (postId: string) => {
    postStore.update((posts) => posts.filter((post) => post.id !== postId));
};

export const updatePost = (postId: string, updatedPost: Post) => {
    postStore.update((posts) =>
        posts.map((post) => (post.id === postId ? updatedPost : post))
    );
};