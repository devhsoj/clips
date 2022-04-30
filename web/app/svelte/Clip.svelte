<!-- svelte-ignore a11y-media-has-caption -->

<script lang="ts">
    import { getClip, incrementViews } from '../lib/clip'
    import { link } from 'svelte-spa-router'
    import { getMe } from '../lib/user'
    import { onMount } from 'svelte'

    export let params = {}

    let clip = null
    let user = null

    async function onClipEnd(clip_id) {
        try {
            await incrementViews(clip_id)

            clip.views += 1
        } catch(err) {
            console.trace(err)
        }
    }

    async function loadClip() {
        try {
            clip = await getClip({
                clip_id:params['clip_id']
            })
        } catch(err) {
            console.trace(err)
        }
    }

    onMount(async () => {
        try {
            user = await getMe()
        } catch (err) {
            console.trace(err)
        }
    })
</script>

<svelte:head>
    <title>{clip?.title} ({clip?.creator}) - Project Clips</title>

    <meta property="og:site_name" content="Project Clips" />
    <meta property="og:title" content="{clip?.title}" />

    <meta name="description" content="{clip?.description}" />
    <meta property="og:description" content="{clip?.description}" />

    <meta property="og:url" content="{window.location.origin}" />
    <meta property="og:video" content="{window.location.origin}/api/clips/view/{clip?.clip_id}" />
    <meta property="og:video:url" content="{window.location.origin}/api/clips/view/{clip?.clip_id}" />
    <meta property="og:video:type" content="{clip?.type}" />
    <meta property="og:video:width" content="1920">
    <meta property="og:video:height" content="1080">
</svelte:head>

<div class="content-wrapper" on:load={loadClip()}>
    <section class="wrapper bg-soft-primary">
        <div class="container pt-5 pb-5 text-center">
            <div class="row">
                <div class="col-md-9 col-lg-7 col-xl-6 mx-auto">
                    <h2 class="display-1 mb-3">Project Clips</h2>
                </div>
            </div>
        </div>
    </section>

    <div class="container">
        <div class="row gx-md-11 gx-lg-0">
            <div id="clipsContainer" class="col-xl-9 order-xl-2">
                <section class="wrapper pt-8 pb-2">
                    {#if clip === null}
                        <p>Clip not found.</p>
                    {:else}
                        <section class="container clip">
                            <h3 class="clip-link"><a href="/clip/{clip.clip_id}" use:link style="color:inherit;text-decoration:inherit;">{clip.title}</a></h3>
                            <p class="small" style="color:grey;">
                                {clip.description} - <strong>{clip.creator}</strong>
                            </p>
                            <p class="small">
                                {clip.views} Views
                            </p>
                            <hr class="my-4" />

                            <video class="player" on:ended={onClipEnd(clip.clip_id)} preload="metadata" controls>
                                <source type={clip.type} src="/api/clip/view/{clip.clip_id}">
                            </video>
                        </section>
                        <hr class="my-8" />
                    {/if}
                </section>
            </div>
            <aside class="col-xl-3 sidebar sticky-sidebar mt-md-0 py-16 d-none d-xl-block">
                <nav>
                    <ul class="nav list-unstyled lh-lg text-reset d-block">
                        <li class="nav-item"><a class="nav-link" href="/" use:link>Home</a></li>
                        {#if user}
                            <li class="nav-item"><a class="nav-link" href="/upload" use:link>Upload</a></li>
                        {:else}
                            <li class="nav-item"><a class="nav-link" href="/login" use:link>Login</a></li>
                        {/if}
                    </ul>
                </nav>
            </aside>
        </div>
    </div>
</div>

<style>
    section::-webkit-scrollbar {
        display:none;
    }

    video {
        width:100%;
        height:100%;
    }

    .clip {
        overflow-y:scroll;
        scrollbar-width:none;
        padding-left:0;
        padding-right:0;
    }

    .clip-link:hover {
        color:#3f78e0;
        cursor:pointer;
    }
</style>