import { component$ } from '@builder.io/qwik';
import type { DocumentHead } from '@builder.io/qwik-city';
import Hero from '~/components/home/hero/hero';

export default component$(() => {
    return <Hero />;
});

export const head: DocumentHead = {
    title: "Rafaleo - Open Source Performance Testing Tool ",
    meta: [
        {
            name: "description",
            content: "Testez et optimisez vos applications pour le long terme avec des charges réalistes grâce à Rafaleo !",
        },
    ],
};
