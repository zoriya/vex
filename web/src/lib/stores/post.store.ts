import type { Post } from "$lib/types";
import { writable } from 'svelte/store';

const initialState: Post[] = [];

export const postStore = writable<Post[]>(initialState);

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